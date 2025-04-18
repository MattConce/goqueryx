package clauses

type Update struct {
	Table   string
	Columns []string
}

func NewUpdate(table string, columns []string) *Update {
	return &Update{
		Table:   table,
		Columns: columns,
	}
}
