package main

import (
	"fmt"
	"testing"
)

func TestMatchAny(t *testing.T) {
	field := "title"
	value := make([]string, 0)
	value = append(value, "fashion", "politics", "economics")

	var expr = MatchAny(field, value)

	sb := expr.ToBuilder()

	str := sb.String()

	fmt.Print(str)

}
