package mocks

import (
	"github.com/Drinnn/go-expert-events/pkg/events"
	"github.com/stretchr/testify/mock"
	"sync"
)

type EventHandler struct {
	ID   int
	Mock mock.Mock
}

func (eh *EventHandler) Handle(e events.EventInterface, wg *sync.WaitGroup) {
	eh.Mock.Called(e)
	wg.Done()
}
