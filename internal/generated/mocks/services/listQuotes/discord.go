// Code generated by mockery. DO NOT EDIT.

package listQuotes_test

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

// GuildMember provides a mock function with given fields: ctx, guildID, userID
func (_m *MockDiscord) GuildMember(ctx context.Context, guildID models.ID, userID models.ID) (*discordgo.Member, error) {
	ret := _m.Called(ctx, guildID, userID)

	if len(ret) == 0 {
		panic("no return value specified for GuildMember")
	}

	var r0 *discordgo.Member
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, models.ID, models.ID) (*discordgo.Member, error)); ok {
		return rf(ctx, guildID, userID)
	}
	if rf, ok := ret.Get(0).(func(context.Context, models.ID, models.ID) *discordgo.Member); ok {
		r0 = rf(ctx, guildID, userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*discordgo.Member)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, models.ID, models.ID) error); ok {
		r1 = rf(ctx, guildID, userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockDiscord_GuildMember_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GuildMember'
type MockDiscord_GuildMember_Call struct {
	*mock.Call
}

// GuildMember is a helper method to define mock.On call
//   - ctx context.Context
//   - guildID models.ID
//   - userID models.ID
func (_e *MockDiscord_Expecter) GuildMember(ctx interface{}, guildID interface{}, userID interface{}) *MockDiscord_GuildMember_Call {
	return &MockDiscord_GuildMember_Call{Call: _e.mock.On("GuildMember", ctx, guildID, userID)}
}

func (_c *MockDiscord_GuildMember_Call) Run(run func(ctx context.Context, guildID models.ID, userID models.ID)) *MockDiscord_GuildMember_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.ID), args[2].(models.ID))
	})
	return _c
}

func (_c *MockDiscord_GuildMember_Call) Return(_a0 *discordgo.Member, _a1 error) *MockDiscord_GuildMember_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockDiscord_GuildMember_Call) RunAndReturn(run func(context.Context, models.ID, models.ID) (*discordgo.Member, error)) *MockDiscord_GuildMember_Call {
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

// SendEmbeds provides a mock function with given fields: ctx, channelID, embeds
func (_m *MockDiscord) SendEmbeds(ctx context.Context, channelID models.ID, embeds []*discordgo.MessageEmbed) error {
	ret := _m.Called(ctx, channelID, embeds)

	if len(ret) == 0 {
		panic("no return value specified for SendEmbeds")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.ID, []*discordgo.MessageEmbed) error); ok {
		r0 = rf(ctx, channelID, embeds)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockDiscord_SendEmbeds_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SendEmbeds'
type MockDiscord_SendEmbeds_Call struct {
	*mock.Call
}

// SendEmbeds is a helper method to define mock.On call
//   - ctx context.Context
//   - channelID models.ID
//   - embeds []*discordgo.MessageEmbed
func (_e *MockDiscord_Expecter) SendEmbeds(ctx interface{}, channelID interface{}, embeds interface{}) *MockDiscord_SendEmbeds_Call {
	return &MockDiscord_SendEmbeds_Call{Call: _e.mock.On("SendEmbeds", ctx, channelID, embeds)}
}

func (_c *MockDiscord_SendEmbeds_Call) Run(run func(ctx context.Context, channelID models.ID, embeds []*discordgo.MessageEmbed)) *MockDiscord_SendEmbeds_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(models.ID), args[2].([]*discordgo.MessageEmbed))
	})
	return _c
}

func (_c *MockDiscord_SendEmbeds_Call) Return(_a0 error) *MockDiscord_SendEmbeds_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockDiscord_SendEmbeds_Call) RunAndReturn(run func(context.Context, models.ID, []*discordgo.MessageEmbed) error) *MockDiscord_SendEmbeds_Call {
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
