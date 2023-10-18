package migration

import (
	"github.com/rodericusifo/fiber-template/internal/app/model/database/sql"
)

var (
	AutoMigrateModelList = []any{
		&sql.User{},
		&sql.Employee{},
	}
)
