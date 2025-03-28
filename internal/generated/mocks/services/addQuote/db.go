// Code generated by mockery. DO NOT EDIT.

package addQuote_test

import (
	context "context"

	models "github.com/nijeti/cinema-keeper/internal/models"
	mock "github.com/stretchr/testify/mock"
)

// MockDb is an autogenerated mock type for the db type
type MockDb struct {
	mock.Mock
}

type MockDb_Expecter struct {
	mock *mock.Mock
}

func (_m *MockDb) EXPECT() *MockDb_Expecter {
	return &MockDb_Expecter{mock: &_m.Mock}
}

// AddUserQuoteInGuild provides a mock function with given fields: ctx, quote
func (_m *MockDb) AddUserQuoteInGuild(ctx context.Context, quote *models.Quote) error {
	ret := _m.Called(ctx, quote)

	if len(ret) == 0 {
		panic("no return value specified for AddUserQuoteInGuild")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Quote) error); ok {
		r0 = rf(ctx, quote)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDb_AddUserQuoteInGuild_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUserQuoteInGuild'
type MockDb_AddUserQuoteInGuild_Call struct {
	*mock.Call
}

// AddUserQuoteInGuild is a helper method to define mock.On call
//   - ctx context.Context
//   - quote *models.Quote
func (_e *MockDb_Expecter) AddUserQuoteInGuild(ctx interface{}, quote interface{}) *MockDb_AddUserQuoteInGuild_Call {
	return &MockDb_AddUserQuoteInGuild_Call{Call: _e.mock.On("AddUserQuoteInGuild", ctx, quote)}
}

func (_c *MockDb_AddUserQuoteInGuild_Call) Run(run func(ctx context.Context, quote *models.Quote)) *MockDb_AddUserQuoteInGuild_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Quote))
	})
	return _c
}

func (_c *MockDb_AddUserQuoteInGuild_Call) Return(_a0 error) *MockDb_AddUserQuoteInGuild_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDb_AddUserQuoteInGuild_Call) RunAndReturn(run func(context.Context, *models.Quote) error) *MockDb_AddUserQuoteInGuild_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockDb creates a new instance of MockDb. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockDb(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockDb {
	mock := &MockDb{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
