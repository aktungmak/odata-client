// this package provides utilities for interacting with
// ccm in a variety of different ways.
package ccm

import (
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const (
	EP_TOKENSERVICE     = "/rest/v0/TokenService/Actions/TokenService.Token"
	EP_ALARMCOLLECTION  = "/rest/v0/AlarmServices/AlarmListService/AlarmEntryCollection"
	EP_SYSTEMCOLLECTION = "/rest/v0/Systems"
	TRUST_BAD_CERT      = true
)

// a basic client for API interactions. handles token management
type Client struct {
	Host     string
	Username string
	Password string
	Token    string
}

func NewClient(host, uname, pass string) (*Client, error) {
	c := &Client{host, uname, pass, ""}
	var err error
	c.Token, err = GetToken(host, uname, pass)
	if err != nil {
		return c, err
	} else {
		return c, nil
	}
}

// make a request with authentication, return raw http.Response
func (c *Client) DoRaw(meth, uri string) (*http.Response, error) {
	tr := &http.Transport{}
	if TRUST_BAD_CERT {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	client := &http.Client{Transport: tr}

	req, err := http.NewRequest(meth, uri, nil)
	var retries int = 2
	var res *http.Response
	for retries > 0 {
		req.Header.Set("Authorization", "Bearer "+c.Token)

		res, err = client.Do(req)
		if err != nil {
			return nil, err
		} else if res.StatusCode == 401 {
			//print("token was stale, getting a new one\n")
			c.Token, err = GetToken(c.Host, c.Username, c.Password)
			if err != nil {
				return nil, err
			}

		} else {
			break
		}
		retries -= 1
	}
	return res, nil

}

// perform a get request with authentication
func (c *Client) Get(ep string) ([]byte, error) {
	uri := "https://" + c.Host + ep

	res, err := c.DoRaw("GET", uri)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	return data, err

}

// convenience method to grab the current alarms
func (c *Client) GetAlarms() (*AlarmCollection, error) {
	data, err := c.Get(EP_ALARMCOLLECTION)
	if err != nil {
		return nil, err
	}
	return NewAlarmCollection(data)
}

// convenience method to grab all systems
func (c *Client) GetSystems() (*SystemCollection, error) {
	data, err := c.Get(EP_SYSTEMCOLLECTION)
	if err != nil {
		return nil, err
	}
	sc, err := NewSystemCollection(data)

	// need to actually populate the Members field
	for _, mem := range sc.Links.Members {
		data, err := c.Get(mem.Id)
		s, err := NewSystem(data)
		if err != nil {
			continue
		} else {
			sc.Members = append(sc.Members, *s)
		}
	}
	return sc, err
}

// get a bearer token for the specified user:pass
func GetToken(host, uname, pass string) (string, error) {
	tr := &http.Transport{}
	if TRUST_BAD_CERT {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	client := &http.Client{Transport: tr}
	uri := "https://" + host + EP_TOKENSERVICE

	req, err := http.NewRequest("GET", uri, nil)
	req.SetBasicAuth(uname, pass)

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	data := struct {
		AccessToken string `json:"access_token"`
	}{}
	err = decoder.Decode(&data)
	if err != nil {
		return "", err
	} else {
		return data.AccessToken, nil
	}
}
