package query

// Match creates a single field:value term.
func Match(field, value string) Expr {
	return Term(field, value)
}
