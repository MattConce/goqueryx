package builder

import (
	"strings"
)

func buildFrom(qb *QueryBuilder, b *strings.Builder) {
	b.WriteString(" FROM ")
	b.WriteString(qb.fromClause.Table)
}
