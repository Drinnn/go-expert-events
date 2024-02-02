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
	suite.handlerOne = mocks.EventHandler{ID: 1}
	suite.handlerTwo = mocks.EventHandler{ID: 2}
	suite.handlerThree = mocks.EventHandler{ID: 3}
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_Success() {
	err := suite.eventDispatcher.Register(suite.eventOne.GetName(), &suite.handlerOne)
	suite.NoError(err)
	suite.Equal(1, len(suite.eventDispatcher.Handlers[suite.eventOne.GetName()]))

	err = suite.eventDispatcher.Register(suite.eventOne.GetName(), &suite.handlerTwo)
	suite.NoError(err)
	suite.Equal(2, len(suite.eventDispatcher.Handlers[suite.eventOne.GetName()]))

	assert.Equal(suite.T(), &suite.handlerOne, suite.eventDispatcher.Handlers[suite.eventOne.GetName()][0])
	assert.Equal(suite.T(), &suite.handlerTwo, suite.eventDispatcher.Handlers[suite.eventOne.GetName()][1])
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}