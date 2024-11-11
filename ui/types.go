package ui

import (
	"fyne.io/fyne/v2"
	"github.com/RemyJohnny/pixl/apptype"
	"github.com/RemyJohnny/pixl/swatch"
)

type AppInit struct {
	PixlWindow fyne.Window
	State      *apptype.State
	Swatches   []*swatch.Swatch
}
