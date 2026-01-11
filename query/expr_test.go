package query

import "testing"

func TestNestedExpression(t *testing.T) {
	expr := And(
		Term("title", "politics"),
		Not(Term("title", "sports")),
	)

	got := Build(expr)
	expected := "(title:politics AND NOT title:sports)"

	if got != expected {
		t.Fatalf("expected %q, got %q", expected, got)
	}
}
func FuzzBuildDoesNotPanic(f *testing.F) {
	f.Add("a", "1")
	f.Add("title", "politics")

	f.Fuzz(func(t *testing.T, field, value string) {
		defer func() {
			if recover() != nil {
				t.Fatal("builder panicked")
			}
		}()

		expr := Term(field, value)
		_ = Build(expr)
	})
}
func TestExpressionBehaviourTable(t *testing.T) {
	tests := []struct {
		name     string
		expr     Expr
		expected string
	}{
		{"simple term", Term("a", "1"), "a:1"},
		{"not", Not(Term("a", "1")), "NOT a:1"},
		{"and", And(Term("a", "1"), Term("b", "2")), "(a:1 AND b:2)"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Build(tt.expr); got != tt.expected {
				t.Fatalf("expected %q, got %q", tt.expected, got)
			}
		})
	}
}
