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

func NewNode() *Node {
	return &Node{
		parent: nil,
		kind:   "",
		path:   "",
		name:   "",
		icon:   "",
	}
}

func (node *Node) Initialize(parent *tview.TreeNode, name string, path string, isCollection bool) *Node {
	node.parent = parent
	node.kind = "API"
	node.path = path
	node.name = name
	node.icon = ""

	if isCollection {
		node.kind = "COLLECTION"
		node.icon = ""
	}

	return node
}

func (node *Node) Parent() (parent *tview.TreeNode) {
	return node.parent
}

func (node *Node) ParentNode() (parent *Node) {
	reference, _ := node.parent.GetReference().(*Node)
	return reference
}

func (node *Node) SetParent(parent *tview.TreeNode) *Node {
	node.parent = parent
	return node
}

func (node *Node) Kind() (kind string) {
	return node.kind
}

func (node *Node) SetKind(kind string) *Node {
	if kind != "API" && kind != "COLLECTION" {
		panic("Invalid Node Type")
	}

	node.kind = kind
	return node
}

func (node *Node) Path(parent bool) (path string) {
	if !parent || node.Collection() {
		return node.path
	}

	path, ok := strings.CutSuffix(node.path, node.Name(false, true))

	if !ok {
		panic("Invalid Path Detected")
	}

	return path
}

func (node *Node) SetPath(path string) *Node {
	node.path = path
	return node
}

func (node *Node) Name(icon bool, extension bool) string {
	name := node.name

	if icon {
		name = fmt.Sprintf("%s %s", node.icon, name)
	}

	if !extension {
		name = strings.Split(name, ".")[0]
	}

	return name
}

func (node *Node) SetName(name string) *Node {
	node.name = name
	return node
}

func (node *Node) Icon() (icon string) {
	return node.icon
}

func (node *Node) SetIcon(icon string) *Node {
	node.icon = icon
	return node
}

func (node *Node) Collection() bool {
	return node.kind == "COLLECTION"
}

func (node *Node) Render() *tview.TreeNode {
	return tview.
		NewTreeNode(node.Name(true, false)).
		SetSelectable(true).
		SetExpanded(false).
		SetReference(node)
}
