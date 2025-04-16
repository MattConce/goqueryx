package clauses

type Insert struct {
	Table   string
	Columns []string
	Values  [][]any
}

func NewInsert(table string, columns []string, values [][]any) *Insert {
	return &Insert{Table: table, Columns: columns, Values: values}
}
