// Package builder provides a easy SQL query builder for constructing
// sql statements in a safe and idiomatic way.
//
// Example:
//
//	qb := builder.New().
//	    Select("id", "name").
//	    From("users").
//	    Where("age > ?", []any{18}).
//	    Build()
package builder

import (
	"errors"
	"strings"

	"github.com/MattConce/goqueryx/clauses"
)

type QueryBuilder struct {
	selectClause  *clauses.Select
	fromClause    *clauses.From
	whereClause   []*clauses.Where
	joinClause    []*clauses.Join
	orderByClause *clauses.OrderBy
	groupByClause *clauses.GroupBy
	limitClause   *clauses.Limit
	offsetClause  *clauses.Offset
}

func New() *QueryBuilder {
	return &QueryBuilder{}
}

func (qb *QueryBuilder) Select(columns ...string) *QueryBuilder {
	qb.selectClause = clauses.NewSelect(columns...)
	return qb
}

func (qb *QueryBuilder) From(table string) *QueryBuilder {
	qb.fromClause = clauses.NewFrom(table)
	return qb
}

func (qb *QueryBuilder) Where(condition string, args []any) *QueryBuilder {
	qb.whereClause = append(qb.whereClause, clauses.NewWhere(condition, args))
	return qb
}

func (qb *QueryBuilder) GroupBy(columns ...string) *QueryBuilder {
	qb.groupByClause = clauses.NewGroupBy(columns...)
	return qb
}

func (qb *QueryBuilder) OrderBy(columns ...string) *QueryBuilder {
	qb.orderByClause = clauses.NewOrderBy(columns...)
	return qb
}

func (qb *QueryBuilder) Limit(limit any) *QueryBuilder {
	qb.limitClause = clauses.NewLimit(limit)
	return qb
}

func (qb *QueryBuilder) Offset(offset any) *QueryBuilder {
	qb.offsetClause = clauses.NewOffset(offset)
	return qb
}

func (qb *QueryBuilder) Join(table, condition string, args []any) *QueryBuilder {
	qb.joinClause = append(qb.joinClause, clauses.NewInnerJoin(table, condition, args))
	return qb
}

func (qb *QueryBuilder) LeftJoin(table, condition string, args []any) *QueryBuilder {
	qb.joinClause = append(qb.joinClause, clauses.NewLeftJoin(table, condition, args))
	return qb
}

func (qb *QueryBuilder) Build() (string, []any, error) {
	if qb.selectClause == nil || len(qb.selectClause.Columns) == 0 {
		return "", nil, errors.New("select clause is required and must have columns")
	}
	if qb.fromClause == nil || qb.fromClause.Table == "" {
		return "", nil, errors.New("from clause is required and table cannot be empty")
	}

	var sqlBuilder strings.Builder
	var args []any

	buildSelect(qb, &sqlBuilder)
	buildFrom(qb, &sqlBuilder)

	args = buildJoins(qb, &sqlBuilder, args)
	args = buildWhere(qb, &sqlBuilder, args)
	args = buildGroupBy(qb, &sqlBuilder, args)
	args = buildOrderBy(qb, &sqlBuilder, args)
	args = buildLimt(qb, &sqlBuilder, args)
	args = buildOffset(qb, &sqlBuilder, args)

	return sqlBuilder.String(), args, nil
}
