package events

import "errors"

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

func (e *EventDispatcher) Clear() {
	e.Handlers = make(map[string][]EventHandlerInterface)
}
