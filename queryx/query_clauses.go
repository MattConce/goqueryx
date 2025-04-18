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

func buildHaving(qb *QueryBuilder, b *strings.Builder, args []any) []any {
	if len(qb.havingClause) == 0 {
		return args
	}

	conditions := make([]string, 0, len(qb.havingClause))
	for _, w := range qb.havingClause {
		conditions = append(conditions, w.Condition)
		args = append(args, w.Args...)
	}

	b.WriteString(" HAVING ")
	b.WriteString(strings.Join(conditions, " AND "))
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

func buildInsert(qb *QueryBuilder, b *strings.Builder, args []any) []any {
	if qb.insertClause != nil {
		clause := qb.insertClause
		b.WriteString(fmt.Sprintf("INSERT INTO %s (%s)",
			clause.Table,
			strings.Join(clause.Columns, ", ")))
	}
	return args
}

func buildUpdate(qb *QueryBuilder, b *strings.Builder, args []any) []any {
	if qb.updateClause != nil {
		clause := qb.updateClause

		b.WriteString(fmt.Sprintf("UPDATE %s SET ", clause.Table))

		setClauses := make([]string, len(clause.Columns))
		for i, col := range clause.Columns {
			setClauses[i] = fmt.Sprintf("%s = ?", col)
		}
		b.WriteString(strings.Join(setClauses, ", "))
	}
	return args
}

func buildDelete(qb *QueryBuilder, b *strings.Builder, args []any) []any {
	if qb.deleteClause != nil {
		b.WriteString(fmt.Sprintf("DELETE FROM %s", qb.deleteClause.Table))
	}
	return args
}

func buildValues(qb *QueryBuilder, b *strings.Builder, args []any) []any {
	values := qb.valuesClause.Args
	placeholders := strings.Repeat("?, ", len(values))
	placeholders = strings.TrimSuffix(placeholders, ", ")
	b.WriteString(" VALUES (" + placeholders + ")")
	return append(args, values...)
}

func buildMultiValues(qb *QueryBuilder, b *strings.Builder, args []any) []any {
	placeholders := make([]string, 0, len(qb.multiValuesClause.Args))
	b.WriteString(" VALUES ")
	for _, row := range qb.multiValuesClause.Args {
		placeholders = append(placeholders,
			fmt.Sprintf("(%s)", strings.TrimSuffix(strings.Repeat("?, ", len(row)), ", ")))
		args = append(args, row...)
	}

	b.WriteString(strings.Join(placeholders, ", "))
	return args
}
