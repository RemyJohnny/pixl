package main

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"github.com/RemyJohnny/pixl/apptype"
	"github.com/RemyJohnny/pixl/swatch"
	"github.com/RemyJohnny/pixl/ui"
)

func main() {
	pixlApp := app.New()
	pixlWindow := pixlApp.NewWindow("pixl")

	state := apptype.State{
		BrushColor:     color.NRGBA{255, 255, 255, 255},
		SwatchSelected: 0,
	}

	appInit := ui.AppInit{
		PixlWindow: pixlWindow,
		State:      &state,
		Swatches:   make([]*swatch.Swatch, 0, 64),
	}

	ui.Setup(&appInit)

	appInit.PixlWindow.ShowAndRun()
}
