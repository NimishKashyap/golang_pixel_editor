package ui

import (
	"fyne.io/fyne/v2"
	"github.com/NimishKashyap/nixpix/apptype"
	"github.com/NimishKashyap/nixpix/pxcanvas"
	"github.com/NimishKashyap/nixpix/swatch"
)

type AppInit struct {
	PixlCanvas *pxcanvas.PxCanvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
