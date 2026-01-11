package query

import "strings"

type NotExpr struct {
	Expr Expr
}

func (n NotExpr) Build(sb *strings.Builder) {
	sb.WriteString("NOT ")
	n.Expr.Build(sb)
}
