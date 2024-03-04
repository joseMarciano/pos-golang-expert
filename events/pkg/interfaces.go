package pkg

import "time"

type EventInterface interface {
	GetName() string
	GetDateTime() time.Time
	GetPayload() interface{} // any type
}

type EventHandlerInterface interface {
	Handle(event EventInterface)
}

type EventDispatcherInterface interface {
	Register(eventName string, handler EventHandlerInterface) error
	Unregister(eventName string, handler EventHandlerInterface) error
	HasListeners(eventName string, handler EventHandlerInterface) bool
	Clear() error
	Dispatch(event EventInterface) error
}
