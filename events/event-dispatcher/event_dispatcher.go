package event_dispatcher

import (
	"errors"
	"github.com/devfullcycle/fcutils/pkg"
)

type EventDispatcher struct {
	handlers map[string][]pkg.EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]pkg.EventHandlerInterface),
	}
}

func (e *EventDispatcher) Register(eventName string, handler pkg.EventHandlerInterface) error {
	if _, ok := e.handlers[eventName]; ok {
		for _, h := range e.handlers[eventName] {
			if h == handler {
				return errors.New("handler " + eventName + " already registered")
			}
		}

	}

	e.handlers[eventName] = append(e.handlers[eventName], handler)

	return nil
}
