package draw2dAnimation

import (
	"image/color"
	"math"
)

type Figure struct {
	subClass              Figurer
	id                    int
	depth                 int
	startPoint            Point
	rotationDegrees       float64
	fillColor             color.RGBA
	strokeColor           color.RGBA
	isFilled              bool
	updateType            UpdateType
	updateTranslation     Point
	updateRotationDegrees float64
	updateMethod          func(Figurer)
}

func NewFigure() *Figure {
	return NewFigure3(0, Point{0, 0}, 0.0)
}

func NewFigure3(depth int, startPoint Point, rotationDegrees float64) *Figure {
	nextFigureId++
	return &Figure{
		id:              nextFigureId,
		depth:           depth,
		startPoint:      startPoint,
		rotationDegrees: rotationDegrees,
		fillColor:       color.RGBA{255, 255, 255, 255},
		strokeColor:     color.RGBA{0, 0, 0, 255},
		isFilled:        false}
}

func (this *Figure) GetBase() *Figure {
	return this
}

func (this *Figure) SetSubClass(value Figurer) {
	this.subClass = value
}

func (this *Figure) getId() int {
	return this.id
}

func (this *Figure) GetDepth() int {
	return this.depth
}

func (this *Figure) SetDepth(value int) {
	this.depth = value
}

func (this *Figure) GetStartPoint() Point {
	return Point{this.startPoint.X, this.startPoint.Y}
}

func (this *Figure) SetStartPoint(value Point) {
	this.startPoint = value
}

func (this *Figure) GetRotationDegrees() float64 {
	return this.rotationDegrees
}

func (this *Figure) SetRotationDegrees(value float64) {
	this.rotationDegrees = value
}

func (this *Figure) GetFillColor() color.RGBA {
	return this.fillColor
}

func (this *Figure) SetFillColor(value color.RGBA) {
	this.fillColor = value
	this.isFilled = true
}

func (this *Figure) GetStrokeColor() color.RGBA {
	return this.strokeColor
}

func (this *Figure) SetStrokeColor(value color.RGBA) {
	this.strokeColor = value
	this.isFilled = true
}

func (this *Figure) GetIsFilled() bool {
	return this.isFilled
}

func (this *Figure) SetIsFilled(value bool) {
	this.isFilled = value
}

func (this *Figure) GetUpdateRotationDegrees() float64 {
	return this.updateRotationDegrees
}

func (this *Figure) SetUpdateRotationDegrees(value float64) {
	if value == 0.0 {
		this.updateType &^= Rotation
	} else {
		this.updateType |= Rotation
	}

	this.updateRotationDegrees = value
}

func (this *Figure) GetUpdateTranslation() Point {
	return Point{
		this.updateTranslation.X,
		this.updateTranslation.Y}
}

func (this *Figure) SetUpdateTranslation(value Point) {
	if value.X == 0 && value.Y == 0 {
		this.updateType &^= Translation
	} else {
		this.updateType |= Translation
	}

	this.updateTranslation = value
}

func (this *Figure) GetUpdateMethod() func(Figurer) {
	return this.updateMethod
}

func (this *Figure) SetUpdateMethod(value func(Figurer)) {
	if value == nil {
		this.updateType &^= Custom
	} else {
		this.updateType |= Custom
	}

	this.updateMethod = value
}

func (this Figure) Draw() {
	graphicContext := GetTheImageGraphicContext()
	graphicContext.Save()
	graphicContext.Translate(this.startPoint.X, this.startPoint.Y)
	graphicContext.Rotate(this.rotationDegrees * (math.Pi / 180.0))

	if this.isFilled {
		graphicContext.SetFillColor(this.fillColor)
	}

	graphicContext.SetStrokeColor(this.strokeColor)

	this.subClass.Visualize()

	if this.isFilled {
		graphicContext.FillStroke()
	} else {
		graphicContext.Stroke()
	}

	graphicContext.Restore()
}

func (this *Figure) Update() {
	if (this.updateType & Custom) != 0 {
		this.updateMethod(this)
	}

	if (this.updateType & Translation) != 0 {
		this.startPoint.X += this.updateTranslation.X
		this.startPoint.Y += this.updateTranslation.Y
	}

	if (this.updateType & Rotation) != 0 {
		this.rotationDegrees += this.updateRotationDegrees
	}
}

func (this *Figure) Visualize() {
}
