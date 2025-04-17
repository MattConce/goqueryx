package clauses

type Delete struct {
	Table     string
	Where     string
	WhereArgs []any
}

func NewDelete(table string, where string, whereArgs []any) *Delete {
	return &Delete{
		Table:     table,
		Where:     where,
		WhereArgs: whereArgs,
	}
}
