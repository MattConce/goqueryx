package builder

import (
	"reflect"
	"strings"
	"testing"

	"github.com/MattConce/goqueryx/clauses"
)

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
