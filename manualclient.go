package odata

import (
	"bytes"
	"crypto/tls"
	"net/http"
)

const (
	AUTH_RETRIES = 2
)

// this implements the odata.Client interface
// but uses a user-provided token for auth
type ManualClient struct {
	token    string
	insecure bool
	client   *http.Client
}

func NewManualClient(token string, acceptBadCert bool) *ManualClient {
	c := &ManualClient{token: token}

	c.insecure = acceptBadCert
	tr := &http.Transport{}
	if acceptBadCert {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	c.client = &http.Client{Transport: tr}

	return c
}

func (c ManualClient) DoRaw(meth, uri string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(meth, uri, bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+c.token)

	var res *http.Response
	res, err = c.client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c ManualClient) Insecure() bool {
	return c.insecure
}

func (c ManualClient) Token() string {
	return c.token
}
