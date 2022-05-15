package main

import (
	"fmt"
	"testing"
)

func TestExpress(t *testing.T) {
	constraintLeft := NewConstraint("title", "politics")

	constraintRight := NewConstraint("title", "fashion")

	expression := NewExpressionCtor(And, *NewExpression(*constraintLeft), *NewExpression(*constraintRight))

	sb := expression.ToBuilder()
	query := sb.String()

	fmt.Print(query)
}

func TestNotExpr(t *testing.T) {
	constraintLeft := NewConstraint("title", "politics")

	constraintRight := NewConstraint("title", "fashion")

	expression := NewExpressionCtor(And, *NewExpression(*constraintLeft), *NewExpression(*constraintRight))

	expression = expression.NotExpr("title", "sports")

	sb := expression.ToBuilder()
	query := sb.String()

	fmt.Print(query)
}

func TestAndExpr(t *testing.T) {
	constraintLeft := NewConstraint("title", "politics")

	constraintRight := NewConstraint("title", "fashion")

	expression := NewExpressionCtor(And, *NewExpression(*constraintLeft), *NewExpression(*constraintRight))

	expression = expression.AndExpr("title", "sports")

	sb := expression.ToBuilder()
	query := sb.String()

	fmt.Print(query)
}

func TestOrExpr(t *testing.T) {
	constraintLeft := NewConstraint("title", "politics")

	constraintRight := NewConstraint("title", "fashion")

	expression := NewExpressionCtor(And, *NewExpression(*constraintLeft), *NewExpression(*constraintRight))

	expression = expression.OrExpr("title", "sports")

	sb := expression.ToBuilder()
	query := sb.String()

	fmt.Print(query)
}
