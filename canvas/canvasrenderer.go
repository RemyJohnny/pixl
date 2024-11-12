package canvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type CanvasRenderer struct {
	Canvas       *Canvas
	canvasImage  *canvas.Image
	canvasBorder []canvas.Line
	canvasCursor []fyne.CanvasObject
}

func (r *CanvasRenderer) SetCursor(objects []fyne.CanvasObject) {
	r.canvasCursor = objects
}

// widgetRenderer interface implementation
func (r CanvasRenderer) MinSize() fyne.Size {
	return r.Canvas.DrawingArea
}

func (r CanvasRenderer) Objects() []fyne.CanvasObject {
	objects := make([]fyne.CanvasObject, 0, 5)
	for i := 0; i < len(r.canvasBorder); i++ {
		objects = append(objects, &r.canvasBorder[i])
	}
	objects = append(objects, r.canvasImage)
	objects = append(objects, r.canvasCursor...)
	return objects
}

func (r *CanvasRenderer) Destroy() {}

func (r *CanvasRenderer) Layout(size fyne.Size) {
	r.LayoutCanvas(size)
	r.LayoutBorder(size)
}

func (r CanvasRenderer) Refresh() {
	if r.Canvas.reloadImage {
		r.canvasImage = canvas.NewImageFromImage(r.Canvas.PixelData)
		r.canvasImage.ScaleMode = canvas.ImageScalePixels
		r.canvasImage.FillMode = canvas.ImageFillContain
		r.Canvas.reloadImage = false
	}
	r.Layout(r.Canvas.Size())
	canvas.Refresh(r.canvasImage)
}

func (r *CanvasRenderer) LayoutCanvas(size fyne.Size) {
	imgPxWidth := r.Canvas.PxCols
	imgPxHeight := r.Canvas.PxRows
	pxSize := r.Canvas.PxSize
	r.canvasImage.Move(fyne.NewPos(r.Canvas.CanvasOffset.X, r.Canvas.CanvasOffset.Y))
	r.canvasImage.Resize(fyne.Size{Width: float32(imgPxWidth * pxSize), Height: float32(imgPxHeight * pxSize)})

}

func (r *CanvasRenderer) LayoutBorder(size fyne.Size) {
	offset := r.Canvas.CanvasOffset
	imgHeight := r.canvasImage.Size().Height
	imgWidth := r.canvasImage.Size().Width

	left := &r.canvasBorder[0]
	left.Position1 = fyne.NewPos(offset.X, offset.Y)
	left.Position2 = fyne.NewPos(offset.X, offset.Y+imgHeight)

	top := &r.canvasBorder[1]
	top.Position1 = fyne.NewPos(offset.X, offset.Y)
	top.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y)

	right := &r.canvasBorder[2]
	right.Position1 = fyne.NewPos(offset.X+imgWidth, offset.Y)
	right.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y+imgHeight)

	bottom := &r.canvasBorder[3]
	bottom.Position1 = fyne.NewPos(offset.X, offset.Y+imgHeight)
	bottom.Position2 = fyne.NewPos(offset.X+imgWidth, offset.Y+imgHeight)
}
