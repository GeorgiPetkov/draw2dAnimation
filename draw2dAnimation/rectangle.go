package draw2dAnimation

// A figure type representing a rectangle. Change width and height for adjusting the figure to the desired ratio.
type Rectangle struct {
	*Figure
	Length float64
	Width  float64
}

// Constructor setting current struct's fields and default values for the base struct
func NewRectangle(length float64, width float64, lineWidth float64) *Rectangle {
	rectangle := &Rectangle{NewFigure(), length, width}
	rectangle.SetLineWidth(lineWidth)
	rectangle.SetSubClass(rectangle)

	return rectangle
}

// Constructor setting both base struct's and current struct's fields.
func NewRectangle6(
	length float64, width float64, depth int, startPoint Point, rotationDegrees float64, lineWidth float64) *Rectangle {
	rectangle := &Rectangle{NewFigure4(depth, startPoint, rotationDegrees, lineWidth), length, width}
	rectangle.SetSubClass(rectangle)

	return rectangle
}

// Defines the visualization of the figure according to position (0, 0).
func (this *Rectangle) Visualize() {
	graphicContext := GetTheImageGraphicContext()
	graphicContext.MoveTo(0, 0)
	graphicContext.LineTo(this.Length, 0)
	graphicContext.LineTo(this.Length, this.Width)
	graphicContext.LineTo(0, this.Width)
}
