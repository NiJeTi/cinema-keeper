// Code generated by mockery. DO NOT EDIT.

package diceRoll_test

import (
	context "context"

	discordgo "github.com/bwmarrin/discordgo"

	mock "github.com/stretchr/testify/mock"
)

// MockDiscord is an autogenerated mock type for the discord type
type MockDiscord struct {
	mock.Mock
}

type MockDiscord_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDiscord) EXPECT() *MockDiscord_Expecter {
	return &MockDiscord_Expecter{mock: &_m.Mock}
}

// Respond provides a mock function with given fields: ctx, i, response
func (_m *MockDiscord) Respond(ctx context.Context, i *discordgo.Interaction, response *discordgo.InteractionResponse) error {
	ret := _m.Called(ctx, i, response)

	if len(ret) == 0 {
		panic("no return value specified for Respond")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *discordgo.Interaction, *discordgo.InteractionResponse) error); ok {
		r0 = rf(ctx, i, response)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDiscord_Respond_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Respond'
type MockDiscord_Respond_Call struct {
	*mock.Call
}

// Respond is a helper method to define mock.On call
//   - ctx context.Context
//   - i *discordgo.Interaction
//   - response *discordgo.InteractionResponse
func (_e *MockDiscord_Expecter) Respond(ctx interface{}, i interface{}, response interface{}) *MockDiscord_Respond_Call {
	return &MockDiscord_Respond_Call{Call: _e.mock.On("Respond", ctx, i, response)}
}

func (_c *MockDiscord_Respond_Call) Run(run func(ctx context.Context, i *discordgo.Interaction, response *discordgo.InteractionResponse)) *MockDiscord_Respond_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*discordgo.Interaction), args[2].(*discordgo.InteractionResponse))
	})
	return _c
}

func (_c *MockDiscord_Respond_Call) Return(_a0 error) *MockDiscord_Respond_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDiscord_Respond_Call) RunAndReturn(run func(context.Context, *discordgo.Interaction, *discordgo.InteractionResponse) error) *MockDiscord_Respond_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDiscord creates a new instance of MockDiscord. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDiscord(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDiscord {
	mock := &MockDiscord{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
