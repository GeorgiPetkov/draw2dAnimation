package draw2dAnimation

import (
	"math"
)

type ComposedFigure struct {
	*Figure
	figures *figuresCollection
}

func NewComposedFigure() *ComposedFigure {
	composedFigure := &ComposedFigure{NewFigure(), newFiguresCollection()}
	composedFigure.SetSubClass(composedFigure)

	return composedFigure
}

func NewComposedFigure3(depth int, startPoint Point, rotationDegrees float64) *ComposedFigure {
	composedFigure := &ComposedFigure{NewFigure3(depth, startPoint, rotationDegrees), newFiguresCollection()}
	composedFigure.SetSubClass(composedFigure)

	return composedFigure
}

func (this *ComposedFigure) Visualize() {
}

func (this *ComposedFigure) AddFigure(name string, figure Figurer) {
	this.figures.add(name, figure)
}

func (this *ComposedFigure) RemoveFigure(name string) {
	this.figures.remove(name)
}

func (this *ComposedFigure) GetFigureByName(name string) Figurer {
	return this.figures.getByName(name)
}

func (this *ComposedFigure) Update() {
	this.GetBase().Update()
	this.figures.traverse(func(figure Figurer) {
		figure.Update()
	})
}

func (this *ComposedFigure) Draw() {
	this.figures.traverse(func(figure Figurer) {
		graphicContext := GetTheImageGraphicContext()
		graphicContext.Translate(this.startPoint.X, this.startPoint.Y)
		graphicContext.Rotate(this.rotationDegrees * (math.Pi / 180.0))

		figure.Draw()

		graphicContext.Rotate(-this.rotationDegrees * (math.Pi / 180.0))
		graphicContext.Translate(-this.startPoint.X, -this.startPoint.Y)
	})
}