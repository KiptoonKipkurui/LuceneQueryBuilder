package query

import "testing"

func TestMatchAny(t *testing.T) {
	expr := MatchAny("tag", []string{"a", "b", "c"})
	got := Build(expr)
	expected := "((tag:a OR tag:b) OR tag:c)"

	if got != expected {
		t.Fatalf("expected %q, got %q", expected, got)
	}
}

func TestMatchAll(t *testing.T) {
	expr := MatchAll("tag", []string{"a", "b", "c"})
	got := Build(expr)
	expected := "((tag:a AND tag:b) AND tag:c)"

	if got != expected {
		t.Fatalf("expected %q, got %q", expected, got)
	}
}
