// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"

// IRequestBody is an autogenerated mock type for the IRequestBody type
type IRequestBody struct {
	mock.Mock
}

type IRequestBody_Expecter struct {
	mock *mock.Mock
}

func (_m *IRequestBody) EXPECT() *IRequestBody_Expecter {
	return &IRequestBody_Expecter{mock: &_m.Mock}
}

// CustomValidateRequestBody provides a mock function with no fields
func (_m *IRequestBody) CustomValidateRequestBody() error {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for CustomValidateRequestBody")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IRequestBody_CustomValidateRequestBody_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CustomValidateRequestBody'
type IRequestBody_CustomValidateRequestBody_Call struct {
	*mock.Call
}

// CustomValidateRequestBody is a helper method to define mock.On call
func (_e *IRequestBody_Expecter) CustomValidateRequestBody() *IRequestBody_CustomValidateRequestBody_Call {
	return &IRequestBody_CustomValidateRequestBody_Call{Call: _e.mock.On("CustomValidateRequestBody")}
}

func (_c *IRequestBody_CustomValidateRequestBody_Call) Run(run func()) *IRequestBody_CustomValidateRequestBody_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *IRequestBody_CustomValidateRequestBody_Call) Return(_a0 error) *IRequestBody_CustomValidateRequestBody_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IRequestBody_CustomValidateRequestBody_Call) RunAndReturn(run func() error) *IRequestBody_CustomValidateRequestBody_Call {
	_c.Call.Return(run)
	return _c
}

// NewIRequestBody creates a new instance of IRequestBody. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIRequestBody(t interface {
	mock.TestingT
	Cleanup(func())
}) *IRequestBody {
	mock := &IRequestBody{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
