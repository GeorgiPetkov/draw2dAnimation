package draw2dAnimation

// A figure type representing a line.
type Line struct {
	*Figure
	Length float64
	Width  float64
}

// Constructor setting current struct's fields and default values for the base struct
func NewLine(length float64, width float64) *Line {
	line := &Line{NewFigure(), length, width}
	line.SetSubClass(line)

	return line
}

// Constructor setting both base struct's and current struct's fields.
func NewLine5(
	length float64, width float64, depth int, startPoint Point, rotationDegrees float64) *Line {
	line := &Line{NewFigure3(depth, startPoint, rotationDegrees), length, width}
	line.SetSubClass(line)

	return line
}

// Defines the visualization of the figure according to position (0, 0).
func (this *Line) Visualize() {
	graphicContext := GetTheImageGraphicContext()
	graphicContext.SetLineWidth(this.Width)
	graphicContext.LineTo(this.Length, 0)
	graphicContext.Stroke()
}
