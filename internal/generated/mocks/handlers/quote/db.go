// Code generated by mockery. DO NOT EDIT.

package quote

import (
	context "context"

	models "github.com/nijeti/cinema-keeper/internal/models"
	mock "github.com/stretchr/testify/mock"

	types "github.com/nijeti/cinema-keeper/internal/types"
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

// AddUserQuoteOnGuild provides a mock function with given fields: ctx, _a1
func (_m *MockDb) AddUserQuoteOnGuild(ctx context.Context, _a1 *models.Quote) error {
	ret := _m.Called(ctx, _a1)

	if len(ret) == 0 {
		panic("no return value specified for AddUserQuoteOnGuild")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.Quote) error); ok {
		r0 = rf(ctx, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDb_AddUserQuoteOnGuild_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddUserQuoteOnGuild'
type MockDb_AddUserQuoteOnGuild_Call struct {
	*mock.Call
}

// AddUserQuoteOnGuild is a helper method to define mock.On call
//   - ctx context.Context
//   - _a1 *models.Quote
func (_e *MockDb_Expecter) AddUserQuoteOnGuild(ctx interface{}, _a1 interface{}) *MockDb_AddUserQuoteOnGuild_Call {
	return &MockDb_AddUserQuoteOnGuild_Call{Call: _e.mock.On("AddUserQuoteOnGuild", ctx, _a1)}
}

func (_c *MockDb_AddUserQuoteOnGuild_Call) Run(run func(ctx context.Context, _a1 *models.Quote)) *MockDb_AddUserQuoteOnGuild_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.Quote))
	})
	return _c
}

func (_c *MockDb_AddUserQuoteOnGuild_Call) Return(_a0 error) *MockDb_AddUserQuoteOnGuild_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDb_AddUserQuoteOnGuild_Call) RunAndReturn(run func(context.Context, *models.Quote) error) *MockDb_AddUserQuoteOnGuild_Call {
	_c.Call.Return(run)
	return _c
}

// GetUserQuotesOnGuild provides a mock function with given fields: ctx, authorID, guildID
func (_m *MockDb) GetUserQuotesOnGuild(ctx context.Context, authorID types.ID, guildID types.ID) ([]*models.Quote, error) {
	ret := _m.Called(ctx, authorID, guildID)

	if len(ret) == 0 {
		panic("no return value specified for GetUserQuotesOnGuild")
	}

	var r0 []*models.Quote
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, types.ID, types.ID) ([]*models.Quote, error)); ok {
		return rf(ctx, authorID, guildID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, types.ID, types.ID) []*models.Quote); ok {
		r0 = rf(ctx, authorID, guildID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*models.Quote)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, types.ID, types.ID) error); ok {
		r1 = rf(ctx, authorID, guildID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDb_GetUserQuotesOnGuild_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetUserQuotesOnGuild'
type MockDb_GetUserQuotesOnGuild_Call struct {
	*mock.Call
}

// GetUserQuotesOnGuild is a helper method to define mock.On call
//   - ctx context.Context
//   - authorID types.ID
//   - guildID types.ID
func (_e *MockDb_Expecter) GetUserQuotesOnGuild(ctx interface{}, authorID interface{}, guildID interface{}) *MockDb_GetUserQuotesOnGuild_Call {
	return &MockDb_GetUserQuotesOnGuild_Call{Call: _e.mock.On("GetUserQuotesOnGuild", ctx, authorID, guildID)}
}

func (_c *MockDb_GetUserQuotesOnGuild_Call) Run(run func(ctx context.Context, authorID types.ID, guildID types.ID)) *MockDb_GetUserQuotesOnGuild_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(types.ID), args[2].(types.ID))
	})
	return _c
}

func (_c *MockDb_GetUserQuotesOnGuild_Call) Return(_a0 []*models.Quote, _a1 error) *MockDb_GetUserQuotesOnGuild_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDb_GetUserQuotesOnGuild_Call) RunAndReturn(run func(context.Context, types.ID, types.ID) ([]*models.Quote, error)) *MockDb_GetUserQuotesOnGuild_Call {
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
