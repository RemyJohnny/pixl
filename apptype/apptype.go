package apptype

import (
	"image/color"

	"fyne.io/fyne/v2"
)

type BrushType = int

type CanvasConfig struct {
	drawingArea    fyne.Size
	CanvasOffset   fyne.Position
	pxRows, PxCols int
	PxSize         int
}

type State struct {
	BrushColor     color.Color
	BrushType      int
	SwatchSelected int
	FilePath       string
}

func (s *State) SetFilePath(path string) {
	s.FilePath = path
}
