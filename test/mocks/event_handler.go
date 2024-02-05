package mocks

import (
	"github.com/Drinnn/go-expert-events/pkg/events"
	"github.com/stretchr/testify/mock"
)

type EventHandler struct {
	ID   int
	Mock mock.Mock
}

func (eh *EventHandler) Handle(e events.EventInterface) {
	eh.Mock.Called(e)
}
