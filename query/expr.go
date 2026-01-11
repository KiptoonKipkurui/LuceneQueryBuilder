package query

import "strings"

// Expr represents any valid Lucene query expression.
type Expr interface {
	Build(*strings.Builder)
}
