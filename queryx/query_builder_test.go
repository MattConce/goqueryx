package queryx

import (
	"reflect"
	"testing"
)

func TestQueryBuilder_Build_Success(t *testing.T) {
	qb := NewQuery().
		Select("id", "name", "email").
		From("users").
		Where("id = ?", []any{1}).
		Where("status = ?", []any{true})

	sql, args, err := qb.Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedExpr := "SELECT id, name, email FROM users WHERE id = ? AND status = ?"
	expectedArgs := []any{1, true}

	if sql != expectedExpr {
		t.Errorf("expected SQL:\n%s\ngot:\n%s", expectedExpr, sql)
	}
	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("expected args: %v, got: %v", expectedArgs, args)
	}
}

func TestQueryBuilder_Build_MissingSelect(t *testing.T) {
	qb := NewQuery().From("users")
	_, _, err := qb.Build()
	if err == nil {
		t.Error("expected an error when select clause is missing, but got none")
	}
}

func TestQueryBuilder_Build_MissingFrom(t *testing.T) {
	qb := NewQuery().Select("id", "name")
	_, _, err := qb.Build()
	if err == nil {
		t.Error("expected an error when from clause is missing, but got none")
	}
}

func TestQueryBuilder_Build_NoWhere(t *testing.T) {
	qb := NewQuery().Select("id", "name").From("users")
	sql, args, err := qb.Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedExpr := "SELECT id, name FROM users"

	if sql != expectedExpr {
		t.Errorf("expected SQL:\n%s\ngot:\n%s", expectedExpr, sql)
	}
	if len(args) != 0 {
		t.Errorf("expected no args, got: %v", args)
	}
}

func TestQueryBuilder_Build_Join(t *testing.T) {
	qb := NewQuery().
		Select("id", "name").
		From("users").
		Join("address", "users.id = address.user_id", nil)

	sql, args, err := qb.Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedExpr := "SELECT id, name FROM users INNER JOIN address ON users.id = address.user_id"

	if sql != expectedExpr {
		t.Errorf("expected SQL:\n%s\ngot:\n%s", expectedExpr, sql)
	}
	if len(args) != 0 {
		t.Errorf("expected no args, got: %v", args)
	}
}

func TestQueryBuilder_Build_LeftJoin(t *testing.T) {
	qb := NewQuery().
		Select("id", "name").
		From("users").
		LeftJoin("address", "users.id = address.user_id", nil)

	sql, args, err := qb.Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedExpr := "SELECT id, name FROM users LEFT JOIN address ON users.id = address.user_id"

	if sql != expectedExpr {
		t.Errorf("expected SQL:\n%s\ngot:\n%s", expectedExpr, sql)
	}
	if len(args) != 0 {
		t.Errorf("expected no args, got: %v", args)
	}
}

func TestQueryBuilder_Build_OrderBy(t *testing.T) {
	qb := NewQuery().Select("id", "name").From("users").OrderBy("name DESC", "id ASC")
	sql, args, err := qb.Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedExpr := "SELECT id, name FROM users ORDER BY name DESC, id ASC"

	if sql != expectedExpr {
		t.Errorf("expected SQL:\n%s\ngot:\n%s", expectedExpr, sql)
	}
	if len(args) != 0 {
		t.Errorf("expected no args, got: %v", args)
	}
}

func TestQueryBuilder_Build_Limit_Offset(t *testing.T) {
	qb := NewQuery().
		Select("id", "name", "email").
		From("users").
		Limit(10).
		Offset(10)

	sql, args, err := qb.Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedExpr := "SELECT id, name, email FROM users LIMIT ? OFFSET ?"
	expectedArgs := []any{10, 10}

	if sql != expectedExpr {
		t.Errorf("expected SQL:\n%s\ngot:\n%s", expectedExpr, sql)
	}
	if !reflect.DeepEqual(args, expectedArgs) {
		t.Errorf("expected args: %v, got: %v", expectedArgs, args)
	}
}

func TestQueryBuilder_Build_GroupBy(t *testing.T) {
	qb := NewQuery().Select("count(id) AS count", "name").From("users").GroupBy("name")
	sql, args, err := qb.Build()
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expectedExpr := "SELECT count(id) AS count, name FROM users GROUP BY name"

	if sql != expectedExpr {
		t.Errorf("expected SQL:\n%s\ngot:\n%s", expectedExpr, sql)
	}
	if len(args) != 0 {
		t.Errorf("expected no args, got: %v", args)
	}
}
