package draw2dAnimation

import (
	"image/color"
)

// Represents the base interface of all figures. For information on the methods see struct Figure.
type Figurer interface {
	GetDepth() int
	SetDepth(value int)
	GetStartPoint() Point
	SetStartPoint(value Point)
	GetRotationDegrees() float64
	SetRotationDegrees(value float64)
	GetFillColor() color.RGBA
	SetFillColor(value color.RGBA)
	GetStrokeColor() color.RGBA
	SetStrokeColor(value color.RGBA)
	GetIsFilled() bool
	SetIsFilled(value bool)
	GetLineWidth() float64
	SetLineWidth(value float64)
	GetScale() Point
	SetScale(value Point)
	GetUpdateTranslation() Point
	SetUpdateTranslation(value Point)
	GetUpdateRotationDegrees() float64
	SetUpdateRotationDegrees(value float64)
	GetUpdateMethod() func(Figurer)
	SetUpdateMethod(value func(Figurer))
	Draw()
	Update()
	Visualize()
	getId() int
}
