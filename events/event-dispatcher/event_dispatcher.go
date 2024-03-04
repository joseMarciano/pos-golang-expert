package event_dispatcher

import (
	"errors"
	"github.com/devfullcycle/fcutils/pkg"
	"sync"
)

type EventDispatcher struct {
	handlers map[string][]pkg.EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]pkg.EventHandlerInterface),
	}
}

func (e *EventDispatcher) Dispatch(event pkg.EventInterface) error {
	if handlers, ok := e.handlers[event.GetName()]; ok {
		wg := sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, &wg)
		}
		wg.Wait() // wait here until all handlers done
	}

	return nil
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

func (e *EventDispatcher) Clear() error {
	e.handlers = make(map[string][]pkg.EventHandlerInterface)
	return nil
}

func (e *EventDispatcher) Has(eventName string, handler pkg.EventHandlerInterface) bool {
	if _, ok := e.handlers[eventName]; ok {
		for _, h := range e.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}

	return false
}

func (e *EventDispatcher) Remove(eventName string, handler pkg.EventHandlerInterface) error {
	if _, ok := e.handlers[eventName]; ok {
		for i, h := range e.handlers[eventName] {
			if h == handler {
				e.handlers[eventName] = append(e.handlers[eventName][:i], e.handlers[eventName][i+1:]...)
				return nil
			}
		}
	}

	return nil

}
