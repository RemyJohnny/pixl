package ui

import (
	"fyne.io/fyne/v2"
	"github.com/RemyJohnny/pixl/apptype"
	"github.com/RemyJohnny/pixl/canvas"
	"github.com/RemyJohnny/pixl/swatch"
)

type AppInit struct {
	PixlCanvas *canvas.Canvas
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
