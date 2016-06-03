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
	EP_TOKENSERVICE = "/rest/v0/TokenService/Actions/TokenService.Token"
    EP_ALARMCOLLECTION = "/rest/v0/AlarmServices/AlarmListService/AlarmEntryCollection"
	TRUST_BAD_CERT  = true
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

// perform a get request with authentication
func (c *Client) Get(ep string) ([]byte, error) {
	tr := &http.Transport{}
	if TRUST_BAD_CERT {
		tr.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	client := &http.Client{Transport: tr}
	uri := "https://" + c.Host + ep

	req, err := http.NewRequest("GET", uri, nil)
	req.Header.Add("Authorization", "Bearer "+c.Token)

	res, err := client.Do(req)
	if err != nil {
        // TODO check if token is stale and retry
		return nil, err
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

// get a bearer token for the specified user:pass combo
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
