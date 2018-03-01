package resource

import (
	"log"
	"reflect"

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

/*
Interactive Resource Loader. This object is returned by ResourceLoader when performing an interactive load. It allows to load with high granularity, so this is mainly useful for displaying load bars/percentages.
*/
type ResourceInteractiveLoader struct {
	Reference
}

func (o *ResourceInteractiveLoader) BaseClass() string {
	return "ResourceInteractiveLoader"
}

/*
   Return the loaded resource (only if loaded). Otherwise, returns null.
*/
func (o *ResourceInteractiveLoader) GetResource() *Resource {
	log.Println("Calling ResourceInteractiveLoader.GetResource()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_resource", goArguments, "*Resource")

	returnValue := goRet.Interface().(*Resource)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the load stage. The total amount of stages can be queried with [method get_stage_count]
*/
func (o *ResourceInteractiveLoader) GetStage() gdnative.Int {
	log.Println("Calling ResourceInteractiveLoader.GetStage()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_stage", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return the total amount of stages (calls to [method poll]) needed to completely load this resource.
*/
func (o *ResourceInteractiveLoader) GetStageCount() gdnative.Int {
	log.Println("Calling ResourceInteractiveLoader.GetStageCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_stage_count", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Poll the load. If OK is returned, this means poll will have to be called again. If ERR_FILE_EOF is returned, them the load has finished and the resource can be obtained by calling [method get_resource].
*/
func (o *ResourceInteractiveLoader) Poll() gdnative.Int {
	log.Println("Calling ResourceInteractiveLoader.Poll()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "poll", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *ResourceInteractiveLoader) Wait() gdnative.Int {
	log.Println("Calling ResourceInteractiveLoader.Wait()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "wait", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   ResourceInteractiveLoaderImplementer is an interface for ResourceInteractiveLoader objects.
*/
type ResourceInteractiveLoaderImplementer interface {
	Class
}