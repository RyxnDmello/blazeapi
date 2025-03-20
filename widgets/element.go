package widgets

import (
	"github.com/rivo/tview"
)

type Element struct {
	entity tview.Primitive
	active bool
}

// Entity returns the entity of the Element type.
//
// The entity value is retrieved from the Element instance. This function
// provides access to the current tview.Primitive.
//
// # Returns
//
//	entity tview.Primitive
//
// # Usage
//
//	entity := element.Entity()
func (element Element) Entity() (entity tview.Primitive) {
	return element.entity
}

// SetEntity sets the entity of the Element type.
//
// The entity value is updated to the specified tview.Primitive. Method chaining
// is supported by returning the modified Element.
//
// # Parameters
//
//	entity tview.Primitive
//
// # Returns
//
//	Element
//
// # Usage
//
//	element = element.SetEntity(nil)
func (element Element) SetEntity(entity tview.Primitive) Element {
	element.entity = entity
	return element
}

// Active returns the active state of the Element type.
//
// The active state is retrieved from the Element instance. This function
// provides access to the current active setting.
//
// # Returns
//
//	active bool
//
// # Usage
//
//	active := element.Active()
func (element Element) Active() (active bool) {
	return element.active
}

// SetActive sets the active state of the Element type.
//
// The active state is updated to the specified boolean. Method chaining is
// supported by returning the modified Element.
//
// # Parameters
//
//	active bool
//
// # Returns
//
//	Element
//
// # Usage
//
//	element = element.SetActive(false)
func (element Element) SetActive(active bool) Element {
	element.active = active
	return element
}
