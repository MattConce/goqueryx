package clauses

type From struct {
	Table string
}

func NewFrom(table string) *From {
	return &From{Table: table}
}
