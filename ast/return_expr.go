package ast

import "strings"

type Return_expr struct {
	Type string
	Expr string
}

func parse_return_expr(line string) (n interface{}) {
	groups := groupsFromRegex(
		`
	type:(?P<type>.*) +
	expr:(?P<expr>.*) +
	`,
		line,
	)
	return Return_expr{
		Type: strings.TrimSpace(groups["type"]),
		Expr: strings.TrimSpace(groups["expr"]),
	}
}