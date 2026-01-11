package query

func toExprArray(field string, values []string) []Expr {
	exprs := make([]Expr, 0, len(values))
	for _, v := range values {
		exprs = append(exprs, Match(field, v))
	}
	return exprs
}
