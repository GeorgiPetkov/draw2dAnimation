package draw2dAnimation

import (
	"image/color"
	"math"
)

// An adroid figure. Change width and height for adjusting the figure to the desired ratio.
type Android struct {
	*ComposedFigure
	BodyWidth  float64
	BodyHeight float64
}

// Constructor setting current struct's fields and default values for the base struct
func NewAndroid(bodyWidth float64, bodyHeight float64, lineWidth float64, fillColor color.RGBA) *Android {
	android := &Android{NewComposedFigure(), bodyWidth, bodyHeight}
	android.SetLineWidth(lineWidth)
	android.SetSubClass(android)
	android.InitializeFigures(fillColor)

	return android
}

// Constructor setting both base struct's and current struct's fields.
func NewAndroid7(
	bodyWidth float64, bodyHeight float64, fillColor color.RGBA,
	depth int, startPoint Point, rotationDegrees float64, lineWidth float64) *Android {
	android := &Android{NewComposedFigure3(depth, startPoint, rotationDegrees), bodyWidth, bodyHeight}
	android.SetLineWidth(lineWidth)
	android.SetSubClass(android)
	android.InitializeFigures(fillColor)

	return android
}

// Called by constructors to set initial state of the figure. Can also be used for reset.
func (this *Android) InitializeFigures(fillColor color.RGBA) {
	radius := 5.0

	lineWidth := this.GetLineWidth()
	downBodyPart := NewRoundRectangle(radius, this.BodyWidth, 9.0/8*this.BodyHeight, lineWidth)
	this.AddFigure("DownBodyPart", downBodyPart)
	upBodyPart := NewRectangle(this.BodyWidth, this.BodyHeight, lineWidth)
	this.AddFigure("upBodyPart", upBodyPart)

	leftArm := NewRoundRectangle(radius, 0.2*this.BodyWidth, this.BodyHeight-2*lineWidth, lineWidth)
	leftArm.SetStartPoint(Point{-0.2*this.BodyWidth - lineWidth, lineWidth})
	this.AddFigure("leftArm", leftArm)
	rightArm := NewRoundRectangle(radius, 0.2*this.BodyWidth, this.BodyHeight-2*lineWidth, lineWidth)
	rightArm.SetStartPoint(Point{this.BodyWidth + lineWidth, lineWidth})
	this.AddFigure("rightArm", rightArm)

	leftLeg := NewRoundRectangle(radius, 0.2*this.BodyWidth, 5.0/8*this.BodyHeight, lineWidth)
	leftLeg.SetStartPoint(Point{0.2 * this.BodyWidth, this.BodyHeight - lineWidth})
	this.AddFigure("leftLeg", leftLeg)

	rightLeg := NewRoundRectangle(radius, 0.2*this.BodyWidth, 5.0/8*this.BodyHeight, lineWidth)
	rightLeg.SetStartPoint(Point{0.6 * this.BodyWidth, this.BodyHeight - lineWidth})
	this.AddFigure("rightLeg", rightLeg)

	head := NewEllipsis5(
		0.5*this.BodyWidth, 5.0/8*this.BodyHeight, -1, Point{0.5 * this.BodyWidth, -lineWidth}, 0.0, lineWidth)
	this.AddFigure("head", head)

	leftEye := NewCircle4(radius, 0, Point{0.3 * this.BodyWidth, -0.3 * this.BodyWidth}, lineWidth)
	this.AddFigure("leftEye", leftEye)
	rightEye := NewCircle4(radius, 0, Point{0.7 * this.BodyWidth, -0.3 * this.BodyWidth}, lineWidth)
	this.AddFigure("rightEye", rightEye)

	antennaDeltaX := 0.1 * this.BodyWidth
	antennaDeltaY := 1.5 / 8 * this.BodyHeight
	antennaLength := math.Sqrt(antennaDeltaX*antennaDeltaX + antennaDeltaY*antennaDeltaY)
	antennaAngle := -math.Asin(antennaDeltaY/antennaLength) * 180 / math.Pi

	leftAntenna := NewLine5(
		antennaLength, 0, Point{0.3 * this.BodyWidth, -5.0 / 8 * this.BodyHeight}, 180-antennaAngle, lineWidth)
	this.AddFigure("leftAntenna", leftAntenna)
	rightAntenna := NewLine5(
		antennaLength, 0, Point{0.7 * this.BodyWidth, -5.0 / 8 * this.BodyHeight}, antennaAngle, lineWidth)
	this.AddFigure("rightAntenna", rightAntenna)

	this.figures.traverse(func(figure Figurer) {
		figure.SetFillColor(fillColor)
	})
}
