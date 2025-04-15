package builder

import (
	"reflect"
	"strings"
	"testing"

	"github.com/MattConce/goqueryx/clauses"
)

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
