package main

func Match(field, value string) *Expression {
	return NewExpressionConstraint(field, value)
}

func ToArray(field string, values []string) []*Expression {
	arr := make([]*Expression, 0)

	for _, v := range values {
		arr = append(arr, Match(field, v))
	}
	return arr
}

type AggrFunc func(expr1, expr2 *Expression) *Expression

func MatchAny(field string, value []string) *Expression {
	return MatchAnyExpr(ToArray(field, value))
}

func MatchAnyExpr(exprs []*Expression) *Expression {
	return Aggregate(exprs, func(expr1, expr2 *Expression) *Expression {
		return expr1.Or(*expr2)
	})
}

func MatchAll(field string, value []string) *Expression {
	return MatchAllExpr(ToArray(field, value))
}

func MatchAllExpr(exprs []*Expression) *Expression {
	return Aggregate(exprs, func(expr1, expr2 *Expression) *Expression {
		return expr1.And(*expr2)
	})
}

func Aggregate(exprs []*Expression, aggrFn func(expr1, expr2 *Expression) *Expression) *Expression {

	var agg = exprs[0]

	for i := 1; i < len(exprs); i++ {
		agg = aggrFn(agg, exprs[i])

	}
	return agg
}
