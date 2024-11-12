package canvas

import (
	"image"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/widget"
	"github.com/RemyJohnny/pixl/apptype"
)

type CanvasMouseState struct {
	previousCoord *fyne.PointEvent
}

type Canvas struct {
	widget.BaseWidget
	apptype.CanvasConfig
	renderer    *CanvasRenderer
	PixelData   image.Image
	mouseState  CanvasMouseState
	appState    *apptype.State
	reloadImage bool
}

func (c *Canvas) Bounds() image.Rectangle {
	x0 := int(c.CanvasOffset.X)
	y0 := int(c.CanvasOffset.Y)
	x1 := int(c.PxCols*c.PxSize + int(c.CanvasOffset.X))
	y1 := int(c.PxRows*c.PxSize + int(c.CanvasOffset.Y))
	return image.Rect(x0, y0, x1, y1)
}

func InBounds(pos fyne.Position, bounds image.Rectangle) bool {
	if pos.X >= float32(bounds.Min.X) && pos.X < float32(bounds.Max.X) &&
		pos.Y >= float32(bounds.Min.Y) && pos.Y < float32(bounds.Max.Y) {
		return true
	}
	return false
}

func NewBlankImage(cols, rows int, c color.Color) image.Image {
	img := image.NewNRGBA(image.Rect(0, 0, cols, rows))
	for y := 0; y < rows; y++ {
		for x := 0; x < cols; x++ {
			img.Set(x, y, c)
		}
	}
	return img
}

func NewCanvas(state *apptype.State, config apptype.CanvasConfig) *Canvas {
	canvas := &Canvas{
		CanvasConfig: config,
		appState:     state,
	}
	canvas.PixelData = NewBlankImage(canvas.PxCols, canvas.PxRows, color.NRGBA{128, 128, 128, 128})
	canvas.ExtendBaseWidget(canvas)
	return canvas
}

func (c *Canvas) CreateRenderer() fyne.WidgetRenderer {
	canvasImage := canvas.NewImageFromImage(c.PixelData)
	canvasImage.ScaleMode = canvas.ImageScalePixels
	canvasImage.FillMode = canvas.ImageFillContain

	canvasBorder := make([]canvas.Line, 4)
	for i := 0; i < len(canvasBorder); i++ {
		canvasBorder[i].StrokeColor = color.NRGBA{100, 100, 100, 255}
		canvasBorder[i].StrokeWidth = 2
	}
	renderer := &CanvasRenderer{
		Canvas:       c,
		canvasImage:  canvasImage,
		canvasBorder: canvasBorder,
	}

	c.renderer = renderer
	return renderer
}

func (c *Canvas) TryPan(previousCoord *fyne.PointEvent, ev *desktop.MouseEvent) {
	if previousCoord != nil && ev.Button == desktop.MouseButtonTertiary {
		c.Pan((*previousCoord), ev.PointEvent)
	}
}

//Brushable Interface

func (c *Canvas) SetColor(color color.Color, x, y int) {
	if nrgba, ok := c.PixelData.(*image.NRGBA); ok {
		nrgba.Set(x, y, color)
	}
	if rgba, ok := c.PixelData.(*image.RGBA); ok {
		rgba.Set(x, y, color)
	}
	c.Refresh()
}

func (c *Canvas) MouseToCanvasXY(ev *desktop.MouseEvent) (*int, *int) {
	bounds := c.Bounds()
	if !InBounds(ev.Position, bounds) {
		return nil, nil
	}
	pxSize := float32(c.PxSize)
	xOffset := c.CanvasOffset.X
	yOffset := c.CanvasOffset.Y

	x := int((ev.Position.X - xOffset) / pxSize)
	y := int((ev.Position.Y - yOffset) / pxSize)

	return &x, &y
}
