package query

import "strings"

type Constraint struct {
	Field string
	Value string
}

func NewConstraint(field, value string) Constraint {
	if field == "" {
		panic("field cannot be null")
	}
	return Constraint{
		Field: field,
		Value: value,
	}
}

func (c Constraint) Build(sb *strings.Builder) {
	sb.WriteString(c.Field)
	sb.WriteString(":")
	sb.WriteString(c.Value)
}
