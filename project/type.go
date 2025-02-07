package project

import (
	"fmt"
	"strings"
)

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

func (node *Node) Path(parent bool) string {
	if node.IsCollection() {
		return node.path
	}

	if !parent {
		return node.path
	}

	path, valid := strings.CutSuffix(node.path, node.Name(false, true))

	if !valid {
		panic("Invalid Path Detected")
	}

	return path
}

func (node *Node) Name(icon bool, extension bool) string {
	var name string = node.name

	if icon {
		name = fmt.Sprintf("%s %s", node.icon, name)
	}

	if !extension {
		name = strings.Split(name, ".")[0]
	}

	return name
}

func (node *Node) Icon() string {
	return node.icon
}

func (node *Node) IsCollection() bool {
	return node.kind == "COLLECTION"
}
