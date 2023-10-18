package types

type SelectQuerySQLOperation struct {
	Field    string
	Alias    string
	Function string
}

type SelectJoinQuerySQLOperation struct {
	Field string
}

type SearchQuerySQLOperation struct {
	Field    string
	Operator string
	Value    any
}

type JoinQuerySQLOperation struct {
	Relation string
	Selects  []SelectJoinQuerySQLOperation
	Searches [][]SearchQuerySQLOperation
}

type InnerJoinQuerySQLOperation struct {
	Relation string
	Selects  []SelectJoinQuerySQLOperation
	Searches [][]SearchQuerySQLOperation
}

type OrderQuerySQLOperation struct {
	Field      string
	Descending bool
}

type GroupQuerySQLOperation struct {
	Field    string
	Function string
}

type QuerySQL struct {
	Selects     []SelectQuerySQLOperation
	Searches    [][]SearchQuerySQLOperation
	Joins       []JoinQuerySQLOperation
	InnerJoins  []InnerJoinQuerySQLOperation
	Orders      []OrderQuerySQLOperation
	Groups      []GroupQuerySQLOperation
	Distinct    bool
	WithDeleted bool
	Limit       int
	Offset      int
}
