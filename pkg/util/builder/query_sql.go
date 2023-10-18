package builder

import (
	"gorm.io/gorm"

	pkg_constant "github.com/rodericusifo/fiber-template/pkg/constant"
	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
	pkg_util_generator "github.com/rodericusifo/fiber-template/pkg/util/generator"
	pkg_util_merger "github.com/rodericusifo/fiber-template/pkg/util/merger"
)

func BuildQuerySQL(tableName string, db *gorm.DB, query *pkg_types.QuerySQL, dialect pkg_constant.DialectDatabaseSQL) *gorm.DB {
	q := db

	if len(query.Joins) > 0 {
		for _, join := range query.Joins {
			if len(join.Selects) > 0 || len(join.Searches) > 0 {
				qj := db.Begin()
				if len(join.Selects) > 0 {
					querySlice := pkg_util_generator.GenerateSelectJoinQuerySQLSlice(join.Relation, pkg_util_merger.MergeSlices(true, join.Selects, pkg_constant.DEFAULT_SELECTS_JOIN_COLUMNS), dialect)
					if query.Distinct {
						qj = qj.Distinct(querySlice)
					} else {
						qj = qj.Select(querySlice)
					}
				}
				if len(join.Searches) > 0 {
					queryString, bindValues := pkg_util_generator.GenerateWhereQuerySQLStringAndBindValues(join.Relation, join.Searches, dialect)
					qj = qj.Where(queryString, bindValues...)
				}
				qj.Commit()
				q = q.Joins(join.Relation, qj)
			} else {
				q = q.Joins(join.Relation)
			}
		}
	}
	if len(query.InnerJoins) > 0 {
		for _, innerJoin := range query.InnerJoins {
			if len(innerJoin.Selects) > 0 || len(innerJoin.Searches) > 0 {
				qj := db.Begin()
				if len(innerJoin.Selects) > 0 {
					querySlice := pkg_util_generator.GenerateSelectJoinQuerySQLSlice(innerJoin.Relation, pkg_util_merger.MergeSlices(true, innerJoin.Selects, pkg_constant.DEFAULT_SELECTS_JOIN_COLUMNS), dialect)
					if query.Distinct {
						qj = qj.Distinct(querySlice)
					} else {
						qj = qj.Select(querySlice)
					}
				}
				if len(innerJoin.Searches) > 0 {
					queryString, bindValues := pkg_util_generator.GenerateWhereQuerySQLStringAndBindValues(innerJoin.Relation, innerJoin.Searches, dialect)
					qj = qj.Where(queryString, bindValues...)
				}
				qj.Commit()
				q = q.InnerJoins(innerJoin.Relation, qj)
			} else {
				q = q.InnerJoins(innerJoin.Relation)
			}
		}
	}
	if len(query.Selects) > 0 {
		querySlice := pkg_util_generator.GenerateSelectQuerySQLSlice(tableName, pkg_util_merger.MergeSlices(true, query.Selects, pkg_constant.DEFAULT_SELECTS_COLUMNS), dialect)
		if query.Distinct {
			q = q.Distinct(querySlice)
		} else {
			q = q.Select(querySlice)
		}
	}
	if len(query.Searches) > 0 {
		queryString, bindValues := pkg_util_generator.GenerateWhereQuerySQLStringAndBindValues(tableName, query.Searches, dialect)
		q = q.Where(queryString, bindValues...)
	}
	if len(query.Orders) > 0 {
		queryString := pkg_util_generator.GenerateOrderQuerySQLString(tableName, query.Orders, dialect)
		q = q.Order(queryString)
	}
	if len(query.Groups) > 0 {
		queryString := pkg_util_generator.GenerateGroupQuerySQLString(tableName, query.Groups, dialect)
		q = q.Group(queryString)
	}
	if query.WithDeleted {
		q = q.Unscoped()
	}
	if query.Limit != 0 {
		q = q.Limit(query.Limit)
	}
	if query.Offset != 0 {
		q = q.Offset(query.Offset)
	}

	return q
}
