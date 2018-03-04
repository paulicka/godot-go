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

//func NewAudioEffectFromPointer(ptr gdnative.Pointer) AudioEffect {
func NewAudioEffectFromPointer(ptr gdnative.Pointer) AudioEffect {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := AudioEffect{}
	obj.SetBaseObject(owner)

	return obj
}

/*
Base resource for audio bus. Applies an audio effect on the bus that the resource is applied on.
*/
type AudioEffect struct {
	Resource
	owner gdnative.Object
}

func (o *AudioEffect) BaseClass() string {
	return "AudioEffect"
}

// SetBaseObject will internally set the Godot object inside the struct.
// This is used to call parent methods.
func (o *AudioEffect) SetBaseObject(object gdnative.Object) {
	o.owner = object
}

func (o *AudioEffect) GetBaseObject() gdnative.Object {
	return o.owner
}