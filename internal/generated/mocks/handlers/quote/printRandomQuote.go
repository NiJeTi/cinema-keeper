// Code generated by mockery. DO NOT EDIT.

package quote_test

import (
	context "context"

	discordgo "github.com/bwmarrin/discordgo"
	mock "github.com/stretchr/testify/mock"
)

// MockPrintRandomQuote is an autogenerated mock type for the printRandomQuote type
type MockPrintRandomQuote struct {
	mock.Mock
}

type MockPrintRandomQuote_Expecter struct {
	mock *mock.Mock
}

func (_m *MockPrintRandomQuote) EXPECT() *MockPrintRandomQuote_Expecter {
	return &MockPrintRandomQuote_Expecter{mock: &_m.Mock}
}

// Exec provides a mock function with given fields: ctx, i
func (_m *MockPrintRandomQuote) Exec(ctx context.Context, i *discordgo.Interaction) error {
	ret := _m.Called(ctx, i)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *discordgo.Interaction) error); ok {
		r0 = rf(ctx, i)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockPrintRandomQuote_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockPrintRandomQuote_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - i *discordgo.Interaction
func (_e *MockPrintRandomQuote_Expecter) Exec(ctx interface{}, i interface{}) *MockPrintRandomQuote_Exec_Call {
	return &MockPrintRandomQuote_Exec_Call{Call: _e.mock.On("Exec", ctx, i)}
}

func (_c *MockPrintRandomQuote_Exec_Call) Run(run func(ctx context.Context, i *discordgo.Interaction)) *MockPrintRandomQuote_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*discordgo.Interaction))
	})
	return _c
}

func (_c *MockPrintRandomQuote_Exec_Call) Return(_a0 error) *MockPrintRandomQuote_Exec_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockPrintRandomQuote_Exec_Call) RunAndReturn(run func(context.Context, *discordgo.Interaction) error) *MockPrintRandomQuote_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockPrintRandomQuote creates a new instance of MockPrintRandomQuote. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockPrintRandomQuote(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockPrintRandomQuote {
	mock := &MockPrintRandomQuote{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
