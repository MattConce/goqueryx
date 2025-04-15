package builder

import (
	"strings"
)

func buildGroupBy(qb *QueryBuilder, b *strings.Builder, args []any) []any {
	if qb.groupByClause != nil {
		b.WriteString(" GROUP BY ")
		b.WriteString(strings.Join(qb.groupByClause.Columns, ", "))
	}
	return args
}
