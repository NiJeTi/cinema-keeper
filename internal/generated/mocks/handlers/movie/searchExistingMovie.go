// Code generated by mockery. DO NOT EDIT.

package movie_test

import (
	context "context"

	discordgo "github.com/bwmarrin/discordgo"
	mock "github.com/stretchr/testify/mock"
)

// MockSearchExistingMovie is an autogenerated mock type for the searchExistingMovie type
type MockSearchExistingMovie struct {
	mock.Mock
}

type MockSearchExistingMovie_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSearchExistingMovie) EXPECT() *MockSearchExistingMovie_Expecter {
	return &MockSearchExistingMovie_Expecter{mock: &_m.Mock}
}

// Exec provides a mock function with given fields: ctx, i, title
func (_m *MockSearchExistingMovie) Exec(ctx context.Context, i *discordgo.Interaction, title string) error {
	ret := _m.Called(ctx, i, title)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *discordgo.Interaction, string) error); ok {
		r0 = rf(ctx, i, title)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSearchExistingMovie_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockSearchExistingMovie_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - i *discordgo.Interaction
//   - title string
func (_e *MockSearchExistingMovie_Expecter) Exec(ctx interface{}, i interface{}, title interface{}) *MockSearchExistingMovie_Exec_Call {
	return &MockSearchExistingMovie_Exec_Call{Call: _e.mock.On("Exec", ctx, i, title)}
}

func (_c *MockSearchExistingMovie_Exec_Call) Run(run func(ctx context.Context, i *discordgo.Interaction, title string)) *MockSearchExistingMovie_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*discordgo.Interaction), args[2].(string))
	})
	return _c
}

func (_c *MockSearchExistingMovie_Exec_Call) Return(_a0 error) *MockSearchExistingMovie_Exec_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSearchExistingMovie_Exec_Call) RunAndReturn(run func(context.Context, *discordgo.Interaction, string) error) *MockSearchExistingMovie_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSearchExistingMovie creates a new instance of MockSearchExistingMovie. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSearchExistingMovie(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSearchExistingMovie {
	mock := &MockSearchExistingMovie{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}