package canvas

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/driver/desktop"
)

func (c *Canvas) Scrolled(ev *fyne.ScrollEvent) {
	c.scale(int(ev.Scrolled.DY))
	c.Refresh()
}
func (c *Canvas) MouseMoved(ev *desktop.MouseEvent) {
	c.TryPan(c.mouseState.previousCoord, ev)
	c.Refresh()
	c.mouseState.previousCoord = &ev.PointEvent
}

func (c *Canvas) MouseIn(env *desktop.MouseEvent) {}
func (c *Canvas) MouseOut()                       {}
