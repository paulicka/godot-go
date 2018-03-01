package stylebox

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
This stylebox can be used to achieve all kinds of looks without the need of a texture. Those properties are customizable: - Color - Border width (individual width for each border) - Rounded corners (individual radius for each corner) - Shadow About corner radius: Setting corner radius to high values is allowed. As soon as corners would overlap the stylebox will switch to a relative system. Example: [codeblock] height = 30 corner_radius_top_left = 50 corner_radius_bottom_left = 100 [/codeblock] The relative system now would take the 1:2 ratio of the two left corners to calculate the actual corner width. Both corners added will [b]never[/b] be more than the height. Result: [codeblock] corner_radius_top_left: 10 corner_radius_bottom_left: 20 [/codeblock]
*/
type StyleBoxFlat struct {
	StyleBox
}

func (o *StyleBoxFlat) BaseClass() string {
	return "StyleBoxFlat"
}

/*
   Undocumented
*/
func (o *StyleBoxFlat) GetAaSize() gdnative.Int {
	log.Println("Calling StyleBoxFlat.GetAaSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_aa_size", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) GetBgColor() *Color {
	log.Println("Calling StyleBoxFlat.GetBgColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_bg_color", goArguments, "*Color")

	returnValue := goRet.Interface().(*Color)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) GetBorderBlend() gdnative.Bool {
	log.Println("Calling StyleBoxFlat.GetBorderBlend()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_border_blend", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) GetBorderColor() *Color {
	log.Println("Calling StyleBoxFlat.GetBorderColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_border_color", goArguments, "*Color")

	returnValue := goRet.Interface().(*Color)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) GetBorderWidth(margin gdnative.Int) gdnative.Int {
	log.Println("Calling StyleBoxFlat.GetBorderWidth()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(margin)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_border_width", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*

 */
func (o *StyleBoxFlat) GetBorderWidthMin() gdnative.Int {
	log.Println("Calling StyleBoxFlat.GetBorderWidthMin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_border_width_min", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) GetCornerDetail() gdnative.Int {
	log.Println("Calling StyleBoxFlat.GetCornerDetail()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_corner_detail", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) GetCornerRadius(corner gdnative.Int) gdnative.Int {
	log.Println("Calling StyleBoxFlat.GetCornerRadius()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(corner)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_corner_radius", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) GetExpandMargin(margin gdnative.Int) gdnative.Float {
	log.Println("Calling StyleBoxFlat.GetExpandMargin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(margin)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_expand_margin", goArguments, "gdnative.Float")

	returnValue := goRet.Interface().(gdnative.Float)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) GetShadowColor() *Color {
	log.Println("Calling StyleBoxFlat.GetShadowColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_shadow_color", goArguments, "*Color")

	returnValue := goRet.Interface().(*Color)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) GetShadowSize() gdnative.Int {
	log.Println("Calling StyleBoxFlat.GetShadowSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "get_shadow_size", goArguments, "gdnative.Int")

	returnValue := goRet.Interface().(gdnative.Int)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) IsAntiAliased() gdnative.Bool {
	log.Println("Calling StyleBoxFlat.IsAntiAliased()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_anti_aliased", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) IsDrawCenterEnabled() gdnative.Bool {
	log.Println("Calling StyleBoxFlat.IsDrawCenterEnabled()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 0, 0)

	// Call the parent method.

	goRet := o.callParentMethod(o.BaseClass(), "is_draw_center_enabled", goArguments, "gdnative.Bool")

	returnValue := goRet.Interface().(gdnative.Bool)

	log.Println("  Got return value: ", returnValue)
	return returnValue

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) SetAaSize(size gdnative.Int) {
	log.Println("Calling StyleBoxFlat.SetAaSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(size)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_aa_size", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) SetAntiAliased(antiAliased gdnative.Bool) {
	log.Println("Calling StyleBoxFlat.SetAntiAliased()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(antiAliased)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_anti_aliased", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) SetBgColor(color *Color) {
	log.Println("Calling StyleBoxFlat.SetBgColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(color)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_bg_color", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) SetBorderBlend(blend gdnative.Bool) {
	log.Println("Calling StyleBoxFlat.SetBorderBlend()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(blend)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_border_blend", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) SetBorderColor(color *Color) {
	log.Println("Calling StyleBoxFlat.SetBorderColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(color)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_border_color", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) SetBorderWidth(margin gdnative.Int, width gdnative.Int) {
	log.Println("Calling StyleBoxFlat.SetBorderWidth()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(margin)
	goArguments[1] = reflect.ValueOf(width)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_border_width", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetBorderWidthAll(width gdnative.Int) {
	log.Println("Calling StyleBoxFlat.SetBorderWidthAll()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(width)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_border_width_all", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) SetCornerDetail(detail gdnative.Int) {
	log.Println("Calling StyleBoxFlat.SetCornerDetail()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(detail)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_corner_detail", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) SetCornerRadius(corner gdnative.Int, radius gdnative.Int) {
	log.Println("Calling StyleBoxFlat.SetCornerRadius()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(corner)
	goArguments[1] = reflect.ValueOf(radius)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_corner_radius", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetCornerRadiusAll(radius gdnative.Int) {
	log.Println("Calling StyleBoxFlat.SetCornerRadiusAll()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(radius)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_corner_radius_all", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetCornerRadiusIndividual(radiusTopLeft gdnative.Int, radiusTopRight gdnative.Int, radiusBottomRight gdnative.Int, radiusBottomLeft gdnative.Int) {
	log.Println("Calling StyleBoxFlat.SetCornerRadiusIndividual()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 4, 4)
	goArguments[0] = reflect.ValueOf(radiusTopLeft)
	goArguments[1] = reflect.ValueOf(radiusTopRight)
	goArguments[2] = reflect.ValueOf(radiusBottomRight)
	goArguments[3] = reflect.ValueOf(radiusBottomLeft)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_corner_radius_individual", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) SetDrawCenter(drawCenter gdnative.Bool) {
	log.Println("Calling StyleBoxFlat.SetDrawCenter()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(drawCenter)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_draw_center", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) SetExpandMargin(margin gdnative.Int, size gdnative.Float) {
	log.Println("Calling StyleBoxFlat.SetExpandMargin()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 2, 2)
	goArguments[0] = reflect.ValueOf(margin)
	goArguments[1] = reflect.ValueOf(size)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_expand_margin", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetExpandMarginAll(size gdnative.Float) {
	log.Println("Calling StyleBoxFlat.SetExpandMarginAll()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(size)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_expand_margin_all", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*

 */
func (o *StyleBoxFlat) SetExpandMarginIndividual(sizeLeft gdnative.Float, sizeTop gdnative.Float, sizeRight gdnative.Float, sizeBottom gdnative.Float) {
	log.Println("Calling StyleBoxFlat.SetExpandMarginIndividual()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 4, 4)
	goArguments[0] = reflect.ValueOf(sizeLeft)
	goArguments[1] = reflect.ValueOf(sizeTop)
	goArguments[2] = reflect.ValueOf(sizeRight)
	goArguments[3] = reflect.ValueOf(sizeBottom)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_expand_margin_individual", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) SetShadowColor(color *Color) {
	log.Println("Calling StyleBoxFlat.SetShadowColor()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(color)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_shadow_color", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   Undocumented
*/
func (o *StyleBoxFlat) SetShadowSize(size gdnative.Int) {
	log.Println("Calling StyleBoxFlat.SetShadowSize()")

	// Build out the method's arguments
	goArguments := make([]reflect.Value, 1, 1)
	goArguments[0] = reflect.ValueOf(size)

	// Call the parent method.

	o.callParentMethod(o.BaseClass(), "set_shadow_size", goArguments, "")

	log.Println("  Function successfully completed.")

}

/*
   StyleBoxFlatImplementer is an interface for StyleBoxFlat objects.
*/
type StyleBoxFlatImplementer interface {
	Class
}