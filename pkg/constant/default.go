package constant

import (
	pkg_types "github.com/rodericusifo/fiber-template/pkg/types"
)

type DefaultSelects []pkg_types.SelectQuerySQLOperation
type DefaultSelectsJoin []pkg_types.SelectJoinQuerySQLOperation

var (
	DEFAULT_SELECTS_COLUMNS      = DefaultSelects([]pkg_types.SelectQuerySQLOperation{})
	DEFAULT_SELECTS_JOIN_COLUMNS = DefaultSelectsJoin([]pkg_types.SelectJoinQuerySQLOperation{})
)
