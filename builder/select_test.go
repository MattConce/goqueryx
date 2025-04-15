package builder

import (
	"strings"
	"testing"

	"github.com/MattConce/goqueryx/clauses"
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
