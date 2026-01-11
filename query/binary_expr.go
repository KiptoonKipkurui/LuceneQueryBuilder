package query

import "strings"

type BinaryExpr struct {
	Op    Operator
	Left  Expr
	Right Expr
}

func (b BinaryExpr) Build(sb *strings.Builder) {
	sb.WriteString("(")
	b.Left.Build(sb)
	sb.WriteString(" ")
	sb.WriteString(string(b.Op))
	sb.WriteString(" ")
	b.Right.Build(sb)
	sb.WriteString(")")
}
