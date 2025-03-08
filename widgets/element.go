package widgets

import (
	"github.com/rivo/tview"
)

type Element struct {
	entity tview.Primitive
	active bool
}

func (element Element) Entity() (entity tview.Primitive) {
	return element.entity
}

func (element Element) SetEntity(entity tview.Primitive) Element {
	element.entity = entity
	return element
}

func (element Element) Active() (active bool) {
	return active
}

func (element Element) SetActive(active bool) Element {
	element.active = active
	return element
}
