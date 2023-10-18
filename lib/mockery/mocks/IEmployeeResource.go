// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	sql "github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	types "github.com/rodericusifo/fiber-template/pkg/types"
)

// IEmployeeResource is an autogenerated mock type for the IEmployeeResource type
type IEmployeeResource struct {
	mock.Mock
}

type IEmployeeResource_Expecter struct {
	mock *mock.Mock
}

func (_m *IEmployeeResource) EXPECT() *IEmployeeResource_Expecter {
	return &IEmployeeResource_Expecter{mock: &_m.Mock}
}

// CountEmployees provides a mock function with given fields: query
func (_m *IEmployeeResource) CountEmployees(query *types.QuerySQL) (int64, error) {
	ret := _m.Called(query)

	if len(ret) == 0 {
		panic("no return value specified for CountEmployees")
	}

	var r0 int64
	var r1 error
	if rf, ok := ret.Get(0).(func(*types.QuerySQL) (int64, error)); ok {
		return rf(query)
	}
	if rf, ok := ret.Get(0).(func(*types.QuerySQL) int64); ok {
		r0 = rf(query)
	} else {
		r0 = ret.Get(0).(int64)
	}

	if rf, ok := ret.Get(1).(func(*types.QuerySQL) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IEmployeeResource_CountEmployees_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CountEmployees'
type IEmployeeResource_CountEmployees_Call struct {
	*mock.Call
}

// CountEmployees is a helper method to define mock.On call
//   - query *types.QuerySQL
func (_e *IEmployeeResource_Expecter) CountEmployees(query interface{}) *IEmployeeResource_CountEmployees_Call {
	return &IEmployeeResource_CountEmployees_Call{Call: _e.mock.On("CountEmployees", query)}
}

func (_c *IEmployeeResource_CountEmployees_Call) Run(run func(query *types.QuerySQL)) *IEmployeeResource_CountEmployees_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*types.QuerySQL))
	})
	return _c
}

func (_c *IEmployeeResource_CountEmployees_Call) Return(_a0 int64, _a1 error) *IEmployeeResource_CountEmployees_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IEmployeeResource_CountEmployees_Call) RunAndReturn(run func(*types.QuerySQL) (int64, error)) *IEmployeeResource_CountEmployees_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteEmployee provides a mock function with given fields: payload
func (_m *IEmployeeResource) DeleteEmployee(payload *sql.Employee) error {
	ret := _m.Called(payload)

	if len(ret) == 0 {
		panic("no return value specified for DeleteEmployee")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*sql.Employee) error); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IEmployeeResource_DeleteEmployee_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteEmployee'
type IEmployeeResource_DeleteEmployee_Call struct {
	*mock.Call
}

// DeleteEmployee is a helper method to define mock.On call
//   - payload *sql.Employee
func (_e *IEmployeeResource_Expecter) DeleteEmployee(payload interface{}) *IEmployeeResource_DeleteEmployee_Call {
	return &IEmployeeResource_DeleteEmployee_Call{Call: _e.mock.On("DeleteEmployee", payload)}
}

func (_c *IEmployeeResource_DeleteEmployee_Call) Run(run func(payload *sql.Employee)) *IEmployeeResource_DeleteEmployee_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*sql.Employee))
	})
	return _c
}

func (_c *IEmployeeResource_DeleteEmployee_Call) Return(_a0 error) *IEmployeeResource_DeleteEmployee_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IEmployeeResource_DeleteEmployee_Call) RunAndReturn(run func(*sql.Employee) error) *IEmployeeResource_DeleteEmployee_Call {
	_c.Call.Return(run)
	return _c
}

// FindEmployees provides a mock function with given fields: query
func (_m *IEmployeeResource) FindEmployees(query *types.QuerySQL) ([]*sql.Employee, error) {
	ret := _m.Called(query)

	if len(ret) == 0 {
		panic("no return value specified for FindEmployees")
	}

	var r0 []*sql.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(*types.QuerySQL) ([]*sql.Employee, error)); ok {
		return rf(query)
	}
	if rf, ok := ret.Get(0).(func(*types.QuerySQL) []*sql.Employee); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*sql.Employee)
		}
	}

	if rf, ok := ret.Get(1).(func(*types.QuerySQL) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IEmployeeResource_FindEmployees_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FindEmployees'
type IEmployeeResource_FindEmployees_Call struct {
	*mock.Call
}

// FindEmployees is a helper method to define mock.On call
//   - query *types.QuerySQL
func (_e *IEmployeeResource_Expecter) FindEmployees(query interface{}) *IEmployeeResource_FindEmployees_Call {
	return &IEmployeeResource_FindEmployees_Call{Call: _e.mock.On("FindEmployees", query)}
}

func (_c *IEmployeeResource_FindEmployees_Call) Run(run func(query *types.QuerySQL)) *IEmployeeResource_FindEmployees_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*types.QuerySQL))
	})
	return _c
}

func (_c *IEmployeeResource_FindEmployees_Call) Return(_a0 []*sql.Employee, _a1 error) *IEmployeeResource_FindEmployees_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IEmployeeResource_FindEmployees_Call) RunAndReturn(run func(*types.QuerySQL) ([]*sql.Employee, error)) *IEmployeeResource_FindEmployees_Call {
	_c.Call.Return(run)
	return _c
}

// FirstEmployee provides a mock function with given fields: query
func (_m *IEmployeeResource) FirstEmployee(query *types.QuerySQL) (*sql.Employee, error) {
	ret := _m.Called(query)

	if len(ret) == 0 {
		panic("no return value specified for FirstEmployee")
	}

	var r0 *sql.Employee
	var r1 error
	if rf, ok := ret.Get(0).(func(*types.QuerySQL) (*sql.Employee, error)); ok {
		return rf(query)
	}
	if rf, ok := ret.Get(0).(func(*types.QuerySQL) *sql.Employee); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Employee)
		}
	}

	if rf, ok := ret.Get(1).(func(*types.QuerySQL) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IEmployeeResource_FirstEmployee_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FirstEmployee'
type IEmployeeResource_FirstEmployee_Call struct {
	*mock.Call
}

// FirstEmployee is a helper method to define mock.On call
//   - query *types.QuerySQL
func (_e *IEmployeeResource_Expecter) FirstEmployee(query interface{}) *IEmployeeResource_FirstEmployee_Call {
	return &IEmployeeResource_FirstEmployee_Call{Call: _e.mock.On("FirstEmployee", query)}
}

func (_c *IEmployeeResource_FirstEmployee_Call) Run(run func(query *types.QuerySQL)) *IEmployeeResource_FirstEmployee_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*types.QuerySQL))
	})
	return _c
}

func (_c *IEmployeeResource_FirstEmployee_Call) Return(_a0 *sql.Employee, _a1 error) *IEmployeeResource_FirstEmployee_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IEmployeeResource_FirstEmployee_Call) RunAndReturn(run func(*types.QuerySQL) (*sql.Employee, error)) *IEmployeeResource_FirstEmployee_Call {
	_c.Call.Return(run)
	return _c
}

// SaveEmployee provides a mock function with given fields: payload
func (_m *IEmployeeResource) SaveEmployee(payload *sql.Employee) error {
	ret := _m.Called(payload)

	if len(ret) == 0 {
		panic("no return value specified for SaveEmployee")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*sql.Employee) error); ok {
		r0 = rf(payload)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// IEmployeeResource_SaveEmployee_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveEmployee'
type IEmployeeResource_SaveEmployee_Call struct {
	*mock.Call
}

// SaveEmployee is a helper method to define mock.On call
//   - payload *sql.Employee
func (_e *IEmployeeResource_Expecter) SaveEmployee(payload interface{}) *IEmployeeResource_SaveEmployee_Call {
	return &IEmployeeResource_SaveEmployee_Call{Call: _e.mock.On("SaveEmployee", payload)}
}

func (_c *IEmployeeResource_SaveEmployee_Call) Run(run func(payload *sql.Employee)) *IEmployeeResource_SaveEmployee_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*sql.Employee))
	})
	return _c
}

func (_c *IEmployeeResource_SaveEmployee_Call) Return(_a0 error) *IEmployeeResource_SaveEmployee_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *IEmployeeResource_SaveEmployee_Call) RunAndReturn(run func(*sql.Employee) error) *IEmployeeResource_SaveEmployee_Call {
	_c.Call.Return(run)
	return _c
}

// NewIEmployeeResource creates a new instance of IEmployeeResource. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIEmployeeResource(t interface {
	mock.TestingT
	Cleanup(func())
}) *IEmployeeResource {
	mock := &IEmployeeResource{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}