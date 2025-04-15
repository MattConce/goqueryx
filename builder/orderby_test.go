package builder

import (
	"strings"
	"testing"

	"github.com/MattConce/goqueryx/clauses"
)

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
