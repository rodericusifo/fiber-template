package resource

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func (r *EmployeeResource) GetEmployee(query *pkg_types.QuerySQL) (*sql.Employee, error) {
	return r.EmployeeDatabaseSQLRepository.GetEmployee(query)
}
