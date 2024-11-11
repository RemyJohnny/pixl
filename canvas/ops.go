package canvas

import "fyne.io/fyne/v2"

func (c *Canvas) scale(direction int) {
	switch {
	case direction > 0:
		c.PxSize += 1
	case direction < 0:
		if c.PxSize > 2 {
			c.PxSize -= 1
		}
	default:
		c.PxSize = 10

	}
}

func (c *Canvas) Pan(previousCoord, currentCoord fyne.PointEvent) {
	xDiff := currentCoord.Position.X - previousCoord.Position.X
	yDiff := currentCoord.Position.Y - previousCoord.Position.Y

	c.CanvasOffset.X += xDiff
	c.CanvasOffset.Y += yDiff
	c.Refresh()
}
