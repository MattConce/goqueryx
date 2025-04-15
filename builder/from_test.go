package builder

import (
	"strings"
	"testing"

	"github.com/MattConce/goqueryx/clauses"
)

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
