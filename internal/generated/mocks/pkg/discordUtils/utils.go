// Code generated by mockery. DO NOT EDIT.

package discordutils_test

import (
	discordgo "github.com/bwmarrin/discordgo"

	mock "github.com/stretchr/testify/mock"

	types "github.com/nijeti/cinema-keeper/internal/types"
)

// MockUtils is an autogenerated mock type for the Utils type
type MockUtils struct {
	mock.Mock
}

type MockUtils_Expecter struct {
	mock *mock.Mock
}

func (_m *MockUtils) EXPECT() *MockUtils_Expecter {
	return &MockUtils_Expecter{mock: &_m.Mock}
}

// GetVoiceChannelUsers provides a mock function with given fields: guild, channel
func (_m *MockUtils) GetVoiceChannelUsers(guild types.ID, channel types.ID) ([]*discordgo.Member, error) {
	ret := _m.Called(guild, channel)

	if len(ret) == 0 {
		panic("no return value specified for GetVoiceChannelUsers")
	}

	var r0 []*discordgo.Member
	var r1 error
	if rf, ok := ret.Get(0).(func(types.ID, types.ID) ([]*discordgo.Member, error)); ok {
		return rf(guild, channel)
	}
	if rf, ok := ret.Get(0).(func(types.ID, types.ID) []*discordgo.Member); ok {
		r0 = rf(guild, channel)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*discordgo.Member)
		}
	}

	if rf, ok := ret.Get(1).(func(types.ID, types.ID) error); ok {
		r1 = rf(guild, channel)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockUtils_GetVoiceChannelUsers_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetVoiceChannelUsers'
type MockUtils_GetVoiceChannelUsers_Call struct {
	*mock.Call
}

// GetVoiceChannelUsers is a helper method to define mock.On call
//   - guild types.ID
//   - channel types.ID
func (_e *MockUtils_Expecter) GetVoiceChannelUsers(guild interface{}, channel interface{}) *MockUtils_GetVoiceChannelUsers_Call {
	return &MockUtils_GetVoiceChannelUsers_Call{Call: _e.mock.On("GetVoiceChannelUsers", guild, channel)}
}

func (_c *MockUtils_GetVoiceChannelUsers_Call) Run(run func(guild types.ID, channel types.ID)) *MockUtils_GetVoiceChannelUsers_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(types.ID), args[1].(types.ID))
	})
	return _c
}

func (_c *MockUtils_GetVoiceChannelUsers_Call) Return(_a0 []*discordgo.Member, _a1 error) *MockUtils_GetVoiceChannelUsers_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockUtils_GetVoiceChannelUsers_Call) RunAndReturn(run func(types.ID, types.ID) ([]*discordgo.Member, error)) *MockUtils_GetVoiceChannelUsers_Call {
	_c.Call.Return(run)
	return _c
}

// Respond provides a mock function with given fields: interaction, response
func (_m *MockUtils) Respond(interaction *discordgo.InteractionCreate, response *discordgo.InteractionResponse) error {
	ret := _m.Called(interaction, response)

	if len(ret) == 0 {
		panic("no return value specified for Respond")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*discordgo.InteractionCreate, *discordgo.InteractionResponse) error); ok {
		r0 = rf(interaction, response)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockUtils_Respond_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Respond'
type MockUtils_Respond_Call struct {
	*mock.Call
}

// Respond is a helper method to define mock.On call
//   - interaction *discordgo.InteractionCreate
//   - response *discordgo.InteractionResponse
func (_e *MockUtils_Expecter) Respond(interaction interface{}, response interface{}) *MockUtils_Respond_Call {
	return &MockUtils_Respond_Call{Call: _e.mock.On("Respond", interaction, response)}
}

func (_c *MockUtils_Respond_Call) Run(run func(interaction *discordgo.InteractionCreate, response *discordgo.InteractionResponse)) *MockUtils_Respond_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*discordgo.InteractionCreate), args[1].(*discordgo.InteractionResponse))
	})
	return _c
}

func (_c *MockUtils_Respond_Call) Return(_a0 error) *MockUtils_Respond_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockUtils_Respond_Call) RunAndReturn(run func(*discordgo.InteractionCreate, *discordgo.InteractionResponse) error) *MockUtils_Respond_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockUtils creates a new instance of MockUtils. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockUtils(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockUtils {
	mock := &MockUtils{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
