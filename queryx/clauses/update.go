package clauses

type Update struct {
	Table     string
	Columns   []string
	Values    []any
	Where     string
	WhereArgs []any
}

func NewUpdate(table string, columns []string, values []any, where string, whereArgs []any) *Update {
	return &Update{
		Table:     table,
		Columns:   columns,
		Values:    values,
		Where:     where,
		WhereArgs: whereArgs,
	}
}
