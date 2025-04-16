package queryx

import (
	"reflect"
	"strings"
	"testing"

	"github.com/MattConce/goqueryx/queryx/clauses"
)

func TestBuildSelect(t *testing.T) {
	cols := []string{"id", "name"}
	qb := &QueryBuilder{
		selectClause: &clauses.Select{Columns: cols},
	}

	var sql strings.Builder
	buildSelect(qb, &sql)

	expected := "SELECT id, name"
	if sql.String() != expected {
		t.Errorf("\nexpected: %q\ngot: %q", expected, sql.String())
	}
}

func TestBuildFrom(t *testing.T) {
	qb := &QueryBuilder{
		fromClause: &clauses.From{Table: "users"},
	}

	var sql strings.Builder
	buildFrom(qb, &sql)

	expected := " FROM users"
	if sql.String() != expected {
		t.Errorf("\nexpected: %q\ngot: %q", expected, sql.String())
	}
}

func TestBuildWhere(t *testing.T) {
	whereClauses := make([]*clauses.Where, 0)
	whereClauses = append(whereClauses, &clauses.Where{Condition: "age > ?", Args: []any{18}})

	qb := &QueryBuilder{
		whereClause: whereClauses,
	}

	var sql strings.Builder
	args := buildWhere(qb, &sql, nil)

	expectedExpr := " WHERE age > ?"
	expectedArgs := []any{18}

	if sql.String() != expectedExpr {
		t.Errorf("\nexpected: %q\ngot: %q", expectedExpr, sql.String())
	}
	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("expected args: %v, got: %v", expectedArgs, args)
	}
}

func TestBuildJoin_InnerJoin(t *testing.T) {
	joins := make([]*clauses.Join, 0)
	joins = append(joins, &clauses.Join{Type: clauses.InnerJoin, Table: "address", Condition: "user.id = address.user_id", Args: nil})

	qb := &QueryBuilder{
		joinClause: joins,
	}

	var sql strings.Builder
	buildJoins(qb, &sql, nil)

	expected := " INNER JOIN address ON user.id = address.user_id"
	if sql.String() != expected {
		t.Errorf("\nexpected: %q\ngot: %q", expected, sql.String())
	}
}

func TestBuildJoin_LeftJoin(t *testing.T) {
	joins := make([]*clauses.Join, 0)
	joins = append(joins, &clauses.Join{Type: clauses.LeftJoin, Table: "address", Condition: "user.id = address.user_id", Args: nil})

	qb := &QueryBuilder{
		joinClause: joins,
	}

	var sql strings.Builder
	buildJoins(qb, &sql, nil)

	expected := " LEFT JOIN address ON user.id = address.user_id"
	if sql.String() != expected {
		t.Errorf("\nexpected: %q\ngot: %q", expected, sql.String())
	}
}

func TestBuildGroupBy(t *testing.T) {
	cols := []string{"id", "name"}
	qb := &QueryBuilder{
		groupByClause: &clauses.GroupBy{Columns: cols},
	}

	var sql strings.Builder
	buildGroupBy(qb, &sql, nil)

	expected := " GROUP BY id, name"
	if sql.String() != expected {
		t.Errorf("\nexpected: %q\ngot: %q", expected, sql.String())
	}
}

func TestBuildLimit(t *testing.T) {
	limit := 10

	qb := &QueryBuilder{limitClause: &clauses.Limit{Limit: limit}}

	var sql strings.Builder
	args := buildLimt(qb, &sql, nil)

	expectedExpr := " LIMIT ?"
	expectedArgs := []any{limit}

	if sql.String() != expectedExpr {
		t.Errorf("\nexpected: %q\ngot: %q", expectedExpr, sql.String())
	}
	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("expected args: %v, got: %v", expectedArgs, args)
	}
}

func TestBuildOffset(t *testing.T) {
	offset := 10

	qb := &QueryBuilder{offsetClause: &clauses.Offset{Offset: offset}}

	var sql strings.Builder
	args := buildOffset(qb, &sql, nil)

	expectedExpr := " OFFSET ?"
	expectedArgs := []any{offset}

	if sql.String() != expectedExpr {
		t.Errorf("\nexpected: %q\ngot: %q", expectedExpr, sql.String())
	}
	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("expected args: %v, got: %v", expectedArgs, args)
	}
}

func TestBuildOrderBy(t *testing.T) {
	cols := []string{"name DESC", "id ASC"}
	qb := &QueryBuilder{
		orderByClause: &clauses.OrderBy{Columns: cols},
	}

	var sql strings.Builder
	buildOrderBy(qb, &sql, nil)

	expected := " ORDER BY name DESC, id ASC"
	if sql.String() != expected {
		t.Errorf("\nexpected: %q\ngot: %q", expected, sql.String())
	}
}
