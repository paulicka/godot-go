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

//func NewNavigationMeshInstanceFromPointer(ptr gdnative.Pointer) NavigationMeshInstance {
func NewNavigationMeshInstanceFromPointer(ptr gdnative.Pointer) NavigationMeshInstance {
	owner := gdnative.NewObjectFromPointer(ptr)
	obj := NavigationMeshInstance{}
	obj.SetBaseObject(owner)

	return obj
}

/*

 */
type NavigationMeshInstance struct {
	Spatial
	owner gdnative.Object
}

func (o *NavigationMeshInstance) BaseClass() string {
	return "NavigationMeshInstance"
}

/*
        Undocumented
	Args: [], Returns: NavigationMesh
*/
func (o *NavigationMeshInstance) GetNavigationMesh() NavigationMesh {
	//log.Println("Calling NavigationMeshInstance.GetNavigationMesh()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NavigationMeshInstance", "get_navigation_mesh")

	// Call the parent method.
	// NavigationMesh
	retPtr := gdnative.NewEmptyObject()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := NewNavigationMeshFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [], Returns: bool
*/
func (o *NavigationMeshInstance) IsEnabled() gdnative.Bool {
	//log.Println("Calling NavigationMeshInstance.IsEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 0, 0)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NavigationMeshInstance", "is_enabled")

	// Call the parent method.
	// bool
	retPtr := gdnative.NewEmptyBool()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

	// If we have a return type, convert it from a pointer into its actual object.
	ret := gdnative.NewBoolFromPointer(retPtr)
	//log.Println("  Got return value: ", ret)
	return ret
}

/*
        Undocumented
	Args: [{ false enabled bool}], Returns: void
*/
func (o *NavigationMeshInstance) SetEnabled(enabled gdnative.Bool) {
	//log.Println("Calling NavigationMeshInstance.SetEnabled()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromBool(enabled)

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NavigationMeshInstance", "set_enabled")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}

/*
        Undocumented
	Args: [{ false navmesh NavigationMesh}], Returns: void
*/
func (o *NavigationMeshInstance) SetNavigationMesh(navmesh NavigationMesh) {
	//log.Println("Calling NavigationMeshInstance.SetNavigationMesh()")

	// Build out the method's arguments
	ptrArguments := make([]gdnative.Pointer, 1, 1)
	ptrArguments[0] = gdnative.NewPointerFromObject(navmesh.GetBaseObject())

	// Get the method bind
	methodBind := gdnative.NewMethodBind("NavigationMeshInstance", "set_navigation_mesh")

	// Call the parent method.
	// void
	retPtr := gdnative.NewEmptyVoid()
	gdnative.MethodBindPtrCall(methodBind, o.GetBaseObject(), ptrArguments, retPtr)

}