package mocks

import "github.com/Drinnn/go-expert-events/pkg/events"

type EventHandler struct{}

func (eh *EventHandler) HandleEvent(e events.EventInterface) {

}
