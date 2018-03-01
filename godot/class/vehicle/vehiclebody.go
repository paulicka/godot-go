package vehicle

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
This nodes implements all the physics logic needed to simulate a car. It is based on the raycast vehicle system commonly found in physics engines. You will need to add a [CollisionShape] for the main body of your vehicle and add [VehicleWheel] nodes for the wheels. You should also add a [MeshInstance] to this node for the 3D model of your car but this model should not include meshes for the wheels. You should control the vehicle by using the [member brake], [member engine_force], and [member steering] properties and not change the position or orientation of this node directly. Note that the origin point of your VehicleBody will determine the center of gravity of your vehicle so it is better to keep this low and move the [CollisionShape] and [MeshInstance] upwards.
*/
type VehicleBody struct {
	RigidBody
}

func (o *VehicleBody) BaseClass() string {
	return "VehicleBody"
}

/*
   Undocumented
*/
func (o *VehicleBody) GetBrake() gdnative.Float {
	log.Println("Calling VehicleBody.GetBrake()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_brake", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VehicleBody) GetEngineForce() gdnative.Float {
	log.Println("Calling VehicleBody.GetEngineForce()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_engine_force", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VehicleBody) GetSteering() gdnative.Float {
	log.Println("Calling VehicleBody.GetSteering()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_steering", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *VehicleBody) SetBrake(brake gdnative.Float) {
	log.Println("Calling VehicleBody.SetBrake()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(brake)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_brake", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *VehicleBody) SetEngineForce(engineForce gdnative.Float) {
	log.Println("Calling VehicleBody.SetEngineForce()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(engineForce)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_engine_force", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *VehicleBody) SetSteering(steering gdnative.Float) {
	log.Println("Calling VehicleBody.SetSteering()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(steering)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_steering", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   VehicleBodyImplementer is an interface for VehicleBody objects.
*/
type VehicleBodyImplementer interface {
	Class
}