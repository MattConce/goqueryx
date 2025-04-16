package clauses

const (
	LeftJoin  = "LEFT JOIN"
	InnerJoin = "INNER JOIN"
)

type Join struct {
	Type      string
	Table     string
	Condition string
	Args      []any
}

func NewInnerJoin(table string, condition string, args []any) *Join {
	return &Join{Type: InnerJoin, Table: table, Condition: condition, Args: args}
}

func NewLeftJoin(table string, condition string, args []any) *Join {
	return &Join{Type: LeftJoin, Table: table, Condition: condition, Args: args}
}
