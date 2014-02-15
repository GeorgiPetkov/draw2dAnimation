package draw2dAnimation

type Line struct {
	*Figure
	Length float64
	Width  float64
}

func NewLine(length float64, width float64) *Line {
	line := &Line{NewFigure(), length, width}
	line.SetSubClass(line)

	return line
}

func NewLine5(
	length float64, width float64, depth int, startPoint Point, rotationDegrees float64) *Line {
	line := &Line{NewFigure3(depth, startPoint, rotationDegrees), length, width}
	line.SetSubClass(line)

	return line
}

func (this *Line) Visualize() {
	graphicContext := GetTheImageGraphicContext()
	graphicContext.SetLineWidth(this.Width)
	graphicContext.LineTo(this.Length, 0)
	graphicContext.Stroke()
}
