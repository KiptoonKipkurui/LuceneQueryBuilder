package main

import (
	"fmt"
	"strings"
)

type Constraint struct {
	Field string
	Value string
}

// return new constraint
func NewConstraint(field, value string) *Constraint {

	if field == "" {
		panic("field cannot be null")
	}

	return &Constraint{Field: field, Value: value}
}

func (c Constraint) ToBuilder() *strings.Builder {

	var sb strings.Builder

	var formated = fmt.Sprintf("%s:%s", c.Field, c.Value)
	sb.WriteString(formated)

	return &sb
}
