package cumulocity

import (
	"github.com/reubenmiller/go-c8y/pkg/c8y"
	"github.com/reubenmiller/go-c8y/pkg/microservice"
	"os"
	"path"
)

var (
	osRemoveAll = os.RemoveAll
)

type Microservice struct {
	ms *microservice.Microservice
}

func Init() Microservice {
	ms := microservice.NewDefaultMicroservice(microservice.Options{})
	ms.Config.SetDefault("application.name", "c8y-scanner-clamav")
	ms.Config.SetDefault("agent.identityType", "c8y-scanner")
	ms.RegisterMicroserviceAgent()
	return Microservice{ms: ms}
}

func (m *Microservice) DownloadFile(file_id string) string {
	filename, _ := m.ms.Client.Inventory.DownloadBinary(m.ms.WithServiceUser(), file_id)
	return filename
}

func (m *Microservice) CleanupFile(filename string) {
	dir := path.Dir(filename)
	osRemoveAll(dir)
}

func (m *Microservice) RaiseEvent(eventType string, text string) {
	event := c8y.Event{
		Time:   c8y.NewTimestamp(),
		Type:   eventType,
		Text:   text,
		Source: c8y.NewSource(m.ms.GetAgent().ID),
	}
	m.ms.Client.Event.Create(m.ms.WithServiceUser(), event)
}

func (m *Microservice) RaiseCriticalAlarm(alarmType string, text string) {
	alarm := c8y.Alarm{
		Time:     c8y.NewTimestamp(),
		Type:     alarmType,
		Severity: c8y.AlarmSeverityCritical,
		Text:     text,
		Source:   c8y.NewSource(m.ms.GetAgent().ID),
	}
	m.ms.Client.Alarm.Create(m.ms.WithServiceUser(), alarm)
}

func (m *Microservice) RaiseMajorAlarm(alarmType string, text string) {
	alarm := c8y.Alarm{
		Time:     c8y.NewTimestamp(),
		Type:     alarmType,
		Severity: c8y.AlarmSeverityMajor,
		Text:     text,
		Source:   c8y.NewSource(m.ms.GetAgent().ID),
	}
	m.ms.Client.Alarm.Create(m.ms.WithServiceUser(), alarm)
}
