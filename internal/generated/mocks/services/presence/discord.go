// Code generated by mockery. DO NOT EDIT.

package presence_test

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

// SetActivity provides a mock function with given fields: ctx, activity
func (_m *MockDiscord) SetActivity(ctx context.Context, activity *discordgo.Activity) error {
	ret := _m.Called(ctx, activity)

	if len(ret) == 0 {
		panic("no return value specified for SetActivity")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *discordgo.Activity) error); ok {
		r0 = rf(ctx, activity)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDiscord_SetActivity_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetActivity'
type MockDiscord_SetActivity_Call struct {
	*mock.Call
}

// SetActivity is a helper method to define mock.On call
//   - ctx context.Context
//   - activity *discordgo.Activity
func (_e *MockDiscord_Expecter) SetActivity(ctx interface{}, activity interface{}) *MockDiscord_SetActivity_Call {
	return &MockDiscord_SetActivity_Call{Call: _e.mock.On("SetActivity", ctx, activity)}
}

func (_c *MockDiscord_SetActivity_Call) Run(run func(ctx context.Context, activity *discordgo.Activity)) *MockDiscord_SetActivity_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*discordgo.Activity))
	})
	return _c
}

func (_c *MockDiscord_SetActivity_Call) Return(_a0 error) *MockDiscord_SetActivity_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDiscord_SetActivity_Call) RunAndReturn(run func(context.Context, *discordgo.Activity) error) *MockDiscord_SetActivity_Call {
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
