package raycast2d

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
A RayCast represents a line from its origin to its destination position, [code]cast_to[/code]. It is used to query the 2D space in order to find the closest object along the path of the ray. RayCast2D can ignore some objects by adding them to the exception list via [code]add_exception[/code], by setting proper filtering with collision layers, or by filtering object types with type masks. Only enabled raycasts will be able to query the space and report collisions. RayCast2D calculates intersection every physics frame (see [Node]), and the result is cached so it can be used later until the next frame. If multiple queries are required between physics frames (or during the same frame) use [method force_raycast_update] after adjusting the raycast.
*/
type RayCast2D struct {
	Node2D
}

func (o *RayCast2D) BaseClass() string {
	return "RayCast2D"
}

/*
   Adds a collision exception so the ray does not report collisions with the specified node.
*/
func (o *RayCast2D) AddException(node *Object) {
	log.Println("Calling RayCast2D.AddException()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(node)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "add_exception", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Adds a collision exception so the ray does not report collisions with the specified [RID].
*/
func (o *RayCast2D) AddExceptionRid(rid *RID) {
	log.Println("Calling RayCast2D.AddExceptionRid()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(rid)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "add_exception_rid", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Removes all collision exceptions for this ray.
*/
func (o *RayCast2D) ClearExceptions() {
	log.Println("Calling RayCast2D.ClearExceptions()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "clear_exceptions", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Updates the collision information for the ray. Use this method to update the collision information immediately instead of waiting for the next [code]_physics_process[/code] call, for example if the ray or its parent has changed state. Note: [code]enabled == true[/code] is not required for this to work.
*/
func (o *RayCast2D) ForceRaycastUpdate() {
	log.Println("Calling RayCast2D.ForceRaycastUpdate()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "force_raycast_update", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RayCast2D) GetCastTo() *Vector2 {
	log.Println("Calling RayCast2D.GetCastTo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_cast_to", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the closest object the ray is pointing to. Note that this does not consider the length of the ray, so you must also use [method is_colliding] to check if the object returned is actually colliding with the ray. Example: [codeblock] if RayCast2D.is_colliding(): var collider = RayCast2D.get_collider() [/codeblock]
*/
func (o *RayCast2D) GetCollider() *Object {
	log.Println("Calling RayCast2D.GetCollider()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collider", goArguments, "*Object")

	returnValue := goRet.Interface().(*Object)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the collision shape of the closest object the ray is pointing to. Note that this does not consider the length of the ray, so you must also use [method is_colliding] to check if the object returned is actually colliding with the ray. Example: [codeblock] if RayCast2D.is_colliding(): var shape = RayCast2D.get_collider_shape() [/codeblock]
*/
func (o *RayCast2D) GetColliderShape() gdnative.Int {
	log.Println("Calling RayCast2D.GetColliderShape()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collider_shape", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RayCast2D) GetCollisionMask() gdnative.Int {
	log.Println("Calling RayCast2D.GetCollisionMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collision_mask", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return an individual bit on the collision mask.
*/
func (o *RayCast2D) GetCollisionMaskBit(bit gdnative.Int) gdnative.Bool {
	log.Println("Calling RayCast2D.GetCollisionMaskBit()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(bit)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collision_mask_bit", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the normal of the intersecting object's shape at the collision point.
*/
func (o *RayCast2D) GetCollisionNormal() *Vector2 {
	log.Println("Calling RayCast2D.GetCollisionNormal()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collision_normal", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns the collision point at which the ray intersects the closest object. Note: this point is in the [b]global[/b] coordinate system.
*/
func (o *RayCast2D) GetCollisionPoint() *Vector2 {
	log.Println("Calling RayCast2D.GetCollisionPoint()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_collision_point", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RayCast2D) GetExcludeParentBody() gdnative.Bool {
	log.Println("Calling RayCast2D.GetExcludeParentBody()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_exclude_parent_body", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Return whether the closest object the ray is pointing to is colliding with the vector (considering the vector length).
*/
func (o *RayCast2D) IsColliding() gdnative.Bool {
	log.Println("Calling RayCast2D.IsColliding()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_colliding", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RayCast2D) IsEnabled() gdnative.Bool {
	log.Println("Calling RayCast2D.IsEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_enabled", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Removes a collision exception so the ray does report collisions with the specified node.
*/
func (o *RayCast2D) RemoveException(node *Object) {
	log.Println("Calling RayCast2D.RemoveException()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(node)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "remove_exception", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Removes a collision exception so the ray does report collisions with the specified [RID].
*/
func (o *RayCast2D) RemoveExceptionRid(rid *RID) {
	log.Println("Calling RayCast2D.RemoveExceptionRid()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(rid)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "remove_exception_rid", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RayCast2D) SetCastTo(localPoint *Vector2) {
	log.Println("Calling RayCast2D.SetCastTo()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(localPoint)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_cast_to", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RayCast2D) SetCollisionMask(mask gdnative.Int) {
	log.Println("Calling RayCast2D.SetCollisionMask()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mask)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_collision_mask", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Set/clear individual bits on the collision mask. This makes selecting the areas scanned easier.
*/
func (o *RayCast2D) SetCollisionMaskBit(bit gdnative.Int, value gdnative.Bool) {
	log.Println("Calling RayCast2D.SetCollisionMaskBit()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(bit)
	goArguments[1] = reflect.ValueOf(value)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_collision_mask_bit", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RayCast2D) SetEnabled(enabled gdnative.Bool) {
	log.Println("Calling RayCast2D.SetEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enabled)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_enabled", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RayCast2D) SetExcludeParentBody(mask gdnative.Bool) {
	log.Println("Calling RayCast2D.SetExcludeParentBody()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mask)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_exclude_parent_body", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   RayCast2DImplementer is an interface for RayCast2D objects.
*/
type RayCast2DImplementer interface {
	Class
}