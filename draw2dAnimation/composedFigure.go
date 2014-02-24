package draw2dAnimation

import (
	"math"
)

// An abstract figure type. Represents a composition of figures kept as collection. Updates to this figure affect it as hole including each part.
type ComposedFigure struct {
	*Figure
	figures *figuresCollection
}

// Constructor setting current struct's fields and default values for the base struct
func NewComposedFigure() *ComposedFigure {
	composedFigure := &ComposedFigure{NewFigure(), newFiguresCollection()}
	composedFigure.SetSubClass(composedFigure)

	return composedFigure
}

// Constructor setting both base struct's and current struct's fields.
func NewComposedFigure3(depth int, startPoint Point, rotationDegrees float64) *ComposedFigure {
	composedFigure := &ComposedFigure{NewFigure(), newFiguresCollection()}
	composedFigure.SetDepth(depth)
	composedFigure.SetStartPoint(startPoint)
	composedFigure.SetRotationDegrees(rotationDegrees)
	composedFigure.SetSubClass(composedFigure)

	return composedFigure
}

// Defines the visualization of the figure according to position (0, 0).
func (this *ComposedFigure) Visualize() {
}

// Adds figure with string key to the contained collection.
func (this *ComposedFigure) AddFigure(name string, figure Figurer) {
	this.figures.add(name, figure)
}

// Removes figure by string key from the contained collection.
func (this *ComposedFigure) RemoveFigure(name string) {
	this.figures.remove(name)
}

// Gets the figure corresponding to the given string key in the contained collection or nil if not found.
func (this *ComposedFigure) GetFigureByName(name string) Figurer {
	return this.figures.getByName(name)
}

// Updates the figure as a whole and each of its parts.
func (this *ComposedFigure) Update() {
	this.GetBase().Update()
	this.figures.traverse(func(figure Figurer) {
		figure.Update()
	})
}

// Draw the figure by visualizing all its part taking in to account the scale, translation and rotation of the figure.
func (this *ComposedFigure) Draw() {
	graphicContext := GetTheImageGraphicContext()
	
	if this.startPoint.X != 0 || this.startPoint.Y != 0 {
		graphicContext.Translate(this.startPoint.X, this.startPoint.Y)
	}
	
	if this.rotationDegrees != 0.0 {
		graphicContext.Rotate(this.rotationDegrees * (math.Pi / 180.0))
	}
	
	if this.scale.X != 1.0 || this.scale.Y != 1.0 {
		graphicContext.Scale(this.scale.X, this.scale.Y)
	}

	this.figures.traverse(func(figure Figurer) {
		figure.Draw()
	})
	
	if this.scale.X != 1.0 || this.scale.Y != 1.0 {
		graphicContext.Scale(1 / this.scale.X, 1 / this.scale.Y)
	}
	
	if this.startPoint.X != 0 || this.startPoint.Y != 0 {
		graphicContext.Translate(-this.startPoint.X, -this.startPoint.Y)
	}
	
	if this.rotationDegrees != 0.0 {
		graphicContext.Rotate(-this.rotationDegrees * (math.Pi / 180.0))
	}
}
