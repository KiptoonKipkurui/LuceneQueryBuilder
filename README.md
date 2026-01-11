# LuceneQueryBuilder

A lightweight **Go** library for **programmatically building Lucene-style query strings** in a safe, composable, and testable way. The project abstracts common query patterns (terms, ranges, boolean logic, and constraints) into strongly typed builders, reducing string-concatenation errors and improving readability when constructing complex search queries.

This library is useful when working with search engines or platforms that expose a Lucene-compatible query syntax (for example, log analytics, document search, or security event filtering systems).

---

## ✨ Features

* **Type-safe query construction** – avoid brittle string concatenation
* **Composable expressions** – build complex queries from small reusable parts
* **Boolean logic support** – AND / OR / NOT composition
* **Constraint handling** – ranges, equality, and comparison operators
* **Test-driven design** – comprehensive unit tests for correctness
* **Zero dependencies** – pure Go standard library

---

## 📦 Installation

```bash
go get github.com/kiptoonkipkurui/LuceneQueryBuilder
```


---

## 🚀 Quick Start

### Basic Term Query

```go
q := Term("status", "active")
fmt.Println(q.String())
```

Output:

```
status:active
```

---

### Boolean Composition

```go
q := And(
    Term("status", "active"),
    Term("role", "admin"),
)

fmt.Println(q.String())
```

Output:

```
(status:active AND role:admin)
```

---

### OR / NOT Expressions

```go
q := Or(
    Term("env", "prod"),
    Not(Term("env", "dev")),
)

fmt.Println(q.String())
```

Output:

```
(env:prod OR NOT env:dev)
```

---

### Range Constraints

```go
q := Range("timestamp", ">=", "2024-01-01")
fmt.Println(q.String())
```

Output:

```
timestamp:>=2024-01-01
```

---

## 🧱 Core Concepts

### Expressions

An **Expression** represents any valid Lucene query fragment. Every expression implements:

```go
String() string
```

This allows expressions to be nested and composed freely.

---

### Constraints

Constraints define **field-level conditions**, such as:

* Equality
* Inequality
* Greater-than / less-than
* Ranges

They are implemented in `constraint.go` and validated through unit tests.

---

### Matchers

Matchers encapsulate logic for **matching fields against values** and ensure that generated queries follow valid Lucene syntax.

---

## 🧪 Testing

The project follows a test-first approach. All major components are covered by unit tests:

```bash
go test ./...
```

Test files:

* `constraint_test.go`
* `expressions_test.go`
* `matcher_test.go`

These tests verify:

* Correct query string generation
* Operator precedence and grouping
* Edge cases in boolean composition

---

## 📁 Project Structure

```
LuceneQueryBuilder/
├── constraint.go          # Field constraints and comparisons
├── expressions.go         # Boolean and logical expressions
├── matcher.go             # Field/value match helpers
├── *_test.go              # Unit tests
├── go.mod                 # Go module definition
├── README.md              # Project documentation
```

---

## 🔒 Security

If you discover a security issue, please refer to `SECURITY.md` for reporting guidelines.

---

## 🤝 Contributing

Contributions are welcome. Please see `CONTRIBUTING.md` for guidelines and best practices.

---

## 📜 License

This project is licensed under the terms of the **MIT License**. See the `LICENSE` file for details.

---

## 🎯 Use Cases

* Building search filters dynamically in backend services
* Generating Lucene queries from APIs or user-defined rules
* Safer query generation for logging, SIEM, and analytics platforms
* Improving readability and maintainability of complex search logic

---

## 🛣️ Future Improvements

* Query escaping and sanitisation helpers
* Support for fuzzy and wildcard queries
* JSON → Lucene query translation
* Fluent builder API

---

**Author:** Daniel Kiptoon

If you have questions or ideas, feel free to open an issue or start a discussion.
