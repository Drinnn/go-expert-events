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

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Register_Failure_SameHandler() {
	err := suite.eventDispatcher.Register(suite.eventOne.GetName(), &suite.handlerOne)
	suite.NoError(err)
	suite.Equal(1, len(suite.eventDispatcher.Handlers[suite.eventOne.GetName()]))

	err = suite.eventDispatcher.Register(suite.eventOne.GetName(), &suite.handlerOne)
	suite.Error(events.ErrHandlerAlreadyRegistered, err)
	suite.Equal(1, len(suite.eventDispatcher.Handlers[suite.eventOne.GetName()]))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Clear_Success() {
	suite.eventDispatcher.Register(suite.eventOne.GetName(), &suite.handlerOne)
	suite.eventDispatcher.Register(suite.eventOne.GetName(), &suite.handlerTwo)
	suite.eventDispatcher.Register(suite.eventTwo.GetName(), &suite.handlerThree)

	suite.eventDispatcher.Clear()

	suite.Equal(0, len(suite.eventDispatcher.Handlers))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Has_Success() {
	suite.eventDispatcher.Register(suite.eventOne.GetName(), &suite.handlerOne)
	suite.eventDispatcher.Register(suite.eventOne.GetName(), &suite.handlerTwo)

	assert.True(suite.T(), suite.eventDispatcher.Has(suite.eventOne.GetName(), &suite.handlerOne))
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.eventOne.GetName(), &suite.handlerTwo))

	assert.False(suite.T(), suite.eventDispatcher.Has(suite.eventOne.GetName(), &suite.handlerThree))

	assert.False(suite.T(), suite.eventDispatcher.Has(suite.eventTwo.GetName(), &suite.handlerOne))
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Dispatch_Success() {
	suite.handlerOne.Mock.On("Handle", &suite.eventOne)

	suite.eventDispatcher.Register(suite.eventOne.GetName(), &suite.handlerOne)
	suite.eventDispatcher.Dispatch(&suite.eventOne)

	suite.handlerOne.Mock.AssertExpectations(suite.T())
	suite.handlerOne.Mock.AssertNumberOfCalls(suite.T(), "Handle", 1)
}

func (suite *EventDispatcherTestSuite) TestEventDispatcher_Remove_Success() {
	suite.eventDispatcher.Register(suite.eventOne.GetName(), &suite.handlerOne)
	suite.eventDispatcher.Register(suite.eventOne.GetName(), &suite.handlerTwo)

	suite.eventDispatcher.Remove(suite.eventOne.GetName(), &suite.handlerOne)

	suite.Equal(1, len(suite.eventDispatcher.Handlers[suite.eventOne.GetName()]))
	assert.True(suite.T(), suite.eventDispatcher.Has(suite.eventOne.GetName(), &suite.handlerTwo))
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(EventDispatcherTestSuite))
}
