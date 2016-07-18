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
// but uses JWT for authentication
type TokenClient struct {
	Host     string
	Username string
	Password string
	Token    string
	client   *http.Client
}

func NewTokenClient(host, uname, pass string, acceptBadCert bool) (*TokenClient, error) {
	c := &TokenClient{Host: host, Username: uname, Password: pass}

	tr := &http.Transport{}
	if acceptBadCert {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	c.client = &http.Client{Transport: tr}

	var err error
	err = c.GetToken()
	if err != nil {
		return c, err
	} else {
		return c, nil
	}
}

func (c *TokenClient) DoRaw(meth, uri string, body []byte) (*http.Response, error) {
	req, err := http.NewRequest(meth, uri, bytes.NewReader(body))
	var res *http.Response
	for i := 0; i < AUTH_RETRIES; i++ {
		req.Header.Set("Authorization", "Bearer "+c.Token)
		res, err = c.client.Do(req)

		if err != nil {
			return nil, err
		}
		if res.StatusCode == 401 {
			// the token was rejected, get a new one
			err = c.GetToken()
			if err != nil {
				return nil, err
			}
		} else {
			// the request was successful, no need to retry
			break
		}
	}
	return res, nil
}

// request and update our JWT
func (c *TokenClient) GetToken() error {
	uri := "https://" + c.Host + EP_TOKENSERVICE

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return err
	}

	req.SetBasicAuth(c.Username, c.Password)

	res, err := c.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	data := struct {
		AccessToken string `json:"access_token"`
	}{}
	err = decoder.Decode(&data)
	if err != nil {
		return err
	} else {
		c.Token = data.AccessToken
		return nil
	}
}
