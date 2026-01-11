package query

import "testing"

func TestConstraintPanicsOnEmptyField(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Fatal("expected panic for empty field")
		}
	}()

	_ = NewConstraint("", "value")
}
