package main

import (
	"fmt"
	"testing"
)

func TestConstaint(t *testing.T) {
	constraint := NewConstraint("title", "test")

	var sb = constraint.ToBuilder()
	var final = sb.String()

	if final != "title:test" {
		t.Fatal("Failed constraint test")
	}
	fmt.Print(final)
}
