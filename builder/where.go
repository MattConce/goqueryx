package builder

import (
	"strings"
)

func buildWhere(qb *QueryBuilder, b *strings.Builder, args []any) []any {
	if len(qb.whereClause) == 0 {
		return args
	}

	conditions := make([]string, 0, len(qb.whereClause))
	for _, w := range qb.whereClause {
		conditions = append(conditions, w.Condition)
		args = append(args, w.Args...)
	}

	b.WriteString(" WHERE ")
	b.WriteString(strings.Join(conditions, " AND "))
	return args
}
