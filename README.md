# Lucene Query Builder (Go)

A lightweight, immutable Go library for building Lucene-style query strings in a safe, composable, and testable way.

Instead of relying on string concatenation, the library models queries as a small expression tree (AST). This makes complex queries easier to reason about, extend, and test.

---

## Key Features

* Immutable expression tree (AST) design
* Fluent, readable API (`And`, `Or`, `Not`, `MatchAny`, `MatchAll`)
* Correct Lucene boolean semantics
* No string concatenation logic in user code
* Behaviour-driven tests that survive refactors
* Zero external dependencies

---

## Installation

```bash
go get github.com/kiptoonkipkurui/lucenequerybuilder
```

Import the query package:

```go
import "github.com/kiptoonkipkurui/lucenequerybuilder/query"
```

---

## Design Overview

Instead of building query strings directly, this library constructs an expression tree:

* Each node represents a logical operation or constraint
* The tree is rendered to a Lucene query string only at the end
* Expressions are immutable, so chaining never mutates previous queries

This approach avoids:

* Brittle string concatenation
* Invalid intermediate states
* Hidden side effects during chaining

---

## Quick Start

### Simple term

```go
expr := query.Term("status", "active")
q := query.Build(expr)
```

Output:

```
status:active
```

---

### Boolean composition

```go
expr := query.And(
	query.Term("title", "politics"),
	query.Term("title", "fashion"),
)

q := query.Build(expr)
```

Output:

```
(title:politics AND title:fashion)
```

---

### NOT expressions

```go
expr := query.Not(query.Term("status", "draft"))
q := query.Build(expr)
```

Output:

```
NOT status:draft
```

---

### Nested expressions and precedence

```go
expr := query.Or(
	query.Term("author", "alice"),
	query.And(
		query.Term("year", "2024"),
		query.Term("status", "published"),
	),
)

q := query.Build(expr)
```

Output:

```
(author:alice OR (year:2024 AND status:published))
```

---

## Match Helpers

### Match any (OR-reduce)

```go
expr := query.MatchAny("tag", []string{"ai", "ml", "search"})
q := query.Build(expr)
```

Output:

```
((tag:ai OR tag:ml) OR tag:search)
```

---

### Match all (AND-reduce)

```go
expr := query.MatchAll("category", []string{"tech", "backend"})
q := query.Build(expr)
```

Output:

```
(category:tech AND category:backend)
```

---

## Core Concepts

### Expr

```go
type Expr interface {
	Build(*strings.Builder)
}
```

Every query element implements `Expr`, allowing expressions to be freely composed and nested.

---

### Expression Types

* `ConstraintExpr` – field/value pairs (e.g. `title:politics`)
* `BinaryExpr` – AND / OR operations
* `NotExpr` – unary NOT expressions

---

### Build

```go
func Build(expr Expr) string
```

Renders the expression tree into a valid Lucene query string.

---

## Testing Strategy

The test suite focuses on behaviour, not implementation details.

Tests verify:

* Boolean correctness
* Operator precedence
* Chaining stability
* Legacy behaviour compatibility
* Defensive input validation

Example test:

```go
func TestChainedExpression(t *testing.T) {
	expr := query.And(
		query.Term("a", "1"),
		query.Not(query.Term("b", "2")),
	)

	got := query.Build(expr)
	expected := "(a:1 AND NOT b:2)"

	if got != expected {
		t.Fatalf("expected %q, got %q", expected, got)
	}
}
```

---

## Project Structure

```
lucenequerybuilder/
├── README.md
├── SECURITY.md
├── CONTRIBUTING.md
├── go.mod
└── query/
    ├── expr.go
    ├── constraint.go
    ├── binary_expr.go
    ├── not_expr.go
    ├── operators.go
    ├── fluent.go
    ├── aggregate.go
    ├── match.go
    ├── builder.go
    └── *_test.go
```

---

## Input Validation

* Empty fields are rejected at construction time
* Invalid states are prevented by design
* Programmer errors fail fast with clear panics

---

## Future Improvements

* Field/value escaping helpers
* Range queries (`[a TO b]`, `{}`)
* Wildcard and fuzzy matching
* JSON to Lucene query translation
* Optional parser for reverse conversion

---

## Contributing

Contributions are welcome.
Please see `CONTRIBUTING.md` for guidelines.

---

## License

MIT License. See `LICENSE` for details.

---

## Author

Daniel Kiptoon
