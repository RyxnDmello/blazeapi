package project

import (
	"blazeapi/core"
	"blazeapi/query"
	"blazeapi/response"

	"blazeapi/widgets/modal"
	"blazeapi/widgets/tree"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

func InitializeProject(app *tview.Application, query *query.Query, response *response.Response) (project *tview.TreeView, createFileModal *tview.Flex, createFolderModal *tview.Flex, deleteNodeModal *tview.Flex) {
	design := tree.
		NewDesign().
		SetRootColor(0x11111b)

	root := tree.
		NewRoot().
		SetPath("./test").
		SetName("test").
		SetPadding(40).
		SetExpanded(true).
		SetSelectedStyle(0x181825, 0xcdd6f4).
		SetUnselectedStyle(0x11111b, 0xcdd6f4)

	project = tree.
		NewTree().
		SetRoot(root).
		SetDesign(design).
		SetHierarchy(true).
		HandleSelect(func(node *tree.Node, treeNode *tview.TreeNode) {
			if node.IsDir() {
				tree.Toggle(node, treeNode)
				return
			}

			api := core.NewApi().Read(node.GetPath())
			query.SetMethod(api.Method)
			query.SetBody(api.Body)
			query.SetUrl(api.Url)
		}).
		HandleDimension(func(screen tcell.Screen, x, y, width, height int) (int, int, int, int) {
			return x, y + 1, width, height
		}).
		View()

	createFolderModal = messageModal(app, project)
	createFileModal = inputModal(app, project)
	deleteNodeModal = bodyModal(app, project)

	return project, createFileModal, createFolderModal, deleteNodeModal
}

func messageModal(app *tview.Application, _ *tview.TreeView) *tview.Flex {
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

	modal := modal.
		NewModal(messageModal).
		SetGap(1, 0).
		SetBackgroundColor(0x11111b).
		View()

	return modal
}

func inputModal(app *tview.Application, _ *tview.TreeView) *tview.Flex {
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

	modal := modal.
		NewModal(inputModal).
		SetBackgroundColor(0x11111b).
		View()

	return modal
}

func bodyModal(app *tview.Application, _ *tview.TreeView) *tview.Flex {
	areaModal := modal.
		NewAreaModal().
		SetDefaultMessage("Message").
		SetDefaultSuccessButton("Success").
		SetDefaultFailureButton("Failure").
		SetDefaultArea("", "Enter Body", 10, 0).
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

	modal := modal.
		NewModal(areaModal).
		SetBackgroundColor(0x11111b).
		View()

	return modal
}
