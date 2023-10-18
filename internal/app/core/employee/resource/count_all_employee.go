package resource

import (
	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func (r *EmployeeResource) CountAllEmployee(query *pkg_types.QuerySQL) (int, error) {
	return r.EmployeeDatabaseSQLRepository.CountAllEmployee(query)
}
