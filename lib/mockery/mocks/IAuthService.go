// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	input "github.com/rodericusifo/fiber-template/internal/app/core/auth/service/dto/input"
	mock "github.com/stretchr/testify/mock"

	output "github.com/rodericusifo/fiber-template/internal/app/core/auth/service/dto/output"
)

// IAuthService is an autogenerated mock type for the IAuthService type
type IAuthService struct {
	mock.Mock
}

type IAuthService_Expecter struct {
	mock *mock.Mock
}

func (_m *IAuthService) EXPECT() *IAuthService_Expecter {
	return &IAuthService_Expecter{mock: &_m.Mock}
}

// LoginAuth provides a mock function with given fields: payload
func (_m *IAuthService) LoginAuth(payload *input.LoginAuthDTO) (*output.LoginAuthDTO, error) {
	ret := _m.Called(payload)

	if len(ret) == 0 {
		panic("no return value specified for LoginAuth")
	}

	var r0 *output.LoginAuthDTO
	var r1 error
	if rf, ok := ret.Get(0).(func(*input.LoginAuthDTO) (*output.LoginAuthDTO, error)); ok {
		return rf(payload)
	}
	if rf, ok := ret.Get(0).(func(*input.LoginAuthDTO) *output.LoginAuthDTO); ok {
		r0 = rf(payload)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*output.LoginAuthDTO)
		}
	}

	if rf, ok := ret.Get(1).(func(*input.LoginAuthDTO) error); ok {
		r1 = rf(payload)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IAuthService_LoginAuth_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoginAuth'
type IAuthService_LoginAuth_Call struct {
	*mock.Call
}

// LoginAuth is a helper method to define mock.On call
//   - payload *input.LoginAuthDTO
func (_e *IAuthService_Expecter) LoginAuth(payload interface{}) *IAuthService_LoginAuth_Call {
	return &IAuthService_LoginAuth_Call{Call: _e.mock.On("LoginAuth", payload)}
}

func (_c *IAuthService_LoginAuth_Call) Run(run func(payload *input.LoginAuthDTO)) *IAuthService_LoginAuth_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*input.LoginAuthDTO))
	})
	return _c
}

func (_c *IAuthService_LoginAuth_Call) Return(_a0 *output.LoginAuthDTO, _a1 error) *IAuthService_LoginAuth_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IAuthService_LoginAuth_Call) RunAndReturn(run func(*input.LoginAuthDTO) (*output.LoginAuthDTO, error)) *IAuthService_LoginAuth_Call {
	_c.Call.Return(run)
	return _c
}

// RegisterAuth provides a mock function with given fields: payload
func (_m *IAuthService) RegisterAuth(payload *input.RegisterAuthDTO) error {
	ret := _m.Called(payload)

	if len(ret) == 0 {
		panic("no return value specified for RegisterAuth")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*input.RegisterAuthDTO) error); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IAuthService_RegisterAuth_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RegisterAuth'
type IAuthService_RegisterAuth_Call struct {
	*mock.Call
}

// RegisterAuth is a helper method to define mock.On call
//   - payload *input.RegisterAuthDTO
func (_e *IAuthService_Expecter) RegisterAuth(payload interface{}) *IAuthService_RegisterAuth_Call {
	return &IAuthService_RegisterAuth_Call{Call: _e.mock.On("RegisterAuth", payload)}
}

func (_c *IAuthService_RegisterAuth_Call) Run(run func(payload *input.RegisterAuthDTO)) *IAuthService_RegisterAuth_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*input.RegisterAuthDTO))
	})
	return _c
}

func (_c *IAuthService_RegisterAuth_Call) Return(_a0 error) *IAuthService_RegisterAuth_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IAuthService_RegisterAuth_Call) RunAndReturn(run func(*input.RegisterAuthDTO) error) *IAuthService_RegisterAuth_Call {
	_c.Call.Return(run)
	return _c
}

// NewIAuthService creates a new instance of IAuthService. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIAuthService(t interface {
	mock.TestingT
	Cleanup(func())
}) *IAuthService {
	mock := &IAuthService{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
