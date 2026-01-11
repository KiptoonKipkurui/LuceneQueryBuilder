package query

// Operator represents a logical operator in a Lucene query.
type Operator string

const (
	OpAnd Operator = "AND"
	OpOr  Operator = "OR"
)
