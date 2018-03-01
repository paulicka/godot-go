package rigidbody2d

import (
	"log"
	"reflect"
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
This node implements simulated 2D physics. You do not control a RigidBody2D directly. Instead you apply forces to it (gravity, impulses, etc.) and the physics simulation calculates the resulting movement based on its mass, friction, and other physical properties. A RigidBody2D has 4 behavior [member mode]s: Rigid, Static, Character, and Kinematic. [b]Note:[/b] You should not change a RigidBody2D's [code]position[/code] or [code]linear_velocity[/code] every frame or even very often. If you need to directly affect the body's state, use [method _integrate_forces], which allows you to directly access the physics state. If you need to override the default physics behavior, you can write a custom force integration. See [member custom_integrator].
*/
type RigidBody2D struct {
	PhysicsBody2D
}

func (o *RigidBody2D) BaseClass() string {
	return "RigidBody2D"
}

/*
   Undocumented
*/
func (o *RigidBody2D) X_BodyEnterTree(arg0 gdnative.Int) {
	log.Println("Calling RigidBody2D.X_BodyEnterTree()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_body_enter_tree", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) X_BodyExitTree(arg0 gdnative.Int) {
	log.Println("Calling RigidBody2D.X_BodyExitTree()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_body_exit_tree", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) X_DirectStateChanged(arg0 *Object) {
	log.Println("Calling RigidBody2D.X_DirectStateChanged()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(arg0)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_direct_state_changed", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Allows you to read and safely modify the simulation state for the object. Use this instead of [Node._physics_process] if you need to directly change the body's [code]position[/code] or other physics properties. By default it works in addition to the usual physics behavior, but [member custom_integrator] allows you to disable the default behavior and write custom force integration for a body.
*/
func (o *RigidBody2D) X_IntegrateForces(state *Physics2DDirectBodyState) {
	log.Println("Calling RigidBody2D.X_IntegrateForces()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(state)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "_integrate_forces", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Adds a positioned force to the body. Both the force and the offset from the body origin are in global coordinates.
*/
func (o *RigidBody2D) AddForce(offset *Vector2, force *Vector2) {
	log.Println("Calling RigidBody2D.AddForce()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(offset)
	goArguments[1] = reflect.ValueOf(force)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "add_force", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Applies a positioned impulse to the body (which will be affected by the body mass and shape). This is the equivalent of hitting a billiard ball with a cue: a force that is applied instantaneously. Both the impulse and the offset from the body origin are in global coordinates.
*/
func (o *RigidBody2D) ApplyImpulse(offset *Vector2, impulse *Vector2) {
	log.Println("Calling RigidBody2D.ApplyImpulse()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(offset)
	goArguments[1] = reflect.ValueOf(impulse)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "apply_impulse", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) GetAngularDamp() gdnative.Float {
	log.Println("Calling RigidBody2D.GetAngularDamp()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_angular_damp", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody2D) GetAngularVelocity() gdnative.Float {
	log.Println("Calling RigidBody2D.GetAngularVelocity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_angular_velocity", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody2D) GetAppliedForce() *Vector2 {
	log.Println("Calling RigidBody2D.GetAppliedForce()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_applied_force", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody2D) GetAppliedTorque() gdnative.Float {
	log.Println("Calling RigidBody2D.GetAppliedTorque()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_applied_torque", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody2D) GetBounce() gdnative.Float {
	log.Println("Calling RigidBody2D.GetBounce()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_bounce", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Returns a list of the bodies colliding with this one. Use [member contacts_reported] to set the maximum number reported. You must also set [member contact_monitor] to [code]true[/code]. Note that the result of this test is not immediate after moving objects. For performance, list of collisions is updated once per frame and before the physics step. Consider using signals instead.
*/
func (o *RigidBody2D) GetCollidingBodies() *Array {
	log.Println("Calling RigidBody2D.GetCollidingBodies()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_colliding_bodies", goArguments, "*Array")

	returnValue := goRet.Interface().(*Array)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody2D) GetContinuousCollisionDetectionMode() gdnative.Int {
	log.Println("Calling RigidBody2D.GetContinuousCollisionDetectionMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_continuous_collision_detection_mode", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody2D) GetFriction() gdnative.Float {
	log.Println("Calling RigidBody2D.GetFriction()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_friction", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody2D) GetGravityScale() gdnative.Float {
	log.Println("Calling RigidBody2D.GetGravityScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_gravity_scale", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody2D) GetInertia() gdnative.Float {
	log.Println("Calling RigidBody2D.GetInertia()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_inertia", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody2D) GetLinearDamp() gdnative.Float {
	log.Println("Calling RigidBody2D.GetLinearDamp()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_linear_damp", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody2D) GetLinearVelocity() *Vector2 {
	log.Println("Calling RigidBody2D.GetLinearVelocity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_linear_velocity", goArguments, "*Vector2")

	returnValue := goRet.Interface().(*Vector2)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody2D) GetMass() gdnative.Float {
	log.Println("Calling RigidBody2D.GetMass()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_mass", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody2D) GetMaxContactsReported() gdnative.Int {
	log.Println("Calling RigidBody2D.GetMaxContactsReported()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_max_contacts_reported", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody2D) GetMode() gdnative.Int {
	log.Println("Calling RigidBody2D.GetMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_mode", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody2D) GetWeight() gdnative.Float {
	log.Println("Calling RigidBody2D.GetWeight()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_weight", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody2D) IsAbleToSleep() gdnative.Bool {
	log.Println("Calling RigidBody2D.IsAbleToSleep()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_able_to_sleep", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody2D) IsContactMonitorEnabled() gdnative.Bool {
	log.Println("Calling RigidBody2D.IsContactMonitorEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_contact_monitor_enabled", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody2D) IsSleeping() gdnative.Bool {
	log.Println("Calling RigidBody2D.IsSleeping()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_sleeping", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody2D) IsUsingCustomIntegrator() gdnative.Bool {
	log.Println("Calling RigidBody2D.IsUsingCustomIntegrator()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_using_custom_integrator", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *RigidBody2D) SetAngularDamp(angularDamp gdnative.Float) {
	log.Println("Calling RigidBody2D.SetAngularDamp()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(angularDamp)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_angular_damp", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) SetAngularVelocity(angularVelocity gdnative.Float) {
	log.Println("Calling RigidBody2D.SetAngularVelocity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(angularVelocity)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_angular_velocity", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) SetAppliedForce(force *Vector2) {
	log.Println("Calling RigidBody2D.SetAppliedForce()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(force)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_applied_force", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) SetAppliedTorque(torque gdnative.Float) {
	log.Println("Calling RigidBody2D.SetAppliedTorque()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(torque)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_applied_torque", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Sets the body's velocity on the given axis. The velocity in the given vector axis will be set as the given vector length. This is useful for jumping behavior.
*/
func (o *RigidBody2D) SetAxisVelocity(axisVelocity *Vector2) {
	log.Println("Calling RigidBody2D.SetAxisVelocity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(axisVelocity)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_axis_velocity", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) SetBounce(bounce gdnative.Float) {
	log.Println("Calling RigidBody2D.SetBounce()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(bounce)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_bounce", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) SetCanSleep(ableToSleep gdnative.Bool) {
	log.Println("Calling RigidBody2D.SetCanSleep()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(ableToSleep)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_can_sleep", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) SetContactMonitor(enabled gdnative.Bool) {
	log.Println("Calling RigidBody2D.SetContactMonitor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enabled)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_contact_monitor", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) SetContinuousCollisionDetectionMode(mode gdnative.Int) {
	log.Println("Calling RigidBody2D.SetContinuousCollisionDetectionMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_continuous_collision_detection_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) SetFriction(friction gdnative.Float) {
	log.Println("Calling RigidBody2D.SetFriction()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(friction)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_friction", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) SetGravityScale(gravityScale gdnative.Float) {
	log.Println("Calling RigidBody2D.SetGravityScale()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(gravityScale)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_gravity_scale", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) SetInertia(inertia gdnative.Float) {
	log.Println("Calling RigidBody2D.SetInertia()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(inertia)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_inertia", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) SetLinearDamp(linearDamp gdnative.Float) {
	log.Println("Calling RigidBody2D.SetLinearDamp()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(linearDamp)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_linear_damp", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) SetLinearVelocity(linearVelocity *Vector2) {
	log.Println("Calling RigidBody2D.SetLinearVelocity()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(linearVelocity)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_linear_velocity", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) SetMass(mass gdnative.Float) {
	log.Println("Calling RigidBody2D.SetMass()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mass)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_mass", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) SetMaxContactsReported(amount gdnative.Int) {
	log.Println("Calling RigidBody2D.SetMaxContactsReported()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(amount)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_max_contacts_reported", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) SetMode(mode gdnative.Int) {
	log.Println("Calling RigidBody2D.SetMode()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(mode)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_mode", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) SetSleeping(sleeping gdnative.Bool) {
	log.Println("Calling RigidBody2D.SetSleeping()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(sleeping)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_sleeping", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) SetUseCustomIntegrator(enable gdnative.Bool) {
	log.Println("Calling RigidBody2D.SetUseCustomIntegrator()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(enable)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_use_custom_integrator", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *RigidBody2D) SetWeight(weight gdnative.Float) {
	log.Println("Calling RigidBody2D.SetWeight()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(weight)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_weight", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Returns [code]true[/code] if a collision would result from moving in the given vector. [code]margin[/code] increases the size of the shapes involved in the collision detection, and [code]result[/code] is an object of type [Physics2DTestMotionResult], which contains additional information about the collision (should there be one).
*/
func (o *RigidBody2D) TestMotion(motion *Vector2, margin gdnative.Float, result *Physics2DTestMotionResult) gdnative.Bool {
	log.Println("Calling RigidBody2D.TestMotion()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 3, 3)
	goArguments[0] = reflect.ValueOf(motion)
	goArguments[1] = reflect.ValueOf(margin)
	goArguments[2] = reflect.ValueOf(result)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "test_motion", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   RigidBody2DImplementer is an interface for RigidBody2D objects.
*/
type RigidBody2DImplementer interface {
	Class
}