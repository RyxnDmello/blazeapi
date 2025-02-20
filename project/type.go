package project

import (
	"fmt"
	"strings"

	"github.com/rivo/tview"
)

type Node struct {
	parent *tview.TreeNode
	kind   string
	path   string
	name   string
	icon   string
}

func CreateNode(parent *tview.TreeNode, name string, path string, isCollection bool) (node Node) {
	kind := "API"
	icon := "󰅩"

	if isCollection {
		kind = "COLLECTION"
		icon = ""
	}

	return Node{
		parent: parent,
		kind:   kind,
		path:   path,
		name:   name,
		icon:   icon,
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

func (node *Node) Parent() (parent *tview.TreeNode) {
	return node.parent
}

func (node *Node) ParentNode() (parent *Node) {
	reference, ok := node.parent.GetReference().(Node)

	if !ok {
		return nil
	}

	return &reference
}

func (node *Node) Collapse() {

}
