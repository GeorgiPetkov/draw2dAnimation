package draw2dAnimation

import (
//"math"
)

type Heart struct {
	*Figure
	Width     float64
	Height    float64
	LineWidth float64
}

func NewHeart(width float64, height float64, lineWidth float64) *Heart {
	heart := &Heart{NewFigure(), width, height, lineWidth}
	heart.SetSubClass(heart)

	return heart
}

func NewHeart5(width float64, height float64, lineWidth float64, depth int, startPoint Point) *Heart {
	heart := &Heart{NewFigure3(depth, startPoint, 0), width, height, lineWidth}
	heart.SetSubClass(heart)

	return heart
}

func (this *Heart) Visualize() {
	graphicContext := GetTheImageGraphicContext()
	graphicContext.SetLineWidth(this.LineWidth)

	// left half
	graphicContext.MoveTo(0, this.Height/2)
	graphicContext.CubicCurveTo(
		-this.Width, 1.0/8*this.Height,
		0, -1.0/4*this.Height,
		0, 1.0/8*this.Height)

	// right half
	graphicContext.MoveTo(0, this.Height/2)
	graphicContext.CubicCurveTo(
		this.Width, 1.0/8*this.Height,
		0, -1.0/4*this.Height,
		0, 1.0/8*this.Height)
	graphicContext.Stroke()
}
