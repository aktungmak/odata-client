// this package provides utilities for interacting with
// odata in a variety of different ways.
package odata

import (
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
type Client interface {
	DoRaw(string, string, []byte) (*http.Response, error)
	Insecure() bool
	Token() string
}

// perform a get request with authentication
func Get(c Client, host, ep string) ([]byte, error) {
	uri := "https://" + host + ep

	res, err := c.DoRaw("GET", uri, nil)
	if err != nil {
		return []byte{}, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}
	return data, err
}

// convenience method to grab the current alarms
func GetAlarms(c Client, host string) (*AlarmCollection, error) {
	data, err := Get(c, host, EP_ALARMCOLLECTION)
	if err != nil {
		return nil, err
	}
	return NewAlarmCollection(data)
}

// convenience method to grab all systems
func GetSystems(c Client, host string) (*SystemCollection, error) {
	data, err := Get(c, host, EP_SYSTEMCOLLECTION)
	if err != nil {
		return nil, err
	}
	sc, err := NewSystemCollection(data)

	// need to actually populate the Members field
	for _, mem := range sc.Links.Members {
		data, err := Get(c, host, mem.Id)
		s, err := NewSystem(data)
		if err != nil {
			continue
		} else {
			sc.Members = append(sc.Members, *s)
		}
	}
	return sc, err
}
