package class

import (
	"github.com/shadowapex/godot-go/gdnative"
)

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

//func NewVScrollBarFromPointer(ptr gdnative.Pointer) VScrollBar {
func NewVScrollBarFromPointer(ptr gdnative.Pointer) VScrollBar {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := VScrollBar{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type VScrollBar struct {
	ScrollBar
	owner gdnative.Object
}

func (o *VScrollBar) BaseClass() string {
	return "VScrollBar"
}

// SetBaseObject will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *VScrollBar) SetBaseObject(object gdnative.Object) {
	o.owner = object
}

func (o *VScrollBar) GetBaseObject() gdnative.Object {
	return o.owner
}