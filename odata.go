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

func NewAlarmCollection(data []byte) (*AlarmCollection, error) {
	ac := &AlarmCollection{}
	err := json.Unmarshal(data, ac)
	return ac, err
}
