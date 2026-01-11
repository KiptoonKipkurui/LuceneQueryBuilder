package query

import "strings"

func Term(field, value string) Expr {
	return ConstraintExpr{Constraint: NewConstraint(field, value)}
}

type ConstraintExpr struct {
	Constraint Constraint
}

func (c ConstraintExpr) Build(sb *strings.Builder) {
	c.Constraint.Build(sb)
}

func And(left, right Expr) Expr {
	return BinaryExpr{Op: OpAnd, Left: left, Right: right}
}

func Or(left, right Expr) Expr {
	return BinaryExpr{Op: OpOr, Left: left, Right: right}
}

func Not(expr Expr) Expr {
	return NotExpr{Expr: expr}
}
