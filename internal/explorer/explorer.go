package explorer

import (
	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"

	"github.com/ryxndmello/flame/widgets/modal"
	"github.com/ryxndmello/flame/widgets/tree"
)

type Explorer struct {
	name              string
	path              string
	padding           int
	expanded          bool
	hierarchy         bool
	tree              *tview.TreeView
	addFileModal      *tview.Flex
	addFolderModal    *tview.Flex
	deleteFileModal   *tview.Flex
	deleteFolderModal *tview.Flex
}

func NewExplorer() *Explorer {
	return &Explorer{
		name:              ".",
		path:              ".",
		padding:           40,
		expanded:          true,
		hierarchy:         true,
		tree:              nil,
		addFileModal:      nil,
		addFolderModal:    nil,
		deleteFileModal:   nil,
		deleteFolderModal: nil,
	}
}

func NewDefaultExplorer(app *tview.Application) *Explorer {
	explorer := NewExplorer().
		SetName("Test").
		SetPath("./test").
		UseDefaultTree().
		UseDefaultAddFileModal(app).
		UseDefaultAddFolderModal(app).
		UseDefaultDeleteFileModal(app).
		SetDefaultDeleteFolderModal(app)

	return explorer
}

func (explorer *Explorer) SetName(name string) *Explorer {
	explorer.name = name
	return explorer
}

func (explorer *Explorer) SetPath(path string) *Explorer {
	explorer.path = path
	return explorer
}

func (explorer *Explorer) SetPadding(padding int) *Explorer {
	explorer.padding = min(35, padding)
	return explorer
}

func (explorer *Explorer) SetExpanded(expanded bool) *Explorer {
	explorer.expanded = expanded
	return explorer
}

func (explorer *Explorer) SetHierarchy(hierarchy bool) *Explorer {
	explorer.hierarchy = hierarchy
	return explorer
}

func (explorer *Explorer) SetTree(tree *tree.Tree) *Explorer {
	explorer.tree = tree.
		View()

	return explorer
}

func (explorer *Explorer) UseDefaultTree() *Explorer {
	design := tree.
		NewDesign().
		SetRootColor(0x11111b)

	root := tree.
		NewRoot().
		SetName(explorer.name).
		SetPath(explorer.path).
		SetPadding(explorer.padding).
		SetExpanded(explorer.expanded).
		SetSelectedStyle(0x181825, 0xcdd6f4).
		SetUnselectedStyle(0x11111b, 0xcdd6f4)

	explorer.tree = tree.
		NewTree().
		SetRoot(root).
		SetDesign(design).
		SetHierarchy(explorer.hierarchy).
		HandleSelect(func(node *tree.Node, treeNode *tview.TreeNode) {
			if node.IsDir() {
				tree.Toggle(node, treeNode)
				return
			}
		}).
		HandleDimension(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x, y + 1, width, height
		}).
		View()

	return explorer
}

func (explorer *Explorer) SetAddFileModal(addFileModal *modal.Modal[*modal.InputModal]) *Explorer {
	explorer.addFileModal = addFileModal.
		View()

	return explorer
}

func (explorer *Explorer) UseDefaultAddFileModal(app *tview.Application) *Explorer {
	inputModal := modal.
		NewInputModal().
		SetDefaultMessage("Message").
		SetDefaultSuccessButton("Success").
		SetDefaultFailureButton("Failure").
		SetDefaultInput("", "Enter A Message").
		HandleDefaultSwitch(func(index int) (*tview.Application, int) {
			if index == modal.InputModalSuccessButton {
				return app, modal.InputModalFailureButton
			}

			if index == modal.InputModalFailureButton {
				return app, modal.InputModalInput
			}

			if index == modal.InputModalInput {
				return app, modal.InputModalSuccessButton
			}

			return app, modal.InputModalSuccessButton
		})

	explorer.addFileModal = modal.
		NewModal(inputModal).
		SetBackgroundColor(0x11111b).
		View()

	return explorer
}

func (explorer *Explorer) SetAddFolderModal(addFolderModal *modal.Modal[*modal.InputModal]) *Explorer {
	explorer.addFolderModal = addFolderModal.
		View()

	return explorer
}

func (explorer *Explorer) UseDefaultAddFolderModal(app *tview.Application) *Explorer {
	inputModal := modal.
		NewInputModal().
		SetDefaultMessage("Message").
		SetDefaultSuccessButton("Success").
		SetDefaultFailureButton("Failure").
		SetDefaultInput("", "Enter A Message").
		HandleDefaultSwitch(func(index int) (*tview.Application, int) {
			if index == modal.InputModalSuccessButton {
				return app, modal.InputModalFailureButton
			}

			if index == modal.InputModalFailureButton {
				return app, modal.InputModalInput
			}

			if index == modal.InputModalInput {
				return app, modal.InputModalSuccessButton
			}

			return app, modal.InputModalSuccessButton
		})

	explorer.addFolderModal = modal.
		NewModal(inputModal).
		SetBackgroundColor(0x11111b).
		View()

	return explorer
}

func (explorer *Explorer) SetDeleteFileModal(deleteFileModal *modal.Modal[*modal.MessageModal]) *Explorer {
	explorer.deleteFileModal = deleteFileModal.
		View()

	return explorer
}

func (explorer *Explorer) UseDefaultDeleteFileModal(app *tview.Application) *Explorer {
	messageModal := modal.
		NewMessageModal().
		SetDefaultMessage("Message").
		SetDefaultSuccessButton("Success").
		SetDefaultFailureButton("Failure").
		HandleDefaultSwitch(func(index int) (*tview.Application, int) {
			if index == modal.MessageModalSuccessButton {
				return app, modal.MessageModalFailureButton
			}

			return app, modal.MessageModalSuccessButton
		})

	explorer.deleteFileModal = modal.
		NewModal(messageModal).
		SetGap(1, 0).
		SetBackgroundColor(0x11111b).
		View()

	return explorer
}

func (explorer *Explorer) UseDeleteFolderModal(deleteFolderModal *modal.Modal[*modal.MessageModal]) *Explorer {
	explorer.deleteFolderModal = deleteFolderModal.
		View()

	return explorer
}

func (explorer *Explorer) SetDefaultDeleteFolderModal(app *tview.Application) *Explorer {
	messageModal := modal.
		NewMessageModal().
		SetDefaultMessage("Message").
		SetDefaultSuccessButton("Success").
		SetDefaultFailureButton("Failure").
		HandleDefaultSwitch(func(index int) (*tview.Application, int) {
			if index == modal.MessageModalSuccessButton {
				return app, modal.MessageModalFailureButton
			}

			return app, modal.MessageModalSuccessButton
		})

	explorer.deleteFolderModal = modal.
		NewModal(messageModal).
		SetGap(1, 0).
		SetBackgroundColor(0x11111b).
		View()

	return explorer
}

func (explorer *Explorer) GetName() string {
	return explorer.name
}

func (explorer *Explorer) GetPath() string {
	return explorer.path
}

func (explorer *Explorer) GetPadding() int {
	return explorer.padding
}

func (explorer *Explorer) IsExpanded() bool {
	return explorer.expanded
}

func (explorer *Explorer) IsHierarchical() bool {
	return explorer.hierarchy
}

func (explorer *Explorer) GetTree() *tview.TreeView {
	return explorer.tree
}

func (explorer *Explorer) GetAddFileModal() *tview.Flex {
	return explorer.addFileModal
}

func (explorer *Explorer) GetAddFolderModal() *tview.Flex {
	return explorer.addFolderModal
}

func (explorer *Explorer) GetDeleteFileModal() *tview.Flex {
	return explorer.deleteFileModal
}

func (explorer *Explorer) GetDeleteFolderModal() *tview.Flex {
	return explorer.deleteFolderModal
}
