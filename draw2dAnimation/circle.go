package draw2dAnimation

import (
	"code.google.com/p/draw2d/draw2d"
)

// A figure type representing circle.
type Circle struct {
	*Figure
	Radius float64
}

// Constructor setting current struct's fields and default values for the base struct
func NewCircle(radius float64, lineWidth float64) *Circle {
	circle := &Circle{NewFigure(), radius}
	circle.SetLineWidth(lineWidth)
	circle.SetSubClass(circle)

	return circle
}

// Constructor setting both base struct's and current struct's fields.
func NewCircle4(radius float64, depth int, startPoint Point, lineWidth float64) *Circle {
	circle := &Circle{NewFigure4(depth, startPoint, 0.0, lineWidth), radius}
	circle.SetSubClass(circle)

	return circle
}

// Defines the visualization of the figure according to position (0, 0).
func (this *Circle) Visualize() {
	graphicContext := GetTheImageGraphicContext()
	draw2d.Circle(graphicContext, 0, 0, this.Radius)
}
