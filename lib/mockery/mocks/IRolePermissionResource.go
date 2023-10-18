// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	sql "github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	types "github.com/rodericusifo/fiber-template/pkg/types"
)

// IRolePermissionResource is an autogenerated mock type for the IRolePermissionResource type
type IRolePermissionResource struct {
	mock.Mock
}

type IRolePermissionResource_Expecter struct {
	mock *mock.Mock
}

func (_m *IRolePermissionResource) EXPECT() *IRolePermissionResource_Expecter {
	return &IRolePermissionResource_Expecter{mock: &_m.Mock}
}

// FirstRolePermission provides a mock function with given fields: query
func (_m *IRolePermissionResource) FirstRolePermission(query *types.QuerySQL) (*sql.RolePermission, error) {
	ret := _m.Called(query)

	if len(ret) == 0 {
		panic("no return value specified for FirstRolePermission")
	}

	var r0 *sql.RolePermission
	var r1 error
	if rf, ok := ret.Get(0).(func(*types.QuerySQL) (*sql.RolePermission, error)); ok {
		return rf(query)
	}
	if rf, ok := ret.Get(0).(func(*types.QuerySQL) *sql.RolePermission); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.RolePermission)
		}
	}

	if rf, ok := ret.Get(1).(func(*types.QuerySQL) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IRolePermissionResource_FirstRolePermission_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FirstRolePermission'
type IRolePermissionResource_FirstRolePermission_Call struct {
	*mock.Call
}

// FirstRolePermission is a helper method to define mock.On call
//   - query *types.QuerySQL
func (_e *IRolePermissionResource_Expecter) FirstRolePermission(query interface{}) *IRolePermissionResource_FirstRolePermission_Call {
	return &IRolePermissionResource_FirstRolePermission_Call{Call: _e.mock.On("FirstRolePermission", query)}
}

func (_c *IRolePermissionResource_FirstRolePermission_Call) Run(run func(query *types.QuerySQL)) *IRolePermissionResource_FirstRolePermission_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*types.QuerySQL))
	})
	return _c
}

func (_c *IRolePermissionResource_FirstRolePermission_Call) Return(_a0 *sql.RolePermission, _a1 error) *IRolePermissionResource_FirstRolePermission_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IRolePermissionResource_FirstRolePermission_Call) RunAndReturn(run func(*types.QuerySQL) (*sql.RolePermission, error)) *IRolePermissionResource_FirstRolePermission_Call {
	_c.Call.Return(run)
	return _c
}

// NewIRolePermissionResource creates a new instance of IRolePermissionResource. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIRolePermissionResource(t interface {
	mock.TestingT
	Cleanup(func())
}) *IRolePermissionResource {
	mock := &IRolePermissionResource{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
