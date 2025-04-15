package builder

import (
	"fmt"
	"strings"
)

func buildJoins(qb *QueryBuilder, b *strings.Builder, args []any) []any {
	if len(qb.joinClause) == 0 {
		return args
	}

	joinClauses := make([]string, 0, len(qb.joinClause))
	for _, j := range qb.joinClause {
		joinClauses = append(joinClauses, fmt.Sprintf("%s %s ON %s", j.Type, j.Table, j.Condition))
		args = append(args, j.Args...)
	}

	b.WriteString(" ")
	b.WriteString(strings.Join(joinClauses, " "))
	return args
}
