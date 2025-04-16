package queryx

import (
	"fmt"
	"strings"
)

func buildSelect(qb *QueryBuilder, b *strings.Builder) {
	b.WriteString("SELECT ")
	b.WriteString(strings.Join(qb.selectClause.Columns, ", "))
}

func buildFrom(qb *QueryBuilder, b *strings.Builder) {
	b.WriteString(" FROM ")
	b.WriteString(qb.fromClause.Table)
}

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

func buildGroupBy(qb *QueryBuilder, b *strings.Builder, args []any) []any {
	if qb.groupByClause != nil {
		b.WriteString(" GROUP BY ")
		b.WriteString(strings.Join(qb.groupByClause.Columns, ", "))
	}
	return args
}

func buildLimt(qb *QueryBuilder, b *strings.Builder, args []any) []any {
	if qb.limitClause != nil {
		b.WriteString(" LIMIT ?")
		args = append(args, qb.limitClause.Limit)
	}
	return args
}

func buildOffset(qb *QueryBuilder, b *strings.Builder, args []any) []any {
	if qb.offsetClause != nil {
		b.WriteString(" OFFSET ?")
		args = append(args, qb.offsetClause.Offset)
	}
	return args
}

func buildOrderBy(qb *QueryBuilder, b *strings.Builder, args []any) []any {
	if qb.orderByClause != nil {
		b.WriteString(" ORDER BY ")
		b.WriteString(strings.Join(qb.orderByClause.Columns, ", "))
	}
	return args
}
