package draw2dAnimation

import (
	"image/color"
	"math"
)

type Android struct {
	*ComposedFigure
	BodyWidth  float64
	BodyHeight float64
	LineWidth  float64
}

func NewAndroid(bodyWidth float64, bodyHeight float64, lineWidth float64, fillColor color.RGBA) *Android {
	android := &Android{NewComposedFigure(), bodyWidth, bodyHeight, lineWidth}
	android.SetSubClass(android)
	android.InitializeFigures(fillColor)

	return android
}

func NewAndroid5(
	bodyWidth float64, bodyHeight float64, lineWidth float64,
	depth int, startPoint Point, rotationDegrees float64, fillColor color.RGBA) *Android {
	android := &Android{NewComposedFigure3(depth, startPoint, rotationDegrees), bodyWidth, bodyHeight, lineWidth}
	android.SetSubClass(android)
	android.InitializeFigures(fillColor)

	return android
}

func (this *Android) InitializeFigures(fillColor color.RGBA) {
	radius := 5.0

	downBodyPart := NewRoundRectangle(this.BodyWidth, 9.0/8*this.BodyHeight, this.LineWidth, radius)
	this.AddFigure("DownBodyPart", downBodyPart)
	upBodyPart := NewRectangle(this.BodyWidth, this.BodyHeight, this.LineWidth)
	this.AddFigure("upBodyPart", upBodyPart)

	leftArm := NewRoundRectangle(0.2*this.BodyWidth, this.BodyHeight-2*this.LineWidth, this.LineWidth, radius)
	leftArm.SetStartPoint(Point{-0.2*this.BodyWidth - this.LineWidth, this.LineWidth})
	this.AddFigure("leftArm", leftArm)
	rightArm := NewRoundRectangle(0.2*this.BodyWidth, this.BodyHeight-2*this.LineWidth, this.LineWidth, radius)
	rightArm.SetStartPoint(Point{this.BodyWidth + this.LineWidth, this.LineWidth})
	this.AddFigure("rightArm", rightArm)

	leftLeg := NewRoundRectangle(0.2*this.BodyWidth, 5.0/8*this.BodyHeight, this.LineWidth, radius)
	leftLeg.SetStartPoint(Point{0.2 * this.BodyWidth, this.BodyHeight - this.LineWidth})
	this.AddFigure("leftLeg", leftLeg)

	rightLeg := NewRoundRectangle(0.2*this.BodyWidth, 5.0/8*this.BodyHeight, this.LineWidth, radius)
	rightLeg.SetStartPoint(Point{0.6 * this.BodyWidth, this.BodyHeight - this.LineWidth})
	this.AddFigure("rightLeg", rightLeg)

	head := NewEllipsis5(
		0.5*this.BodyWidth, 5.0/8*this.BodyHeight, this.LineWidth, -1, Point{0.5 * this.BodyWidth, -this.LineWidth})
	this.AddFigure("head", head)

	leftEye := NewCircle4(radius, this.LineWidth, 0, Point{0.3 * this.BodyWidth, -0.3 * this.BodyWidth})
	this.AddFigure("leftEye", leftEye)
	rightEye := NewCircle4(radius, this.LineWidth, 0, Point{0.7 * this.BodyWidth, -0.3 * this.BodyWidth})
	this.AddFigure("rightEye", rightEye)

	antennaDeltaX := 0.1 * this.BodyWidth
	antennaDeltaY := 1.5 / 8 * this.BodyHeight
	antennaLength := math.Sqrt(antennaDeltaX*antennaDeltaX + antennaDeltaY*antennaDeltaY)
	antennaAngle := -math.Asin(antennaDeltaY/antennaLength) * 180 / math.Pi

	leftAntenna := NewLine5(
		antennaLength, this.LineWidth, 0, Point{0.3 * this.BodyWidth, -5.0 / 8 * this.BodyHeight}, 180-antennaAngle)
	this.AddFigure("leftAntenna", leftAntenna)
	rightAntenna := NewLine5(
		antennaLength, this.LineWidth, 0, Point{0.7 * this.BodyWidth, -5.0 / 8 * this.BodyHeight}, antennaAngle)
	this.AddFigure("rightAntenna", rightAntenna)

	this.figures.traverse(func(figure Figurer) {
		figure.SetFillColor(fillColor)
	})
}
