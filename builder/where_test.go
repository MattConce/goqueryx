package builder

import (
	"reflect"
	"strings"
	"testing"

	"github.com/MattConce/goqueryx/clauses"
)

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
