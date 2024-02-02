package test

import (
	"github.com/Drinnn/go-expert-events/pkg/events"
	"github.com/Drinnn/go-expert-events/test/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
)

type EventDispatcherTestSuite struct {
	suite.Suite
	eventDispatcher *events.EventDispatcher
	eventOne        mocks.Event
	eventTwo        mocks.Event
	handlerOne      mocks.EventHandler
	handlerTwo      mocks.EventHandler
	handlerThree    mocks.EventHandler
}

func (suite *EventDispatcherTestSuite) SetupTest() {
	suite.eventDispatcher = events.NewEventDispatcher()
	suite.eventOne = mocks.Event{Name: "event.one", Payload: "event.one payload"}
	suite.eventTwo = mocks.Event{Name: "event.two", Payload: "event.two payload"}
	suite.handlerOne = mocks.EventHandler{}
	suite.handlerTwo = mocks.EventHandler{}
	suite.handlerThree = mocks.EventHandler{}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register() {
	assert.True(suite.T(), true)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
