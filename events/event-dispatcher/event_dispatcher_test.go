package event_dispatcher

import (
	"github.com/devfullcycle/fcutils/pkg"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"sync"
	"testing"
	"time"
)

type TestEvent struct {
	Name    string
	Payload interface{}
}

func (te *TestEvent) GetName() string {
	return te.Name
}
func (te *TestEvent) GetPayload() interface{} {
	return te.Payload
}

func (te *TestEvent) GetDateTime() time.Time {
	return time.Now()
}

type TestEventHandler struct {
	ID int
}

func (th *TestEventHandler) Handle(event pkg.EventInterface, s *sync.WaitGroup) {
}

type EventDispatcherTestSuit struct {
	suite.Suite     // creating a suite of tests using testify
	event           TestEvent
	event2          TestEvent
	handler         TestEventHandler
	handler2        TestEventHandler
	handler3        TestEventHandler
	eventDispatcher EventDispatcher
}

func (suite *EventDispatcherTestSuit) SetupTest() { // run before each test
	suite.event = TestEvent{
		Name:    "test",
		Payload: "PAYLOAD",
	}
	suite.event2 = TestEvent{
		Name:    "test2",
		Payload: "PAYLOAD2",
	}
	suite.handler = TestEventHandler{
		ID: 0,
	}
	suite.handler2 = TestEventHandler{
		ID: 1,
	}
	suite.handler3 = TestEventHandler{
		ID: 2,
	}
	suite.eventDispatcher = EventDispatcher{
		handlers: make(map[string][]pkg.EventHandlerInterface),
	}
}

func (suite *EventDispatcherTestSuit) TestEventDispatcher_Register_ShouldAddHandler() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	assert.Equal(suite.T(), &suite.handler, suite.eventDispatcher.handlers[suite.event.GetName()][0])
	assert.Equal(suite.T(), &suite.handler2, suite.eventDispatcher.handlers[suite.event.GetName()][1])
}
func (suite *EventDispatcherTestSuit) TestEventDispatcher_Register_ShouldReturnError() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.NotNil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
}

func (suite *EventDispatcherTestSuit) TestClear_ShouldClear() {

	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	suite.eventDispatcher.Clear()
	suite.Equal(0, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	err = suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))
}

func (suite *EventDispatcherTestSuit) TestHas_ShouldWorkFine() {
	err := suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.Nil(err)
	suite.Equal(1, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

	hasEvent := suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler)
	assert.True(suite.T(), hasEvent)

	suite.eventDispatcher.Clear()
	hasEvent = suite.eventDispatcher.Has(suite.event.GetName(), &suite.handler)
	assert.False(suite.T(), hasEvent)
}

type MockHandler struct {
	mock.Mock
}

func (m *MockHandler) Handle(event pkg.EventInterface, wg *sync.WaitGroup) {
	m.Called(event)
	wg.Done()
}

func (suite *EventDispatcherTestSuit) TestDispatch_ShouldDispatch() {
	eventHandler := &MockHandler{}
	eventHandler.On("Handle", &suite.event)
	suite.eventDispatcher.Register(suite.event.GetName(), eventHandler)
	suite.eventDispatcher.Dispatch(&suite.event)
	eventHandler.AssertExpectations(suite.T())
	eventHandler.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func (suite *EventDispatcherTestSuit) TestRemove() {
	suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler)
	suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler2)
	suite.eventDispatcher.Register(suite.event.GetName(), &suite.handler3)

	suite.eventDispatcher.Remove(suite.event.GetName(), &suite.handler2)
	suite.Equal(2, len(suite.eventDispatcher.handlers[suite.event.GetName()]))

}

func TestSuite(t *testing.T) { // run this and all test inside the suite will run
	suite.Run(t, new(EventDispatcherTestSuit))
}
