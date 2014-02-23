package draw2dAnimation

import (
	"code.google.com/p/draw2d/draw2d"
)

// A figure type with the shape a rectangle with rounded edges. Change width, height and radius for adjusting the figure to the desired ratio.
type RoundRectangle struct {
	*Rectangle
	Radius float64
}

// Constructor setting current struct's fields and default values for the base struct
func NewRoundRectangle(radius float64, width float64, height float64, lineWidth float64) *RoundRectangle {
	roundRectangle := &RoundRectangle{NewRectangle(width, height, lineWidth), radius}
	roundRectangle.SetSubClass(roundRectangle)

	return roundRectangle
}

// Constructor setting both base struct's and current struct's fields.
func NewRoundRectangle7(
	radius float64, width float64, height float64,
	depth int, startPoint Point, rotationDegrees float64, lineWidth float64) *RoundRectangle {
	roundRectangle := &RoundRectangle{NewRectangle6(width, height, depth, startPoint, rotationDegrees, lineWidth), radius}
	roundRectangle.SetSubClass(roundRectangle)

	return roundRectangle
}

// Defines the visualization of the figure according to position (0, 0).
func (this *RoundRectangle) Visualize() {
	graphicContext := GetTheImageGraphicContext()
	draw2d.RoundRect(graphicContext, 0, 0, this.Width, this.Height, this.Radius, this.Radius)
}
