package query

import "strings"

func Build(expr Expr) string {
	var sb strings.Builder
	expr.Build(&sb)
	return sb.String()
}
