package project

import "fmt"

type Node struct {
	kind string
	path string
	name string
	icon string
}

func CreateNode(name string, path string, isCollection bool) (node Node) {
	kind := "API"
	icon := "󰅩"

	if isCollection {
		kind = "COLLECTION"
		icon = ""
	}

	return Node{
		kind: kind,
		path: path,
		name: name,
		icon: icon,
	}
}

func (node *Node) Kind() string {
	return node.kind
}

func (node *Node) Path() string {
	return node.path
}

func (node *Node) Name(complete bool) string {
	if complete {
		return fmt.Sprintf(" %s %s ", node.icon, node.name)
	}

	return node.name
}

func (node *Node) Icon() string {
	return node.icon
}

func (node *Node) IsCollection() bool {
	return node.kind == "COLLECTION"
}
