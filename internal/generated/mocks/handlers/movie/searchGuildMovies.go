// Code generated by mockery. DO NOT EDIT.

package movie_test

import (
	context "context"

	discordgo "github.com/bwmarrin/discordgo"
	mock "github.com/stretchr/testify/mock"
)

// MockSearchGuildMovies is an autogenerated mock type for the searchGuildMovies type
type MockSearchGuildMovies struct {
	mock.Mock
}

type MockSearchGuildMovies_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSearchGuildMovies) EXPECT() *MockSearchGuildMovies_Expecter {
	return &MockSearchGuildMovies_Expecter{mock: &_m.Mock}
}

// Exec provides a mock function with given fields: ctx, i, title
func (_m *MockSearchGuildMovies) Exec(ctx context.Context, i *discordgo.Interaction, title string) error {
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

// MockSearchGuildMovies_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockSearchGuildMovies_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - i *discordgo.Interaction
//   - title string
func (_e *MockSearchGuildMovies_Expecter) Exec(ctx interface{}, i interface{}, title interface{}) *MockSearchGuildMovies_Exec_Call {
	return &MockSearchGuildMovies_Exec_Call{Call: _e.mock.On("Exec", ctx, i, title)}
}

func (_c *MockSearchGuildMovies_Exec_Call) Run(run func(ctx context.Context, i *discordgo.Interaction, title string)) *MockSearchGuildMovies_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*discordgo.Interaction), args[2].(string))
	})
	return _c
}

func (_c *MockSearchGuildMovies_Exec_Call) Return(_a0 error) *MockSearchGuildMovies_Exec_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSearchGuildMovies_Exec_Call) RunAndReturn(run func(context.Context, *discordgo.Interaction, string) error) *MockSearchGuildMovies_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSearchGuildMovies creates a new instance of MockSearchGuildMovies. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSearchGuildMovies(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSearchGuildMovies {
	mock := &MockSearchGuildMovies{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
