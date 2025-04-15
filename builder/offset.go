package builder

import "strings"

func buildOffset(qb *QueryBuilder, b *strings.Builder, args []any) []any {
	if qb.offsetClause != nil {
		b.WriteString(" OFFSET ?")
		args = append(args, qb.offsetClause.Offset)
	}
	return args
}
