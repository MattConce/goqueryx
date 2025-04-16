package clauses

type Where struct {
	Condition string
	Args      []any
}

func NewWhere(condition string, args []any) *Where {
	return &Where{Condition: condition, Args: args}
}
