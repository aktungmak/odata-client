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
	client   *http.Client
}

func NewBaClient(uname, pass string, acceptBadCert bool) *BaClient {
	c := &BaClient{Username: uname, Password: pass}

	tr := &http.Transport{}
	if acceptBadCert {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	c.client = &http.Client{Transport: tr}
	return c
}

func (c *BaClient) DoRaw(meth, uri string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(meth, uri, bytes.NewReader(body))
	var res *http.Response
	req.SetBasicAuth(c.Username, c.Password)
	res, err = c.client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}
