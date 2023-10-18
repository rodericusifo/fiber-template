package resource

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"

	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func (r *EmployeeResource) GetListEmployeeAndCount(query *pkg_types.QuerySQL) ([]*sql.Employee, int, error) {
	return r.EmployeeDatabaseSQLRepository.GetListEmployeeAndCount(query)
}
