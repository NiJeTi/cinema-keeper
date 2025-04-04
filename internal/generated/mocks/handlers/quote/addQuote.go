// Code generated by mockery. DO NOT EDIT.

package quote_test

import (
	context "context"

	discordgo "github.com/bwmarrin/discordgo"
	mock "github.com/stretchr/testify/mock"

	models "github.com/nijeti/cinema-keeper/internal/models"
)

// MockAddQuote is an autogenerated mock type for the addQuote type
type MockAddQuote struct {
	mock.Mock
}

type MockAddQuote_Expecter struct {
	mock *mock.Mock
}

func (_m *MockAddQuote) EXPECT() *MockAddQuote_Expecter {
	return &MockAddQuote_Expecter{mock: &_m.Mock}
}

// Exec provides a mock function with given fields: ctx, i, authorID, text
func (_m *MockAddQuote) Exec(ctx context.Context, i *discordgo.Interaction, authorID models.DiscordID, text string) error {
	ret := _m.Called(ctx, i, authorID, text)

	if len(ret) == 0 {
		panic("no return value specified for Exec")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *discordgo.Interaction, models.DiscordID, string) error); ok {
		r0 = rf(ctx, i, authorID, text)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockAddQuote_Exec_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Exec'
type MockAddQuote_Exec_Call struct {
	*mock.Call
}

// Exec is a helper method to define mock.On call
//   - ctx context.Context
//   - i *discordgo.Interaction
//   - authorID models.DiscordID
//   - text string
func (_e *MockAddQuote_Expecter) Exec(ctx interface{}, i interface{}, authorID interface{}, text interface{}) *MockAddQuote_Exec_Call {
	return &MockAddQuote_Exec_Call{Call: _e.mock.On("Exec", ctx, i, authorID, text)}
}

func (_c *MockAddQuote_Exec_Call) Run(run func(ctx context.Context, i *discordgo.Interaction, authorID models.DiscordID, text string)) *MockAddQuote_Exec_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*discordgo.Interaction), args[2].(models.DiscordID), args[3].(string))
	})
	return _c
}

func (_c *MockAddQuote_Exec_Call) Return(_a0 error) *MockAddQuote_Exec_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockAddQuote_Exec_Call) RunAndReturn(run func(context.Context, *discordgo.Interaction, models.DiscordID, string) error) *MockAddQuote_Exec_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockAddQuote creates a new instance of MockAddQuote. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockAddQuote(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockAddQuote {
	mock := &MockAddQuote{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
