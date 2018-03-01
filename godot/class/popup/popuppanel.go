package popup

/*------------------------------------------------------------------------------
//   This code was generated by a tool.
//
//   Changes to this file may cause incorrect behavior and will be lost if
//   the code is regenerated. Any updates should be done in
//   "class.go.tmpl" so they can be included in the generated
//   code.
//----------------------------------------------------------------------------*/

/*
Class for displaying popups with a panel background. In some cases it might be simpler to use than [Popup], since it provides a configurable background. If you are making windows, better check [WindowDialog].
*/
type PopupPanel struct {
	Popup
}

func (o *PopupPanel) BaseClass() string {
	return "PopupPanel"
}

/*
   PopupPanelImplementer is an interface for PopupPanel objects.
*/
type PopupPanelImplementer interface {
	Class
}