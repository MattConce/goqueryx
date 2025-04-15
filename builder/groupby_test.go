package builder

import (
	"strings"
	"testing"

	"github.com/MattConce/goqueryx/clauses"
)

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
