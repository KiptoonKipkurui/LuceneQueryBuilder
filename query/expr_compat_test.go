package query

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
				expr := And(
					Term("title", "politics"),
					Term("title", "fashion"),
				)
				return Build(expr)
			},
			expected: "(title:politics AND title:fashion)",
		},
		{
			name: "AND with NOT",
			build: func() string {
				expr := And(
					Term("title", "politics"),
					Term("title", "fashion"),
				)
				expr = And(expr, Not(Term("title", "sports")))
				return Build(expr)
			},
			expected: "((title:politics AND title:fashion) AND NOT title:sports)",
		},
		{
			name: "AND chained with OR",
			build: func() string {
				expr := And(
					Term("title", "politics"),
					Term("title", "fashion"),
				)
				expr = Or(expr, Term("title", "sports"))
				return Build(expr)
			},
			expected: "((title:politics AND title:fashion) OR title:sports)",
		},
		{
			name: "AND chained with AND",
			build: func() string {
				expr := And(
					Term("title", "politics"),
					Term("title", "fashion"),
				)
				expr = And(expr, Term("title", "sports"))
				return Build(expr)
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
	expr := Or(
		Term("author", "alice"),
		And(
			Term("title", "politics"),
			Term("year", "2024"),
		),
	)

	got := Build(expr)
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
	base := Term("a", "1")

	expr := And(base, Term("b", "2"))
	expr = Or(expr, Term("c", "3"))
	expr = And(expr, Not(Term("d", "4")))

	got := Build(expr)
	expected := "(((a:1 AND b:2) OR c:3) AND NOT d:4)"

	if got != expected {
		t.Fatalf("expected %q, got %q", expected, got)
	}
}
