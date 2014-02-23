package draw2dAnimation

// A figure type representing a rectangle. Change width and height for adjusting the figure to the desired ratio.
type Rectangle struct {
	*Figure
	Width  float64
	Height float64
}

// Constructor setting current struct's fields and default values for the base struct
func NewRectangle(width float64, height float64, lineWidth float64) *Rectangle {
	rectangle := &Rectangle{NewFigure(), width, height}
	rectangle.SetLineWidth(lineWidth)
	rectangle.SetSubClass(rectangle)

	return rectangle
}

// Constructor setting both base struct's and current struct's fields.
func NewRectangle6(
	width float64, height float64, depth int, startPoint Point, rotationDegrees float64, lineWidth float64) *Rectangle {
	rectangle := &Rectangle{NewFigure4(depth, startPoint, rotationDegrees, lineWidth), width, height}
	rectangle.SetSubClass(rectangle)

	return rectangle
}

// Gets the width of the rectangle.
func (this *Rectangle) GetWidth() float64 {
	return this.Width
}

// Sets the width of the rectangle.
func (this *Rectangle) SetWidth(value float64) {
	this.Width = value
}

// Gets the height of the rectangle.
func (this *Rectangle) GetHeight() float64 {
	return this.Height
}

// Sets the height of the rectangle.
func (this *Rectangle) SetHeight(value float64) {
	this.Height = value
}

// Defines the visualization of the figure according to position (0, 0).
func (this *Rectangle) Visualize() {
	graphicContext := GetTheImageGraphicContext()
	graphicContext.MoveTo(0, 0)
	graphicContext.LineTo(this.Width, 0)
	graphicContext.LineTo(this.Width, this.Height)
	graphicContext.LineTo(0, this.Height)
	graphicContext.LineTo(0, 0)
}
