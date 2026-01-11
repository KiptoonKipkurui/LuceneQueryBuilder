package query

// AggrFunc defines how two expressions are combined.
type AggrFunc func(left, right Expr) Expr

// Aggregate reduces a slice of expressions into one expression
// using the provided aggregation function.
//
// Panics if exprs is empty (programmer error).
func Aggregate(exprs []Expr, aggrFn AggrFunc) Expr {
	if len(exprs) == 0 {
		panic("cannot aggregate empty expression list")
	}

	agg := exprs[0]
	for i := 1; i < len(exprs); i++ {
		agg = aggrFn(agg, exprs[i])
	}
	return agg
}
