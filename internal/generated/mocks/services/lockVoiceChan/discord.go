// Code generated by mockery. DO NOT EDIT.

package lockVoiceChan_test

import (
	context "context"

	discordgo "github.com/bwmarrin/discordgo"

	mock "github.com/stretchr/testify/mock"

	models "github.com/nijeti/cinema-keeper/internal/models"
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

// EditChannel provides a mock function with given fields: ctx, channelID, edit
func (_m *MockDiscord) EditChannel(ctx context.Context, channelID models.DiscordID, edit *discordgo.ChannelEdit) error {
	ret := _m.Called(ctx, channelID, edit)

	if len(ret) == 0 {
		panic("no return value specified for EditChannel")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.DiscordID, *discordgo.ChannelEdit) error); ok {
		r0 = rf(ctx, channelID, edit)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDiscord_EditChannel_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'EditChannel'
type MockDiscord_EditChannel_Call struct {
	*mock.Call
}

// EditChannel is a helper method to define mock.On call
//   - ctx context.Context
//   - channelID models.DiscordID
//   - edit *discordgo.ChannelEdit
func (_e *MockDiscord_Expecter) EditChannel(ctx interface{}, channelID interface{}, edit interface{}) *MockDiscord_EditChannel_Call {
	return &MockDiscord_EditChannel_Call{Call: _e.mock.On("EditChannel", ctx, channelID, edit)}
}

func (_c *MockDiscord_EditChannel_Call) Run(run func(ctx context.Context, channelID models.DiscordID, edit *discordgo.ChannelEdit)) *MockDiscord_EditChannel_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.DiscordID), args[2].(*discordgo.ChannelEdit))
	})
	return _c
}

func (_c *MockDiscord_EditChannel_Call) Return(_a0 error) *MockDiscord_EditChannel_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDiscord_EditChannel_Call) RunAndReturn(run func(context.Context, models.DiscordID, *discordgo.ChannelEdit) error) *MockDiscord_EditChannel_Call {
	_c.Call.Return(run)
	return _c
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

// UserVoiceState provides a mock function with given fields: ctx, guildID, userID
func (_m *MockDiscord) UserVoiceState(ctx context.Context, guildID models.DiscordID, userID models.DiscordID) (*discordgo.VoiceState, error) {
	ret := _m.Called(ctx, guildID, userID)

	if len(ret) == 0 {
		panic("no return value specified for UserVoiceState")
	}

	var r0 *discordgo.VoiceState
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.DiscordID, models.DiscordID) (*discordgo.VoiceState, error)); ok {
		return rf(ctx, guildID, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.DiscordID, models.DiscordID) *discordgo.VoiceState); ok {
		r0 = rf(ctx, guildID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discordgo.VoiceState)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.DiscordID, models.DiscordID) error); ok {
		r1 = rf(ctx, guildID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDiscord_UserVoiceState_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UserVoiceState'
type MockDiscord_UserVoiceState_Call struct {
	*mock.Call
}

// UserVoiceState is a helper method to define mock.On call
//   - ctx context.Context
//   - guildID models.DiscordID
//   - userID models.DiscordID
func (_e *MockDiscord_Expecter) UserVoiceState(ctx interface{}, guildID interface{}, userID interface{}) *MockDiscord_UserVoiceState_Call {
	return &MockDiscord_UserVoiceState_Call{Call: _e.mock.On("UserVoiceState", ctx, guildID, userID)}
}

func (_c *MockDiscord_UserVoiceState_Call) Run(run func(ctx context.Context, guildID models.DiscordID, userID models.DiscordID)) *MockDiscord_UserVoiceState_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.DiscordID), args[2].(models.DiscordID))
	})
	return _c
}

func (_c *MockDiscord_UserVoiceState_Call) Return(_a0 *discordgo.VoiceState, _a1 error) *MockDiscord_UserVoiceState_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDiscord_UserVoiceState_Call) RunAndReturn(run func(context.Context, models.DiscordID, models.DiscordID) (*discordgo.VoiceState, error)) *MockDiscord_UserVoiceState_Call {
	_c.Call.Return(run)
	return _c
}

// VoiceChannelUsers provides a mock function with given fields: ctx, guildID, channelID
func (_m *MockDiscord) VoiceChannelUsers(ctx context.Context, guildID models.DiscordID, channelID models.DiscordID) ([]*discordgo.Member, error) {
	ret := _m.Called(ctx, guildID, channelID)

	if len(ret) == 0 {
		panic("no return value specified for VoiceChannelUsers")
	}

	var r0 []*discordgo.Member
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.DiscordID, models.DiscordID) ([]*discordgo.Member, error)); ok {
		return rf(ctx, guildID, channelID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.DiscordID, models.DiscordID) []*discordgo.Member); ok {
		r0 = rf(ctx, guildID, channelID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*discordgo.Member)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.DiscordID, models.DiscordID) error); ok {
		r1 = rf(ctx, guildID, channelID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDiscord_VoiceChannelUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'VoiceChannelUsers'
type MockDiscord_VoiceChannelUsers_Call struct {
	*mock.Call
}

// VoiceChannelUsers is a helper method to define mock.On call
//   - ctx context.Context
//   - guildID models.DiscordID
//   - channelID models.DiscordID
func (_e *MockDiscord_Expecter) VoiceChannelUsers(ctx interface{}, guildID interface{}, channelID interface{}) *MockDiscord_VoiceChannelUsers_Call {
	return &MockDiscord_VoiceChannelUsers_Call{Call: _e.mock.On("VoiceChannelUsers", ctx, guildID, channelID)}
}

func (_c *MockDiscord_VoiceChannelUsers_Call) Run(run func(ctx context.Context, guildID models.DiscordID, channelID models.DiscordID)) *MockDiscord_VoiceChannelUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.DiscordID), args[2].(models.DiscordID))
	})
	return _c
}

func (_c *MockDiscord_VoiceChannelUsers_Call) Return(_a0 []*discordgo.Member, _a1 error) *MockDiscord_VoiceChannelUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDiscord_VoiceChannelUsers_Call) RunAndReturn(run func(context.Context, models.DiscordID, models.DiscordID) ([]*discordgo.Member, error)) *MockDiscord_VoiceChannelUsers_Call {
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
