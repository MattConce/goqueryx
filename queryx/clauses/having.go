package clauses

type Having struct {
	Condition string
	Args      []any
}

func NewHaving(condition string, args []any) *Having {
	return &Having{Condition: condition, Args: args}
}
