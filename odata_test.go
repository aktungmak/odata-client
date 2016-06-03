package ccm

import (
	"testing"
)

var sampleAlarm string = `
{
  "@odata.context" : "/rest/v0/$metadata#AlarmEntryCollection",
  "@odata.id" : "/rest/v0/AlarmServices/AlarmListService/AlarmEntryCollection",
  "@odata.type" : "#AlarmEntryCollection.0.4.4.AlarmEntryCollection",
  "Name" : "List of all active alarms",
  "Members" : [ {
    "@odata.context" : "/rest/v0/$metadata#AlarmEntry",
    "@odata.id" : "/rest/v0/AlarmServices/AlarmListService/AlarmEntryCollection/2",
    "@odata.type" : "#AlarmEntry.0.4.4.AlarmEntry",
    "Description" : "Capacity warning threshold reached",
    "Name" : "License capacity usage threshold reached",
    "Id" : "2",
    "ResourceId" : "1b51ce2d-49ab-e311-91f0-b8c75753f1d1",
    "AlarmId" : 2,
    "ProducerId" : "HDS 8000",
    "AlarmRaisedTime" : "2016-05-26T18:41:42.137Z",
    "AdditionalInformation" : " [ MetricName : cpu.user ]  [ MetricType : cpu ] ",
    "AlarmType" : "other (1)",
    "ProbableCause" : "thresholdCrossed (549)",
    "Severity" : "Critical",
    "ThresholdInformation" : {
      "ThresholdId" : "/rest/v0/MonitoringService/Thresholds/1",
      "ThresholdLevel" : 6100000.0,
      "ThresholdValue" : 6100097.0
    }
  }, {
    "@odata.context" : "/rest/v0/$metadata#AlarmEntry",
    "@odata.id" : "/rest/v0/AlarmServices/AlarmListService/AlarmEntryCollection/1",
    "@odata.type" : "#AlarmEntry.0.4.4.AlarmEntry",
    "Description" : "Capacity warning threshold reached",
    "Name" : "License capacity usage threshold reached",
    "Id" : "1",
    "ResourceId" : "5d7b6507-59cc-4d89-986a-a397aeb47edc",
    "AlarmId" : 1,
    "ProducerId" : "HDS 8000",
    "AlarmRaisedTime" : "2016-05-17T18:38:25.366Z",
    "Severity" : "Warning"
  } ],
  "Members@odata.count" : 2
}`

func TestAlarmCollection(t *testing.T) {
	ac, err := NewAlarmCollection([]byte(sampleAlarm))
	if err != nil {
		t.Error("Failed to create AlarmCollection: ", err)
	}
	if len(ac.Members) != 2 {
		t.Error("Incorrect number of Members: ", len(ac.Members))
	}
	a := ac.Members[1]
	if a.Description != "Capacity warning threshold reached" ||
		a.Name != "License capacity usage threshold reached" ||
		a.Id != "1" ||
		a.ResourceId != "5d7b6507-59cc-4d89-986a-a397aeb47edc" ||
		a.AlarmId != 1 ||
		a.ProducerId != "HDS 8000" ||
		a.AlarmRaisedTime != "2016-05-17T18:38:25.366Z" ||
		a.Severity != "Warning" {
		t.Error("Values of Alarm incorrect")
	}
}
