package main

import "testing"

func TestExpression_BooleanComposition(t *testing.T) {
	tests := []struct {
		name     string
		build    func() string
		expected string
	}{
		{
			name: "AND expression",
			build: func() string {
				left := NewConstraint("title", "politics")
				right := NewConstraint("title", "fashion")

				expr := NewExpressionCtor(
					And,
					*NewExpression(*left),
					*NewExpression(*right),
				)

				return expr.ToBuilder().String()
			},
			expected: "(title:politics AND title:fashion)",
		},
		{
			name: "AND with NOT",
			build: func() string {
				left := NewConstraint("title", "politics")
				right := NewConstraint("title", "fashion")

				expr := NewExpressionCtor(
					And,
					*NewExpression(*left),
					*NewExpression(*right),
				).NotExpr("title", "sports")

				return expr.ToBuilder().String()
			},
			expected: "((title:politics AND title:fashion) NOT title:sports)",
		},
		{
			name: "AND chained with OR",
			build: func() string {
				left := NewConstraint("title", "politics")
				right := NewConstraint("title", "fashion")

				expr := NewExpressionCtor(
					And,
					*NewExpression(*left),
					*NewExpression(*right),
				).OrExpr("title", "sports")

				return expr.ToBuilder().String()
			},
			expected: "((title:politics AND title:fashion) OR title:sports)",
		},
		{
			name: "AND chained with AND",
			build: func() string {
				left := NewConstraint("title", "politics")
				right := NewConstraint("title", "fashion")

				expr := NewExpressionCtor(
					And,
					*NewExpression(*left),
					*NewExpression(*right),
				).AndExpr("title", "sports")

				return expr.ToBuilder().String()
			},
			expected: "((title:politics AND title:fashion) AND title:sports)",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.build()
			if got != tt.expected {
				t.Fatalf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}

func TestExpression_NestedPrecedence(t *testing.T) {
	expr := NewExpressionCtor(
		Or,
		*NewExpression(*NewConstraint("author", "alice")),
		*NewExpressionCtor(
			And,
			*NewExpression(*NewConstraint("title", "politics")),
			*NewExpression(*NewConstraint("year", "2024")),
		),
	)

	got := expr.ToBuilder().String()
	expected := "(author:alice OR (title:politics AND year:2024))"

	if got != expected {
		t.Fatalf("expected %q, got %q", expected, got)
	}
}
func TestNewConstraint_EmptyFieldPanicsWithMessage(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Fatal("expected panic, got nil")
		}

		if r != "field cannot be null" {
			t.Fatalf("unexpected panic message: %v", r)
		}
	}()

	_ = NewConstraint("", "")
}

func TestExpression_ChainingIsStable(t *testing.T) {
	expr := NewExpression(*NewConstraint("a", "1"))

	expr = expr.
		AndExpr("b", "2").
		OrExpr("c", "3").
		NotExpr("d", "4")

	got := expr.ToBuilder().String()
	expected := "(((a:1 AND b:2) OR c:3) NOT d:4)"

	if got != expected {
		t.Fatalf("expected %q, got %q", expected, got)
	}
}
