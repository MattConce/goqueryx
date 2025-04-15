package clauses

type Select struct {
	Columns []string
}

func NewSelect(columns ...string) *Select {
	return &Select{Columns: columns}
}
