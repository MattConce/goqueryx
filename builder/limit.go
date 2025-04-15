package builder

import "strings"

func buildLimt(qb *QueryBuilder, b *strings.Builder, args []any) []any {
	if qb.limitClause != nil {
		b.WriteString(" LIMIT ?")
		args = append(args, qb.limitClause.Limit)
	}
	return args
}
