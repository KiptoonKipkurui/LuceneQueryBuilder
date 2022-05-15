package main

import (
	"strings"
)

type Expression struct {
	Operator   Operator
	Left       *Expression
	Right      *Expression
	Constraint *Constraint
	IsEmpty    bool
}

func EmptyExpression() *Expression {
	return &Expression{IsEmpty: true}
}

func NewExpression(constraint Constraint) *Expression {
	return &Expression{
		Constraint: &constraint,
		IsEmpty:    false,
	}
}

func NewExpressionConstraint(field, value string) *Expression {
	constraint := NewConstraint(field, value)

	return NewExpression(*constraint)
}

func NewExpressionCtor(op Operator, left, right Expression) *Expression {
	return &Expression{
		Left:     &left,
		Right:    &right,
		Operator: op,
	}
}

func (e Expression) IsConstraint() bool {
	return e.Constraint != nil
}

func (e Expression) ToBuilder() strings.Builder {
	if e.IsConstraint() {
		return e.Constraint.ToBuilder()
	} else {
		return e.InParens()
	}
}

func (e Expression) InParens() strings.Builder {

	var sb strings.Builder

	sb.WriteString("(")
	lstr := e.Left.ToBuilder()
	l := lstr.String()
	sb.WriteString(l)
	sb.WriteString(" ")
	sb.WriteString(strings.ToUpper(string(e.Operator)))

	sb.WriteString(" ")

	rstr := e.Right.ToBuilder()
	r := rstr.String()
	sb.WriteString(r)

	sb.WriteString(")")
	return sb
}

func (e *Expression) NotExpr(field, value string) *Expression {
	return e.Apply(Not, field, value)
}

func (e *Expression) Not(other Expression) *Expression {
	return e.apply(Not, other)
}

func (e *Expression) OrExpr(field, value string) *Expression {
	return e.Apply(Or, field, value)
}

func (e *Expression) Or(other Expression) *Expression {
	return e.apply(Or, other)
}

func (e *Expression) AndExpr(field, value string) *Expression {
	return e.Apply(And, field, value)
}

func (e *Expression) And(other Expression) *Expression {
	return e.apply(And, other)
}

func (e *Expression) Apply(op Operator, field, value string) *Expression {

	constraint := NewConstraint(field, value)
	return e.apply(op, *NewExpression(*constraint))
}

func (e *Expression) apply(op Operator, other Expression) *Expression {
	//todo: error handling
	if other.IsEmpty {
		return e
	}

	if !e.IsEmpty && !other.IsEmpty {
		return &Expression{
			Operator: op,
			Left:     e,
			Right:    &other,
			IsEmpty:  false,
		}
	}

	switch op {
	case And:
	case Or:
		return &other
	case Not:
		leftConstraint := NewConstraint("*", "*")
		return NewExpressionCtor(op, *NewExpression(*leftConstraint), other)
	default:
		panic("Unknown operator")
	}
	return nil
}

//Operator enum
type Operator string

const (
	And Operator = "And"
	Or  Operator = "Or"
	Not Operator = "Not"
)
