# goqueryx

A lightweight, ORM-free SQL query builder for Go. Designed for simplicity and compatibility with `database/sql` and libraries like `sqlx`.

**No magic** – just straightforward SQL construction.

## Features

- 🛠️ **Simple API**: Chainable methods for building SQL queries
- 🚫 **Zero ORM Dependencies**: Use raw SQL with your preferred driver (`pgx`, `sqlx`, `database/sql`, etc.)
- 🛡️ **SQL Injection Protection**: Parameterized arguments
- 📦 **Minimalist**: Pure GO, no external dependencies

## Installation

```bash
go get github.com/MattConce/goqueryx
```

## Why goqueryx?

Avoid ORM Complexity: Perfect for projects that prefer direct SQL control

Type-Safe Arguments: []any args work with all major SQL drivers

Transparent Queries: See exactly what SQL is being generated

## Usage with sqlx (Recommended)

```go
package main

import (
"fmt"
"github.com/jmoiron/sqlx"
"github.com/your-username/your-repo/builder"
)

func main() {
db := sqlx.MustConnect("postgres", "user=test dbname=test")

    qb := builder.New().
        Select("id", "name").
        From("users").
        Where("active = ?", []any{true})

    sql, args, _ := qb.Build()
    sql = db.Rebind(sql)

    var users []User

    err := db.Select(&users, sql, args...)
    if err != nil {
        panic(err)
    }

}
```

## Basic Usage

```go
qb := builder.New().
Select("id", "email").
From("users").
Where("created_at > ?", []any{"2024-01-01"}).
OrderBy("id DESC")

sql, args, _ := qb.Build()
// SQL: SELECT id, email FROM users WHERE created_at > ? ORDER BY id DESC
// Args: [2024-01-01]
```

## TODO

- Add support for INSERT/UPDATE/DELETE statements

- Support dialect-specific placeholders (e.g., $1 for PostgreSQL)

- Add query validation for unsupported operations

- Improve error messages for missing required clauses

License
MIT. See LICENSE.
