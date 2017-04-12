package odata

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"net/http"
)

const (
	AUTH_RETRIES = 2
)

// this implements the odata.Client interface
// but uses a user-provided token for auth
type ManualClient struct {
	Host   string
	Token  string
	client *http.Client
}

func NewManualClient(host, token string, acceptBadCert bool) (*ManualClient, error) {
	c := &ManualClient{Host: host, Token: token}

	tr := &http.Transport{}
	if acceptBadCert {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	c.client = &http.Client{Transport: tr}

	return c, nil
}

func (c *ManualClient) DoRaw(meth, uri string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(meth, uri, bytes.NewReader(body))
	req.Header.Set("Authorization", "Bearer "+c.Token)

	var res *http.Response
	res, err = c.client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}
