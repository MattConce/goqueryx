package clauses

type GroupBy struct {
	Columns []string
}

func NewGroupBy(columns ...string) *GroupBy {
	return &GroupBy{Columns: columns}
}
