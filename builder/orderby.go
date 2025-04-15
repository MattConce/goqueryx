package builder

import (
	"strings"
)

func buildOrderBy(qb *QueryBuilder, b *strings.Builder, args []any) []any {
	if qb.orderByClause != nil {
		b.WriteString(" ORDER BY ")
		b.WriteString(strings.Join(qb.orderByClause.Columns, ", "))
	}
	return args
}
