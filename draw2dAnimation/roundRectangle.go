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
func NewRoundRectangle(length float64, width float64, lineWidth float64, radius float64) *RoundRectangle {
	rectangle := &RoundRectangle{NewRectangle(length, width, lineWidth), radius}
	rectangle.SetSubClass(rectangle)

	return rectangle
}

// Constructor setting both base struct's and current struct's fields.
func NewRoundRectangle7(
	length float64, width float64, lineWidth float64, radius float64,
	depth int, startPoint Point, rotationDegrees float64) *RoundRectangle {
	rectangle := &RoundRectangle{NewRectangle6(length, width, lineWidth, depth, startPoint, rotationDegrees), radius}
	rectangle.SetSubClass(rectangle)

	return rectangle
}

// Defines the visualization of the figure according to position (0, 0).
func (this *RoundRectangle) Visualize() {
	graphicContext := GetTheImageGraphicContext()
	graphicContext.SetLineWidth(this.LineWidth)
	draw2d.RoundRect(graphicContext, 0, 0, this.Length, this.Width, this.Radius, this.Radius)

	if this.isFilled {
		graphicContext.FillStroke()
	} else {
		graphicContext.Stroke()
	}
}
