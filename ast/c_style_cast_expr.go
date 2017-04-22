package ast

import (
	"github.com/elliotchance/c2go/program"
)

type CStyleCastExpr struct {
	Address  string
	Position string
	Type     string
	Kind     string
	Children []Node
}

func parseCStyleCastExpr(line string) *CStyleCastExpr {
	groups := groupsFromRegex(
		"<(?P<position>.*)> '(?P<type>.*?)' <(?P<kind>.*)>",
		line,
	)

	return &CStyleCastExpr{
		Address:  groups["address"],
		Position: groups["position"],
		Type:     groups["type"],
		Kind:     groups["kind"],
		Children: []Node{},
	}
}

func (n *CStyleCastExpr) render(program *program.Program) (string, string) {
	children := n.Children
	return renderExpression(program, children[0])
}

func (n *CStyleCastExpr) AddChild(node Node) {
	n.Children = append(n.Children, node)
}
