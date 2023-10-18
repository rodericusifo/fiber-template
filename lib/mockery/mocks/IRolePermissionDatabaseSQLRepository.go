// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	sql "github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	types "github.com/rodericusifo/fiber-template/pkg/types"
)

// IRolePermissionDatabaseSQLRepository is an autogenerated mock type for the IRolePermissionDatabaseSQLRepository type
type IRolePermissionDatabaseSQLRepository struct {
	mock.Mock
}

type IRolePermissionDatabaseSQLRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *IRolePermissionDatabaseSQLRepository) EXPECT() *IRolePermissionDatabaseSQLRepository_Expecter {
	return &IRolePermissionDatabaseSQLRepository_Expecter{mock: &_m.Mock}
}

// FirstRolePermission provides a mock function with given fields: query
func (_m *IRolePermissionDatabaseSQLRepository) FirstRolePermission(query *types.QuerySQL) (*sql.RolePermission, error) {
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

// IRolePermissionDatabaseSQLRepository_FirstRolePermission_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FirstRolePermission'
type IRolePermissionDatabaseSQLRepository_FirstRolePermission_Call struct {
	*mock.Call
}

// FirstRolePermission is a helper method to define mock.On call
//   - query *types.QuerySQL
func (_e *IRolePermissionDatabaseSQLRepository_Expecter) FirstRolePermission(query interface{}) *IRolePermissionDatabaseSQLRepository_FirstRolePermission_Call {
	return &IRolePermissionDatabaseSQLRepository_FirstRolePermission_Call{Call: _e.mock.On("FirstRolePermission", query)}
}

func (_c *IRolePermissionDatabaseSQLRepository_FirstRolePermission_Call) Run(run func(query *types.QuerySQL)) *IRolePermissionDatabaseSQLRepository_FirstRolePermission_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*types.QuerySQL))
	})
	return _c
}

func (_c *IRolePermissionDatabaseSQLRepository_FirstRolePermission_Call) Return(_a0 *sql.RolePermission, _a1 error) *IRolePermissionDatabaseSQLRepository_FirstRolePermission_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IRolePermissionDatabaseSQLRepository_FirstRolePermission_Call) RunAndReturn(run func(*types.QuerySQL) (*sql.RolePermission, error)) *IRolePermissionDatabaseSQLRepository_FirstRolePermission_Call {
	_c.Call.Return(run)
	return _c
}

// NewIRolePermissionDatabaseSQLRepository creates a new instance of IRolePermissionDatabaseSQLRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIRolePermissionDatabaseSQLRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IRolePermissionDatabaseSQLRepository {
	mock := &IRolePermissionDatabaseSQLRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}