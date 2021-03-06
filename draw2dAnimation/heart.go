package draw2dAnimation

import (
//"math"
)

// A figure type with the shape of a heart. Change width and height for adjusting the figure to the desired ratio.
type Heart struct {
	*Figure
	Width  float64
	Height float64
}

// Constructor setting current struct's fields and default values for the base struct
func NewHeart(width float64, height float64, lineWidth float64) *Heart {
	heart := &Heart{NewFigure(), width, height}
	heart.SetLineWidth(lineWidth)
	heart.SetSubClass(heart)

	return heart
}

// Constructor setting both base struct's and current struct's fields.
func NewHeart6(
	width float64, height float64, depth int, startPoint Point, rotationDegrees float64, lineWidth float64) *Heart {
	heart := &Heart{NewFigure4(depth, startPoint, rotationDegrees, lineWidth), width, height}
	heart.SetSubClass(heart)

	return heart
}

// Defines the visualization of the figure according to position (0, 0).
func (this *Heart) Visualize() {
	graphicContext := GetTheImageGraphicContext()

	// left half
	graphicContext.MoveTo(0, this.Height/2)
	graphicContext.CubicCurveTo(
		-this.Width, 1.0/8*this.Height,
		0, -1.0/4*this.Height,
		0, 1.0/8*this.Height)

	// right half
	graphicContext.MoveTo(0, 1.0/8*this.Height)
	graphicContext.CubicCurveTo(
		0, -1.0/4*this.Height,
		this.Width, 1.0/8*this.Height,
		0, this.Height/2)
}
