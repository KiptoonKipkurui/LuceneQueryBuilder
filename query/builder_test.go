package query

import "testing"

func TestComplexQuery(t *testing.T) {
	expr := Or(
		And(
			Term("author", "alice"),
			Term("year", "2024"),
		),
		Not(Term("status", "draft")),
	)

	got := Build(expr)
	expected := "((author:alice AND year:2024) OR NOT status:draft)"

	if got != expected {
		t.Fatalf("expected %q, got %q", expected, got)
	}
}
