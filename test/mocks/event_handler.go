package mocks

import "github.com/Drinnn/go-expert-events/pkg/events"

type EventHandler struct {
	ID int
}

func (eh *EventHandler) Handle(e events.EventInterface) {

}
