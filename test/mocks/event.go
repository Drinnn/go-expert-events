package mocks

import "time"

type Event struct {
	Name    string
	Payload interface{}
}

func (e *Event) GetName() string {
	return e.Name
}

func (e *Event) GetPayload() interface{} {
	return e.Payload
}

func (e *Event) GetDateTime() time.Time {
	return time.Now()
}
