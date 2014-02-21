package draw2dAnimation

import (
	"code.google.com/p/draw2d/draw2d"
)

// A figure type that allows visualization of a text with custom font. Change font size and data to adjust the text visulization.
type Text struct {
	*Figure
	FontData draw2d.FontData
	FontSize float64
	Text     string
}

// Constructor setting current struct's fields and default values for the base struct
func NewText(fontData draw2d.FontData, fontSize float64, text string) *Text {
	textFigure := &Text{NewFigure(), fontData, fontSize, text}
	textFigure.SetSubClass(textFigure)

	return textFigure
}

// Constructor setting both base struct's and current struct's fields.
func NewText5(fontData draw2d.FontData, fontSize float64, text string,
	depth int, startPoint Point, rotationDegrees float64) *Text {
	textFigure := &Text{NewFigure(), fontData, fontSize, text}
	textFigure.SetDepth(depth)
	textFigure.SetStartPoint(startPoint)
	textFigure.SetRotationDegrees(rotationDegrees)
	textFigure.SetSubClass(textFigure)

	return textFigure
}

// Defines the visualization of the figure according to position (0, 0).
func (this *Text) Visualize() {
	graphicContext := GetTheImageGraphicContext()
	graphicContext.SetFontSize(this.FontSize)
	graphicContext.SetFontData(this.FontData)

	_, top, _, bottom := graphicContext.GetStringBounds(this.Text)
	graphicContext.Translate(0, bottom-top)

	if this.isFilled {
		graphicContext.FillString(this.Text)
	} else {
		graphicContext.StrokeString(this.Text)
	}
}
