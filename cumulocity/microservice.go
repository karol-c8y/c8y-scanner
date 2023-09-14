package cumulocity

import (
    "github.com/reubenmiller/go-c8y/pkg/microservice"
    "github.com/reubenmiller/go-c8y/pkg/c8y"
)

var (
	ms = microservice.NewDefaultMicroservice(microservice.Options{})
)

func Init() {
	ms.Config.SetDefault("application.name", "c8y-scanner-clamav")
	ms.Config.SetDefault("agent.identityType", "c8y-scanner")
	ms.RegisterMicroserviceAgent()
}

func DownloadFile(file_id string) string {
	filepath, _ := ms.Client.Inventory.DownloadBinary(ms.WithServiceUser(), file_id)
	return filepath
}

func RaiseEvent(eventType string, text string) {
	event := c8y.Event {
		Time:   c8y.NewTimestamp(),
		Type:   eventType,
		Text:   text,
		Source: c8y.NewSource(ms.GetAgent().ID),
	}
	ms.Client.Event.Create(ms.WithServiceUser(), event)
}

func RaiseCriticalAlarm(alarmType string, text string) {
	alarm := c8y.Alarm {
		Time:   	c8y.NewTimestamp(),
		Type:   	alarmType,
		Severity: 	c8y.AlarmSeverityCritical,
		Text:   	text,
		Source: 	c8y.NewSource(ms.GetAgent().ID),
	}
	ms.Client.Alarm.Create(ms.WithServiceUser(), alarm)
}

func RaiseMajorAlarm(alarmType string, text string) {
	alarm := c8y.Alarm {
		Time:   	c8y.NewTimestamp(),
		Type:   	alarmType,
		Severity: 	c8y.AlarmSeverityMajor,
		Text:   	text,
		Source: 	c8y.NewSource(ms.GetAgent().ID),
	}
	ms.Client.Alarm.Create(ms.WithServiceUser(), alarm)
}