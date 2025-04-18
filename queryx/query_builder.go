// Package queryx provides a easy SQL query builder for constructing
// sql statements in a safe and idiomatic way.
//
// Example:
//
//	qb := queryx.NewQuery().
//	    Select("id", "name").
//	    From("users").
//	    Where("age > ?", []any{18}).
//	    Build()
package queryx

import (
	"errors"
	"slices"
	"strings"

	"github.com/MattConce/goqueryx/queryx/clauses"
)

type QueryBuilder struct {
	isCount           bool
	selectClause      *clauses.Select
	insertClause      *clauses.Insert
	updateClause      *clauses.Update
	valuesClause      *clauses.Values
	multiValuesClause *clauses.MultiValues
	deleteClause      *clauses.Delete
	fromClause        *clauses.From
	whereClause       []*clauses.Where
	havingClause      []*clauses.Having
	joinClause        []*clauses.Join
	orderByClause     *clauses.OrderBy
	groupByClause     *clauses.GroupBy
	limitClause       *clauses.Limit
	offsetClause      *clauses.Offset
}

func NewQuery() *QueryBuilder {
	return &QueryBuilder{}
}

func (qb *QueryBuilder) Insert(table string, columns []string) *QueryBuilder {
	qb.insertClause = clauses.NewInsert(table, columns)
	return qb
}

func (qb *QueryBuilder) Update(table string, columns []string) *QueryBuilder {
	qb.updateClause = clauses.NewUpdate(table, columns)
	return qb
}

func (qb *QueryBuilder) Values(values ...any) *QueryBuilder {
	qb.valuesClause = clauses.NewValues(values)
	return qb
}

func (qb *QueryBuilder) MultiValues(values [][]any) *QueryBuilder {
	qb.multiValuesClause = clauses.NewMultiValues(values)
	return qb
}

func (qb *QueryBuilder) Delete(table string) *QueryBuilder {
	qb.deleteClause = clauses.NewDelete(table)
	return qb
}

func (qb *QueryBuilder) CountTotal() *QueryBuilder {
	newQb := qb.cloneForCount()
	newQb.isCount = true
	newQb.orderByClause = nil
	newQb.limitClause = nil
	newQb.offsetClause = nil
	return newQb
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

func (qb *QueryBuilder) Having(condition string, args []any) *QueryBuilder {
	qb.havingClause = append(qb.havingClause, clauses.NewHaving(condition, args))
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
	var sqlBuilder strings.Builder
	var args []any

	switch {
	case qb.isCount:
		if qb.groupByClause != nil && len(qb.groupByClause.Columns) > 0 {
			sqlBuilder.WriteString("SELECT COUNT(*) FROM (SELECT 1")
			buildFrom(qb, &sqlBuilder)
			args = buildJoins(qb, &sqlBuilder, args)
			args = buildWhere(qb, &sqlBuilder, args)
			sqlBuilder.WriteString(") AS subquery")
		} else {
			sqlBuilder.WriteString("SELECT COUNT(*)")
			buildFrom(qb, &sqlBuilder)
			args = buildJoins(qb, &sqlBuilder, args)
			args = buildWhere(qb, &sqlBuilder, args)
		}
	case qb.insertClause != nil:
		if qb.insertClause.Table == "" {
			return "", nil, errors.New("insert requires a table name")
		}
		if len(qb.insertClause.Columns) == 0 {
			return "", nil, errors.New("insert requires columns")
		}
		if qb.valuesClause == nil && qb.multiValuesClause == nil {
			return "", nil, errors.New("insert requires values or multi-values")
		}
		args = buildInsert(qb, &sqlBuilder, args)
		if qb.multiValuesClause != nil {
			args = buildMultiValues(qb, &sqlBuilder, args)
		} else {
			args = buildValues(qb, &sqlBuilder, args)
		}

	case qb.updateClause != nil:
		if qb.updateClause.Table == "" {
			return "", nil, errors.New("update requires a table name")
		}
		if len(qb.updateClause.Columns) == 0 {
			return "", nil, errors.New("update requires columns")
		}
		if qb.valuesClause == nil {
			return "", nil, errors.New("update requires values")
		}
		if len(qb.valuesClause.Args) != len(qb.updateClause.Columns) {
			return "", nil, errors.New("number of values must match columns in update")
		}
		args = buildUpdate(qb, &sqlBuilder, args)
		args = append(args, qb.valuesClause.Args...)
		args = buildWhere(qb, &sqlBuilder, args)

	case qb.deleteClause != nil:
		if qb.deleteClause.Table == "" {
			return "", nil, errors.New("delete requires a table name")
		}
		args = buildDelete(qb, &sqlBuilder, args)

		if len(qb.whereClause) == 0 {
			return "", nil, errors.New("delete requires where condition")
		}
		args = buildWhere(qb, &sqlBuilder, args)

	case qb.selectClause != nil:
		if len(qb.selectClause.Columns) == 0 {
			return "", nil, errors.New("select clause requires columns")
		}
		if qb.fromClause == nil || qb.fromClause.Table == "" {
			return "", nil, errors.New("from clause is required for select")
		}
		buildSelect(qb, &sqlBuilder)
		buildFrom(qb, &sqlBuilder)
		args = buildJoins(qb, &sqlBuilder, args)
		args = buildWhere(qb, &sqlBuilder, args)
		args = buildGroupBy(qb, &sqlBuilder, args)
		args = buildHaving(qb, &sqlBuilder, args)
		args = buildOrderBy(qb, &sqlBuilder, args)
		args = buildLimt(qb, &sqlBuilder, args)
		args = buildOffset(qb, &sqlBuilder, args)

	default:
		return "", nil, errors.New("no query type specified (select/insert/update)")
	}

	return sqlBuilder.String(), args, nil
}

func (qb *QueryBuilder) cloneForCount() *QueryBuilder {
	return &QueryBuilder{
		fromClause:    qb.fromClause,
		whereClause:   slices.Clone(qb.whereClause),
		joinClause:    slices.Clone(qb.joinClause),
		groupByClause: qb.groupByClause,
	}
}
