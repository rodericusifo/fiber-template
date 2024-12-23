package resource

import (
	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func (r *EmployeeResource) CountEmployees(query *pkg_types.QuerySQL) (int64, error) {
	return r.EmployeeDatabaseSQLRepository.CountEmployees(query)
}
