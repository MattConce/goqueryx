package clauses

type OrderBy struct {
	Columns []string
}

func NewOrderBy(columns ...string) *OrderBy {
	return &OrderBy{Columns: columns}
}
