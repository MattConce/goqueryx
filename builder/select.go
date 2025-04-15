package builder

import (
	"strings"
)

func buildSelect(qb *QueryBuilder, b *strings.Builder) {
	b.WriteString("SELECT ")
	b.WriteString(strings.Join(qb.selectClause.Columns, ", "))
}
