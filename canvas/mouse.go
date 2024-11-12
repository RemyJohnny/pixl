package canvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/RemyJohnny/pixl/canvas/brush"
)

func (c *Canvas) Scrolled(ev *fyne.ScrollEvent) {
	c.scale(int(ev.Scrolled.DY))
	c.Refresh()
}
func (c *Canvas) MouseMoved(ev *desktop.MouseEvent) {
	if x, y := c.MouseToCanvasXY(ev); x != nil && y != nil {
		brush.TryBrush(c.appState, c, ev)
		cursor := brush.Cursor(c.CanvasConfig, c.appState.BrushType, ev, *x, *y)
		c.renderer.SetCursor(cursor)
	} else {
		c.renderer.SetCursor(make([]fyne.CanvasObject, 0))
	}
	c.TryPan(c.mouseState.previousCoord, ev)
	c.Refresh()
	c.mouseState.previousCoord = &ev.PointEvent
}

func (c *Canvas) MouseIn(env *desktop.MouseEvent) {}
func (c *Canvas) MouseOut()                       {}

func (c *Canvas) MouseDown(ev *desktop.MouseEvent) {
	brush.TryBrush(c.appState, c, ev)
}
func (c *Canvas) MouseUp(ev *desktop.MouseEvent) {}
