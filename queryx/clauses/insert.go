package clauses

type Insert struct {
	Table   string
	Columns []string
}

func NewInsert(table string, columns []string) *Insert {
	return &Insert{Table: table, Columns: columns}
}
