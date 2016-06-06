package ccm

import (
	"encoding/json"
)

type AlarmCollection struct {
	Members []Alarm `json:"Members"`
}

type Alarm struct {
	Description     string
	Name            string
	Id              string
	ResourceId      string
	AlarmId         int
	ProducerId      string
	AlarmRaisedTime string
	Severity        string
}

type SystemCollection struct {
	Members []System `json:"Members"`
	Links   struct {
		Count   int `json:"Members@odata.count"`
		Members []struct {
			Id string `json:"@odata.id"`
		}
	}
}

type System struct {
	//ManagementState   string `json:"Oem>Ericsson>ManagementState"`
	Id           string
	Name         string
	Manufacturer string
	Model        string
	SKU          string
	SerialNumber string
	UUID         string
	BiosVersion  string
	//ProcessorCount    int    `json:"ProcessorSummary>Count"`
	//ProcessorModel    string `json:"ProcessorSummary>Model"`
	//TotalSystemMemory int    `json:"MemorySummary>TotalSystemMemoryGiB"`
}

func NewAlarmCollection(data []byte) (*AlarmCollection, error) {
	ac := &AlarmCollection{}
	err := json.Unmarshal(data, ac)
	return ac, err
}

// this function populates the Links field, but it does not
// pull the data for the Members field using those links
func NewSystemCollection(data []byte) (*SystemCollection, error) {
	sc := &SystemCollection{}
	err := json.Unmarshal(data, sc)
	return sc, err
}

func NewSystem(data []byte) (*System, error) {
	s := &System{}
	err := json.Unmarshal(data, s)
	return s, err
}
