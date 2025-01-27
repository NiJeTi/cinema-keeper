// Code generated by mockery. DO NOT EDIT.

package addMovie_test

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

// AddMovie provides a mock function with given fields: ctx, movie
func (_m *MockDb) AddMovie(ctx context.Context, movie *models.MovieMeta) (models.ID, error) {
	ret := _m.Called(ctx, movie)

	if len(ret) == 0 {
		panic("no return value specified for AddMovie")
	}

	var r0 models.ID
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *models.MovieMeta) (models.ID, error)); ok {
		return rf(ctx, movie)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *models.MovieMeta) models.ID); ok {
		r0 = rf(ctx, movie)
	} else {
		r0 = ret.Get(0).(models.ID)
	}

	if rf, ok := ret.Get(1).(func(context.Context, *models.MovieMeta) error); ok {
		r1 = rf(ctx, movie)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDb_AddMovie_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddMovie'
type MockDb_AddMovie_Call struct {
	*mock.Call
}

// AddMovie is a helper method to define mock.On call
//   - ctx context.Context
//   - movie *models.MovieMeta
func (_e *MockDb_Expecter) AddMovie(ctx interface{}, movie interface{}) *MockDb_AddMovie_Call {
	return &MockDb_AddMovie_Call{Call: _e.mock.On("AddMovie", ctx, movie)}
}

func (_c *MockDb_AddMovie_Call) Run(run func(ctx context.Context, movie *models.MovieMeta)) *MockDb_AddMovie_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(*models.MovieMeta))
	})
	return _c
}

func (_c *MockDb_AddMovie_Call) Return(_a0 models.ID, _a1 error) *MockDb_AddMovie_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDb_AddMovie_Call) RunAndReturn(run func(context.Context, *models.MovieMeta) (models.ID, error)) *MockDb_AddMovie_Call {
	_c.Call.Return(run)
	return _c
}

// AddMovieToGuild provides a mock function with given fields: ctx, movieID, guildID, addedByID
func (_m *MockDb) AddMovieToGuild(ctx context.Context, movieID models.ID, guildID models.DiscordID, addedByID models.DiscordID) error {
	ret := _m.Called(ctx, movieID, guildID, addedByID)

	if len(ret) == 0 {
		panic("no return value specified for AddMovieToGuild")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.ID, models.DiscordID, models.DiscordID) error); ok {
		r0 = rf(ctx, movieID, guildID, addedByID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDb_AddMovieToGuild_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'AddMovieToGuild'
type MockDb_AddMovieToGuild_Call struct {
	*mock.Call
}

// AddMovieToGuild is a helper method to define mock.On call
//   - ctx context.Context
//   - movieID models.ID
//   - guildID models.DiscordID
//   - addedByID models.DiscordID
func (_e *MockDb_Expecter) AddMovieToGuild(ctx interface{}, movieID interface{}, guildID interface{}, addedByID interface{}) *MockDb_AddMovieToGuild_Call {
	return &MockDb_AddMovieToGuild_Call{Call: _e.mock.On("AddMovieToGuild", ctx, movieID, guildID, addedByID)}
}

func (_c *MockDb_AddMovieToGuild_Call) Run(run func(ctx context.Context, movieID models.ID, guildID models.DiscordID, addedByID models.DiscordID)) *MockDb_AddMovieToGuild_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.ID), args[2].(models.DiscordID), args[3].(models.DiscordID))
	})
	return _c
}

func (_c *MockDb_AddMovieToGuild_Call) Return(_a0 error) *MockDb_AddMovieToGuild_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDb_AddMovieToGuild_Call) RunAndReturn(run func(context.Context, models.ID, models.DiscordID, models.DiscordID) error) *MockDb_AddMovieToGuild_Call {
	_c.Call.Return(run)
	return _c
}

// GuildMovieExists provides a mock function with given fields: ctx, id, guildID
func (_m *MockDb) GuildMovieExists(ctx context.Context, id models.ID, guildID models.DiscordID) (bool, error) {
	ret := _m.Called(ctx, id, guildID)

	if len(ret) == 0 {
		panic("no return value specified for GuildMovieExists")
	}

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.ID, models.DiscordID) (bool, error)); ok {
		return rf(ctx, id, guildID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.ID, models.DiscordID) bool); ok {
		r0 = rf(ctx, id, guildID)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.ID, models.DiscordID) error); ok {
		r1 = rf(ctx, id, guildID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDb_GuildMovieExists_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GuildMovieExists'
type MockDb_GuildMovieExists_Call struct {
	*mock.Call
}

// GuildMovieExists is a helper method to define mock.On call
//   - ctx context.Context
//   - id models.ID
//   - guildID models.DiscordID
func (_e *MockDb_Expecter) GuildMovieExists(ctx interface{}, id interface{}, guildID interface{}) *MockDb_GuildMovieExists_Call {
	return &MockDb_GuildMovieExists_Call{Call: _e.mock.On("GuildMovieExists", ctx, id, guildID)}
}

func (_c *MockDb_GuildMovieExists_Call) Run(run func(ctx context.Context, id models.ID, guildID models.DiscordID)) *MockDb_GuildMovieExists_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.ID), args[2].(models.DiscordID))
	})
	return _c
}

func (_c *MockDb_GuildMovieExists_Call) Return(_a0 bool, _a1 error) *MockDb_GuildMovieExists_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDb_GuildMovieExists_Call) RunAndReturn(run func(context.Context, models.ID, models.DiscordID) (bool, error)) *MockDb_GuildMovieExists_Call {
	_c.Call.Return(run)
	return _c
}

// MovieByImdbID provides a mock function with given fields: ctx, id
func (_m *MockDb) MovieByImdbID(ctx context.Context, id models.ImdbID) (*models.ID, *models.MovieMeta, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for MovieByImdbID")
	}

	var r0 *models.ID
	var r1 *models.MovieMeta
	var r2 error
	if rf, ok := ret.Get(0).(func(context.Context, models.ImdbID) (*models.ID, *models.MovieMeta, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.ImdbID) *models.ID); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.ID)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.ImdbID) *models.MovieMeta); ok {
		r1 = rf(ctx, id)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*models.MovieMeta)
		}
	}

	if rf, ok := ret.Get(2).(func(context.Context, models.ImdbID) error); ok {
		r2 = rf(ctx, id)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// MockDb_MovieByImdbID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'MovieByImdbID'
type MockDb_MovieByImdbID_Call struct {
	*mock.Call
}

// MovieByImdbID is a helper method to define mock.On call
//   - ctx context.Context
//   - id models.ImdbID
func (_e *MockDb_Expecter) MovieByImdbID(ctx interface{}, id interface{}) *MockDb_MovieByImdbID_Call {
	return &MockDb_MovieByImdbID_Call{Call: _e.mock.On("MovieByImdbID", ctx, id)}
}

func (_c *MockDb_MovieByImdbID_Call) Run(run func(ctx context.Context, id models.ImdbID)) *MockDb_MovieByImdbID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.ImdbID))
	})
	return _c
}

func (_c *MockDb_MovieByImdbID_Call) Return(_a0 *models.ID, _a1 *models.MovieMeta, _a2 error) *MockDb_MovieByImdbID_Call {
	_c.Call.Return(_a0, _a1, _a2)
	return _c
}

func (_c *MockDb_MovieByImdbID_Call) RunAndReturn(run func(context.Context, models.ImdbID) (*models.ID, *models.MovieMeta, error)) *MockDb_MovieByImdbID_Call {
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
