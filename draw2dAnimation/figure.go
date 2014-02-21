package draw2dAnimation

import (
	"fmt"
	"image/color"
	"math"
)

// An abstract figure type. Represents a base struct for all figures.
type Figure struct {
	subClass              Figurer
	id                    int
	depth                 int
	startPoint            Point
	rotationDegrees       float64
	fillColor             color.RGBA
	strokeColor           color.RGBA
	isFilled              bool
	updateTypes           updateType
	updateTranslation     Point
	updateRotationDegrees float64
	updateMethod          func(Figurer)
}

// Default constructor.
func NewFigure() *Figure {
	return NewFigure3(0, Point{0, 0}, 0.0)
}

// Constructor accepting depth(layer in the image), startPoint(to which all actions are related) and rotation degrees.
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

// Gets the current instance. Has meaning to be used from struct extending this one if need.
func (this *Figure) GetBase() *Figure {
	return this
}

// Sets the final figure in the extending chain. Should be called in the constructor of each extending struct. Has meaning to be used from struct extending this one to use function as virtual.
func (this *Figure) SetSubClass(value Figurer) {
	this.subClass = value
}

// Gets the unique ID for the figure. Used to maintain order of figures when their depth is equal.
func (this *Figure) getId() int {
	return this.id
}

// Gets the depth(layer) of the figure in the image.
func (this *Figure) GetDepth() int {
	return this.depth
}

// Sets the depth(layer) of the figure in the image.
func (this *Figure) SetDepth(value int) {
	this.depth = value
}

// Gets the start point of the figure.
func (this *Figure) GetStartPoint() Point {
	return Point{this.startPoint.X, this.startPoint.Y}
}

//Sets the start point of the figure.
func (this *Figure) SetStartPoint(value Point) {
	this.startPoint = value
}

// Get the current degrees by which the figure is rotated.
func (this *Figure) GetRotationDegrees() float64 {
	return this.rotationDegrees
}

// Sets the rotation of the figure for the time after the next call of Update()
func (this *Figure) SetRotationDegrees(value float64) {
	this.rotationDegrees = value
}

// Gets the color used to fill the figure.
func (this *Figure) GetFillColor() color.RGBA {
	return this.fillColor
}

// Sets the color to be used to fill the figure. Automatically set the figure as one to be filled.
func (this *Figure) SetFillColor(value color.RGBA) {
	this.fillColor = value
	this.isFilled = true
}

// Gets to color used for drawing the contour of the figure.
func (this *Figure) GetStrokeColor() color.RGBA {
	return this.strokeColor
}

// Sets the color to be used to draw the contour of the figure.
func (this *Figure) SetStrokeColor(value color.RGBA) {
	this.strokeColor = value
}

// Gets whether the figure should be filled or stroked.
func (this *Figure) GetIsFilled() bool {
	return this.isFilled
}

// Sets whether the figure should be filled or stroked.
func (this *Figure) SetIsFilled(value bool) {
	this.isFilled = value
}

// Gets the degrees by which the figure rotates on each call of Update().
func (this *Figure) GetUpdateRotationDegrees() float64 {
	return this.updateRotationDegrees
}

// Sets the degrees by which the figure should rotate on each call of Update().
func (this *Figure) SetUpdateRotationDegrees(value float64) {
	if value == 0.0 {
		this.updateTypes &^= Rotation
	} else {
		this.updateTypes |= Rotation
	}

	this.updateRotationDegrees = value
}

// Gets the vector by which the figure is translated on each call of Update().
func (this *Figure) GetUpdateTranslation() Point {
	return this.updateTranslation
}

// Sets the vector by which the figure should translate on each call of Update().
func (this *Figure) SetUpdateTranslation(value Point) {
	if value.X == 0 && value.Y == 0 {
		this.updateTypes &^= Translation
	} else {
		this.updateTypes |= Translation
	}

	this.updateTranslation = value
}

// Gets the custom update method used to update the figure on each call of Update().
func (this *Figure) GetUpdateMethod() func(Figurer) {
	return this.updateMethod
}

// Sets a custom update method to be used in updating the figure on each call of Update().
func (this *Figure) SetUpdateMethod(value func(Figurer)) {
	if value == nil {
		this.updateTypes &^= Custom
	} else {
		this.updateTypes |= Custom
	}

	this.updateMethod = value
}

// Draws the figure taking into account the translation and rotation of the figure and using the implemented by the extending substruct Visualize() method.
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

// Updates the figure by the custom method, the update translation and the update rotation degrees.
func (this *Figure) Update() {
	if (this.updateTypes & Custom) != 0 {
		this.updateMethod(this)
	}

	if (this.updateTypes & Translation) != 0 {
		this.startPoint.X += this.updateTranslation.X
		this.startPoint.Y += this.updateTranslation.Y
	}

	if (this.updateTypes & Rotation) != 0 {
		this.rotationDegrees += this.updateRotationDegrees
	}
}

// Does nothing. Needed to implement the Figurer interface. Should be ovewritted by extending substruct.
func (this *Figure) Visualize() {
	panic(fmt.Sprintf("Type %T doesn't implement Visualize() method.", this.subClass))
}
