package query

// MatchAny builds an OR expression across all values.
func MatchAny(field string, values []string) Expr {
	return MatchAnyExpr(toExprArray(field, values))
}

// MatchAll builds an AND expression across all values.
func MatchAll(field string, values []string) Expr {
	return MatchAllExpr(toExprArray(field, values))
}

// MatchAnyExpr OR-reduces a list of expressions.
func MatchAnyExpr(exprs []Expr) Expr {
	return Aggregate(exprs, Or)
}

// MatchAllExpr AND-reduces a list of expressions.
func MatchAllExpr(exprs []Expr) Expr {
	return Aggregate(exprs, And)
}
