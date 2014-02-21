package draw2dAnimation

// A figure type representing a line.
type Line struct {
	*Figure
	Length float64
}

// Constructor setting current struct's fields and default values for the base struct
func NewLine(length float64, width float64) *Line {
	line := &Line{NewFigure(), length}
	line.SetLineWidth(width)
	line.SetSubClass(line)

	return line
}

// Constructor setting both base struct's and current struct's fields.
func NewLine5(
	length float64, depth int, startPoint Point, rotationDegrees float64, width float64) *Line {
	line := &Line{NewFigure4(depth, startPoint, rotationDegrees, width), length}
	line.SetSubClass(line)

	return line
}

// Defines the visualization of the figure according to position (0, 0).
func (this *Line) Visualize() {
	graphicContext := GetTheImageGraphicContext()
	graphicContext.LineTo(this.Length, 0)
	graphicContext.Stroke()
}
