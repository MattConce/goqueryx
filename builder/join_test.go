package builder

import (
	"strings"
	"testing"

	"github.com/MattConce/goqueryx/clauses"
)

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
