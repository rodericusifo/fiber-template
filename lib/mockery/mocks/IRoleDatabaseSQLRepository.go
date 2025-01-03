// Code generated by mockery v2.50.0. DO NOT EDIT.

package mocks

import (
	mock "github.com/stretchr/testify/mock"

	sql "github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	types "github.com/rodericusifo/fiber-template/pkg/types"
)

// IRoleDatabaseSQLRepository is an autogenerated mock type for the IRoleDatabaseSQLRepository type
type IRoleDatabaseSQLRepository struct {
	mock.Mock
}

type IRoleDatabaseSQLRepository_Expecter struct {
	mock *mock.Mock
}

func (_m *IRoleDatabaseSQLRepository) EXPECT() *IRoleDatabaseSQLRepository_Expecter {
	return &IRoleDatabaseSQLRepository_Expecter{mock: &_m.Mock}
}

// FirstRole provides a mock function with given fields: query
func (_m *IRoleDatabaseSQLRepository) FirstRole(query *types.QuerySQL) (*sql.Role, error) {
	ret := _m.Called(query)

	if len(ret) == 0 {
		panic("no return value specified for FirstRole")
	}

	var r0 *sql.Role
	var r1 error
	if rf, ok := ret.Get(0).(func(*types.QuerySQL) (*sql.Role, error)); ok {
		return rf(query)
	}
	if rf, ok := ret.Get(0).(func(*types.QuerySQL) *sql.Role); ok {
		r0 = rf(query)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*sql.Role)
		}
	}

	if rf, ok := ret.Get(1).(func(*types.QuerySQL) error); ok {
		r1 = rf(query)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IRoleDatabaseSQLRepository_FirstRole_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FirstRole'
type IRoleDatabaseSQLRepository_FirstRole_Call struct {
	*mock.Call
}

// FirstRole is a helper method to define mock.On call
//   - query *types.QuerySQL
func (_e *IRoleDatabaseSQLRepository_Expecter) FirstRole(query interface{}) *IRoleDatabaseSQLRepository_FirstRole_Call {
	return &IRoleDatabaseSQLRepository_FirstRole_Call{Call: _e.mock.On("FirstRole", query)}
}

func (_c *IRoleDatabaseSQLRepository_FirstRole_Call) Run(run func(query *types.QuerySQL)) *IRoleDatabaseSQLRepository_FirstRole_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*types.QuerySQL))
	})
	return _c
}

func (_c *IRoleDatabaseSQLRepository_FirstRole_Call) Return(_a0 *sql.Role, _a1 error) *IRoleDatabaseSQLRepository_FirstRole_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *IRoleDatabaseSQLRepository_FirstRole_Call) RunAndReturn(run func(*types.QuerySQL) (*sql.Role, error)) *IRoleDatabaseSQLRepository_FirstRole_Call {
	_c.Call.Return(run)
	return _c
}

// NewIRoleDatabaseSQLRepository creates a new instance of IRoleDatabaseSQLRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewIRoleDatabaseSQLRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *IRoleDatabaseSQLRepository {
	mock := &IRoleDatabaseSQLRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
