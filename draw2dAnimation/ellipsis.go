package draw2dAnimation

import (
	"math"
)

type Ellipsis struct {
	*Figure
	RadiusX   float64
	RadiusY   float64
	LineWidth float64
}

func NewEllipsis(radiusX float64, radiusY float64, lineWidth float64) *Ellipsis {
	ellipsis := &Ellipsis{NewFigure(), radiusX, radiusY, lineWidth}
	ellipsis.SetSubClass(ellipsis)

	return ellipsis
}

func NewEllipsis5(radiusX float64, radiusY float64, lineWidth float64, depth int, startPoint Point) *Ellipsis {
	ellipsis := &Ellipsis{NewFigure3(depth, startPoint, 0), radiusX, radiusY, lineWidth}
	ellipsis.SetSubClass(ellipsis)

	return ellipsis
}

func (this *Ellipsis) Visualize() {
	graphicContext := GetTheImageGraphicContext()
	graphicContext.SetLineWidth(this.LineWidth)
	graphicContext.ArcTo(0, 0, this.RadiusX, this.RadiusY, 0, 2*math.Pi)
}
