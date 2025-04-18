package clauses

type Delete struct {
	Table string
}

func NewDelete(table string) *Delete {
	return &Delete{
		Table: table,
	}
}
