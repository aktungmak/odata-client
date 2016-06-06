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

var SampleSystemCollection = `{
  "@odata.context" : "/rest/v0/$metadata#Systems",
  "@odata.id" : "/rest/v0/Systems",
  "@odata.type" : "#ComputerSystem.0.96.0.ComputerSystemCollection",
  "Name" : "Computer System Collection",
  "Links" : {
    "Members" : [ {
      "@odata.id" : "/rest/v0/Systems/1b51ce2d-49ab-e311-91f0-8d33f8f9a2dd"
    }, {
      "@odata.id" : "/rest/v0/Systems/1b51ce2d-49ab-e311-91f0-b8c75753f1d1"
    }, {
      "@odata.id" : "/rest/v0/Systems/1b51ce2d-49ab-e311-91f0-3dacfeb8fa46"
    }, {
      "@odata.id" : "/rest/v0/Systems/1b51ce2d-49ab-e311-91f0-b1f13ed31234"
    }, {
      "@odata.id" : "/rest/v0/Systems/1b51ce2d-49ab-e311-91f0-dedc9b57ec74"
    }, {
      "@odata.id" : "/rest/v0/Systems/1b51ce2d-49ab-e311-91f0-f2c4ab29792e"
    }, {
      "@odata.id" : "/rest/v0/Systems/1b51ce2d-49ab-e311-91f0-33bd8b0fe2a4"
    }, {
      "@odata.id" : "/rest/v0/Systems/1b51ce2d-49ab-e311-91f0-977c9a8d8a74"
    }, {
      "@odata.id" : "/rest/v0/Systems/1b51ce2d-49ab-e311-91f0-1c013d7abf89"
    }, {
      "@odata.id" : "/rest/v0/Systems/1b51ce2d-49ab-e311-91f0-3b77ff4a51df"
    } ],
    "Members@odata.count" : 10
  }
}`

var SampleSystem = `{
  "@odata.context" : "/rest/v0/$metadata#HdsComputerSystem.0.11.0",
  "@odata.id" : "/rest/v0/Systems/1b51ce2d-49ab-e311-91f0-3b77ff4a51df",
  "@odata.type" : "#HdsComputerSystem.0.11.0.HdsComputerSystem",
  "Oem" : {
    "Ericsson" : {
      "ManagementState" : "ManagedFree",
      "MetricsEnabled" : true,
      "BootSourceOverrideTargetAllowableValues" : [ "Pxe", "Hdd" ]
    }
  },
  "Id" : "1b51ce2d-49ab-e311-91f0-3b77ff4a51df",
  "Name" : "3b77ff4a51df",
  "Links" : {
    "Chassis" : [ ],
    "ManagedBy" : [ ],
    "EthernetInterfaces" : {
      "@odata.id" : "/rest/v0/Systems/1b51ce2d-49ab-e311-91f0-3b77ff4a51df/EthernetInterfaces"
    },
    "PoweredBy" : [ ],
    "CooledBy" : [ ]
  },
  "Manufacturer" : "Hewlett-Packard",
  "Model" : "20238",
  "SKU" : "Hewlett-Packard_MT_20238",
  "SerialNumber" : "3385372101747",
  "UUID" : "1b51ce2d-49ab-e311-91f0-3b77ff4a51df",
  "Boot" : { },
  "BiosVersion" : "79CN46WW(V3.05)",
  "ProcessorSummary" : {
    "Count" : 1,
    "Model" : "Core i5"
  },
  "MemorySummary" : {
    "TotalSystemMemoryGiB" : 16
  },
  "Actions" : {
    "http://www.ericsson.com/hds8000#HdsComputerSystem.ConnectRemoteConsole" : {
      "title" : "HdsComputerSystem.ConnectRemoteConsole",
      "target" : "/rest/v0/Systems/1b51ce2d-49ab-e311-91f0-3b77ff4a51df/Actions/HdsComputerSystem.ConnectRemoteConsole"
    },
    "#ComputerSystem.Reset" : {
      "target" : "/rest/v0/Systems/1b51ce2d-49ab-e311-91f0-3b77ff4a51df/Actions/ComputerSystem.Reset",
      "ResetType@DMTF.AllowableValues" : [ "On", "ForceOff", "ForceRestart" ]
    }
  }
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

func TestSystemCollection(t *testing.T) {
	sc, err := NewSystemCollection([]byte(SampleSystemCollection))
	if err != nil {
		t.Error("Failed to create SystemCollection: ", err)
	}
	if len(sc.Links.Members) != sc.Links.Count {
		t.Error("Mismatch between links member count and links count")
	}
	m1 := sc.Links.Members[1]
	if m1.Id != "/rest/v0/Systems/1b51ce2d-49ab-e311-91f0-b8c75753f1d1" {
		t.Error("second link member ID is incorrect")
	}
}

func TestSystem(t *testing.T) {
	s, err := NewSystem([]byte(SampleSystem))
	if err != nil {
		t.Error("Failed to create System: ", err)
	}
	if //s.ManagementState != "ManagedFree" ||
	s.Id != "1b51ce2d-49ab-e311-91f0-3b77ff4a51df" ||
		s.Name != "3b77ff4a51df" ||
		s.Manufacturer != "Hewlett-Packard" ||
		s.Model != "20238" ||
		s.SKU != "Hewlett-Packard_MT_20238" ||
		s.SerialNumber != "3385372101747" ||
		s.UUID != "1b51ce2d-49ab-e311-91f0-3b77ff4a51df" ||
		s.BiosVersion != "79CN46WW(V3.05)" {
		//s.ProcessorCount != 1 ||
		//s.ProcessorModel != "Core i5" ||
		//s.TotalSystemMemory != 16 {
		t.Error("Values of System incorrect")
	}
}
