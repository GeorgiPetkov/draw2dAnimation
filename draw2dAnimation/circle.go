package draw2dAnimation

import (
	"code.google.com/p/draw2d/draw2d"
)

type Circle struct {
	*Figure
	Radius    float64
	LineWidth float64
}

func NewCircle(radius float64, lineWidth float64) *Circle {
	circle := &Circle{NewFigure(), radius, lineWidth}
	circle.SetSubClass(circle)

	return circle
}

func NewCircle4(radius float64, lineWidth float64, depth int, startPoint Point) *Circle {
	circle := &Circle{NewFigure3(depth, startPoint, 0), radius, lineWidth}
	circle.SetSubClass(circle)

	return circle
}

func (this *Circle) Visualize() {
	graphicContext := GetTheImageGraphicContext()
	graphicContext.SetLineWidth(this.LineWidth)
	draw2d.Circle(graphicContext, 0, 0, this.Radius)
}
