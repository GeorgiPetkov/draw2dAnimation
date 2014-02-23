package draw2dAnimation

import (
	"math"
)

// A figure type representing ellipsis. Change RadiusX and RadiusY for adjusting the figure to the desired ratio.
type Ellipsis struct {
	*Figure
	RadiusX float64
	RadiusY float64
}

// Constructor setting current struct's fields and default values for the base struct
func NewEllipsis(radiusX float64, radiusY float64, lineWidth float64) *Ellipsis {
	ellipsis := &Ellipsis{NewFigure(), radiusX, radiusY}
	ellipsis.SetLineWidth(lineWidth)
	ellipsis.SetSubClass(ellipsis)

	return ellipsis
}

// Constructor setting both base struct's and current struct's fields.
func NewEllipsis5(
	radiusX float64, radiusY float64, depth int, startPoint Point, rotationDegrees float64, lineWidth float64) *Ellipsis {
	ellipsis := &Ellipsis{NewFigure4(depth, startPoint, rotationDegrees, lineWidth), radiusX, radiusY}
	ellipsis.SetSubClass(ellipsis)

	return ellipsis
}

// Defines the visualization of the figure according to position (0, 0).
func (this *Ellipsis) Visualize() {
	graphicContext := GetTheImageGraphicContext()
	graphicContext.ArcTo(0, 0, this.RadiusX, this.RadiusY, 0, 2*math.Pi)
}
