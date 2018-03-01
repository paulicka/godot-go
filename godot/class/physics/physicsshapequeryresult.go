package physics

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

 */
type PhysicsShapeQueryResult struct {
	Reference
}

func (o *PhysicsShapeQueryResult) BaseClass() string {
	return "PhysicsShapeQueryResult"
}

/*

 */
func (o *PhysicsShapeQueryResult) GetResultCount() gdnative.Int {
	log.Println("Calling PhysicsShapeQueryResult.GetResultCount()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_result_count", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *PhysicsShapeQueryResult) GetResultObject(idx gdnative.Int) *Object {
	log.Println("Calling PhysicsShapeQueryResult.GetResultObject()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_result_object", goArguments, "*Object")

	returnValue := goRet.Interface().(*Object)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *PhysicsShapeQueryResult) GetResultObjectId(idx gdnative.Int) gdnative.Int {
	log.Println("Calling PhysicsShapeQueryResult.GetResultObjectId()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_result_object_id", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *PhysicsShapeQueryResult) GetResultObjectShape(idx gdnative.Int) gdnative.Int {
	log.Println("Calling PhysicsShapeQueryResult.GetResultObjectShape()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_result_object_shape", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *PhysicsShapeQueryResult) GetResultRid(idx gdnative.Int) *RID {
	log.Println("Calling PhysicsShapeQueryResult.GetResultRid()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(idx)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_result_rid", goArguments, "*RID")

	returnValue := goRet.Interface().(*RID)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   PhysicsShapeQueryResultImplementer is an interface for PhysicsShapeQueryResult objects.
*/
type PhysicsShapeQueryResultImplementer interface {
	Class
}