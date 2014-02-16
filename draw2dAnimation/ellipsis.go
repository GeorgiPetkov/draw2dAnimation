package draw2dAnimation

import (
	"math"
)

// A figure type representing elipsis. Change RadiusX and RadiusY for adjusting the figure to the desired ratio.
type Ellipsis struct {
	*Figure
	RadiusX   float64
	RadiusY   float64
	LineWidth float64
}

// Constructor setting current struct's fields and default values for the base struct
func NewEllipsis(radiusX float64, radiusY float64, lineWidth float64) *Ellipsis {
	ellipsis := &Ellipsis{NewFigure(), radiusX, radiusY, lineWidth}
	ellipsis.SetSubClass(ellipsis)

	return ellipsis
}

// Constructor setting both base struct's and current struct's fields.
func NewEllipsis5(radiusX float64, radiusY float64, lineWidth float64, depth int, startPoint Point) *Ellipsis {
	ellipsis := &Ellipsis{NewFigure3(depth, startPoint, 0), radiusX, radiusY, lineWidth}
	ellipsis.SetSubClass(ellipsis)

	return ellipsis
}

// Defines the visualization of the figure according to position (0, 0).
func (this *Ellipsis) Visualize() {
	graphicContext := GetTheImageGraphicContext()
	graphicContext.SetLineWidth(this.LineWidth)
	graphicContext.ArcTo(0, 0, this.RadiusX, this.RadiusY, 0, 2*math.Pi)
}
