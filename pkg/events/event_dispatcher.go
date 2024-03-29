package events

import (
	"errors"
	"sync"
)

var ErrHandlerAlreadyRegistered = errors.New("handler already registered")

type EventDispatcher struct {
	Handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		Handlers: make(map[string][]EventHandlerInterface),
	}
}

func (e *EventDispatcher) Register(eventName string, handler EventHandlerInterface) error {
	if _, ok := e.Handlers[eventName]; ok {
		for _, registeredHandler := range e.Handlers[eventName] {
			if registeredHandler == handler {
				return ErrHandlerAlreadyRegistered
			}
		}
	}
	e.Handlers[eventName] = append(e.Handlers[eventName], handler)
	return nil
}

func (e *EventDispatcher) Dispatch(event EventInterface) {
	if _, ok := e.Handlers[event.GetName()]; ok {
		waitGroup := &sync.WaitGroup{}
		for _, handler := range e.Handlers[event.GetName()] {
			waitGroup.Add(1)
			go handler.Handle(event, waitGroup)
		}
		waitGroup.Wait()
	}
}

func (e *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) {
	if _, ok := e.Handlers[eventName]; ok {
		for i, registeredHandler := range e.Handlers[eventName] {
			if registeredHandler == handler {
				e.Handlers[eventName] = append(e.Handlers[eventName][:i], e.Handlers[eventName][i+1:]...)
			}
		}
	}
}

func (e *EventDispatcher) Has(eventName string, handler EventHandlerInterface) bool {
	if _, ok := e.Handlers[eventName]; ok {
		for _, registeredHandler := range e.Handlers[eventName] {
			if registeredHandler == handler {
				return true
			}
		}
	}

	return false
}

func (e *EventDispatcher) Clear() {
	e.Handlers = make(map[string][]EventHandlerInterface)
}
