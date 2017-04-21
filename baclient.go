package odata

import (
	"bytes"
	"crypto/tls"
	"net/http"
)

// this implements the odata.Client interface
// but uses HTTP Basic Auth for authentication
type BaClient struct {
	Username string
	Password string
	insecure bool
	client   *http.Client
}

func NewBaClient(uname, pass string, acceptBadCert bool) *BaClient {
	c := &BaClient{Username: uname, Password: pass}

	c.insecure = acceptBadCert
	tr := &http.Transport{}
	if acceptBadCert {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	c.client = &http.Client{Transport: tr}
	return c
}

func (c BaClient) DoRaw(meth, uri string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(meth, uri, bytes.NewReader(body))
	req.SetBasicAuth(c.Username, c.Password)

	var res *http.Response
	res, err = c.client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c BaClient) Insecure() bool {
	return c.insecure
}

func (c BaClient) Token() string {
	return c.Username + ":" + c.Password
}
