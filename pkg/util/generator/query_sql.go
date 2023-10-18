package generator

import (
	"fmt"
	"strings"

	pkg_constant "github.com/rodericusifo/fiber-template/pkg/constant"
	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

func GenerateSelectQuerySQLSlice(tableAlias string, selects []pkg_types.SelectQuerySQLOperation, dialect pkg_constant.DialectDatabaseSQL) []string {
	querySlice := make([]string, 0)
	tableAlias = strings.Replace(tableAlias, ".", "__", -1)

	switch dialect {
	case pkg_constant.MYSQL:
		for _, s := range selects {
			fieldSelectStr := fmt.Sprintf("`%s`.`%s`", tableAlias, s.Field)
			if s.Function != "" {
				function := s.Function
				fieldStr := fmt.Sprintf("`%s`.`%s`", tableAlias, s.Field)
				fieldSelectStr = strings.Replace(function, "$", fieldStr, -1)
			}

			if s.Alias != "" {
				querySlice = append(querySlice, fmt.Sprintf("%s AS `%s`", fieldSelectStr, s.Alias))
			} else {
				querySlice = append(querySlice, fieldSelectStr)
			}
		}
	case pkg_constant.POSTGRES:
		for _, s := range selects {
			fieldSelectStr := fmt.Sprintf(`"%s"."%s"`, tableAlias, s.Field)
			if s.Function != "" {
				function := s.Function
				fieldStr := fmt.Sprintf(`"%s"."%s"`, tableAlias, s.Field)
				fieldSelectStr = strings.Replace(function, "$", fieldStr, -1)
			}

			if s.Alias != "" {
				querySlice = append(querySlice, fmt.Sprintf(`%s AS "%s"`, fieldSelectStr, s.Alias))
			} else {
				querySlice = append(querySlice, fieldSelectStr)
			}
		}
	}

	return querySlice
}

func GenerateSelectJoinQuerySQLSlice(tableAlias string, selects []pkg_types.SelectJoinQuerySQLOperation, dialect pkg_constant.DialectDatabaseSQL) []string {
	querySlice := make([]string, 0)
	tableAlias = strings.Replace(tableAlias, ".", "__", -1)

	switch dialect {
	case pkg_constant.MYSQL:
		for _, s := range selects {
			fieldSelectStr := fmt.Sprintf("`%s`.`%s`", tableAlias, s.Field)

			querySlice = append(querySlice, fieldSelectStr)
		}
	case pkg_constant.POSTGRES:
		for _, s := range selects {
			fieldSelectStr := fmt.Sprintf(`"%s"."%s"`, tableAlias, s.Field)

			querySlice = append(querySlice, fieldSelectStr)
		}
	}

	return querySlice
}

func GenerateWhereQuerySQLStringAndBindValues(tableAlias string, searches [][]pkg_types.SearchQuerySQLOperation, dialect pkg_constant.DialectDatabaseSQL) (string, []any) {
	queryString := ""
	bindValues := make([]any, 0)
	tableAlias = strings.Replace(tableAlias, ".", "__", -1)

	switch dialect {
	case pkg_constant.MYSQL:
		for indexOuter, searchOuter := range searches {
			if indexOuter > 0 {
				queryString += " OR "
			}
			for indexInner, searchInner := range searchOuter {
				if indexInner > 0 {
					queryString += " AND "
				}
				if searchInner.Value != nil {
					queryString += fmt.Sprintf("`%s`.`%s` %s ?", tableAlias, searchInner.Field, searchInner.Operator)
					bindValues = append(bindValues, searchInner.Value)
				} else {
					queryString += fmt.Sprintf("`%s`.`%s` %s", tableAlias, searchInner.Field, searchInner.Operator)
				}
			}
		}
	case pkg_constant.POSTGRES:
		for indexOuter, searchOuter := range searches {
			if indexOuter > 0 {
				queryString += " OR "
			}
			for indexInner, searchInner := range searchOuter {
				if indexInner > 0 {
					queryString += " AND "
				}
				if searchInner.Value != nil {
					queryString += fmt.Sprintf(`"%s"."%s" %s ?`, tableAlias, searchInner.Field, searchInner.Operator)
					bindValues = append(bindValues, searchInner.Value)
				} else {
					queryString += fmt.Sprintf(`"%s"."%s" %s`, tableAlias, searchInner.Field, searchInner.Operator)
				}
			}
		}
	}

	return queryString, bindValues
}

func GenerateOrderQuerySQLString(tableAlias string, orders []pkg_types.OrderQuerySQLOperation, dialect pkg_constant.DialectDatabaseSQL) string {
	querySlice := make([]string, 0)
	tableAlias = strings.Replace(tableAlias, ".", "__", -1)

	switch dialect {
	case pkg_constant.MYSQL:
		for _, order := range orders {
			if order.Descending {
				querySlice = append(querySlice, fmt.Sprintf("`%s`.`%s` DESC", tableAlias, order.Field))
			} else {
				querySlice = append(querySlice, fmt.Sprintf("`%s`.`%s`", tableAlias, order.Field))
			}
		}
	case pkg_constant.POSTGRES:
		for _, order := range orders {
			if order.Descending {
				querySlice = append(querySlice, fmt.Sprintf(`"%s"."%s" DESC`, tableAlias, order.Field))
			} else {
				querySlice = append(querySlice, fmt.Sprintf(`"%s"."%s"`, tableAlias, order.Field))
			}
		}
	}

	return strings.Join(querySlice, ",")
}

func GenerateGroupQuerySQLString(tableAlias string, groups []pkg_types.GroupQuerySQLOperation, dialect pkg_constant.DialectDatabaseSQL) string {
	querySlice := make([]string, 0)
	tableAlias = strings.Replace(tableAlias, ".", "__", -1)

	switch dialect {
	case pkg_constant.MYSQL:
		for _, group := range groups {
			fieldSelectStr := ""
			if group.Function != "" {
				function := group.Function
				fieldStr := fmt.Sprintf("`%s`.`%s`", tableAlias, group.Field)
				fieldSelectStr = strings.Replace(function, "$", fieldStr, -1)
			} else {
				fieldSelectStr = fmt.Sprintf("`%s`.`%s`", tableAlias, group.Field)
			}
			querySlice = append(querySlice, fieldSelectStr)
		}
	case pkg_constant.POSTGRES:
		for _, group := range groups {
			fieldSelectStr := ""
			if group.Function != "" {
				function := group.Function
				fieldStr := fmt.Sprintf(`"%s"."%s"`, tableAlias, group.Field)
				fieldSelectStr = strings.Replace(function, "$", fieldStr, -1)
			} else {
				fieldSelectStr = fmt.Sprintf(`"%s"."%s"`, tableAlias, group.Field)
			}
			querySlice = append(querySlice, fieldSelectStr)
		}
	}

	return strings.Join(querySlice, ",")
}
