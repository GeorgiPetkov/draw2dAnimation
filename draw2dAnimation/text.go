package draw2dAnimation

import (
	"code.google.com/p/draw2d/draw2d"
)

type Text struct {
	*Figure
	FontData draw2d.FontData
	FontSize float64
	Text     string
}

func NewText(fontData draw2d.FontData, fontSize float64, text string) *Text {
	textFigure := &Text{NewFigure(), fontData, fontSize, text}
	textFigure.SetSubClass(textFigure)

	return textFigure
}

func NewText5(fontData draw2d.FontData, fontSize float64, text string,
	depth int, startPoint Point, rotationDegrees float64) *Text {
	textFigure := &Text{NewFigure3(depth, startPoint, 0), fontData, fontSize, text}
	textFigure.SetSubClass(textFigure)

	return textFigure
}

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

/*func (this *Text) Visualize() {
	graphicContext := GetTheImageGraphicContext()
	graphicContext.FillStroke()
	graphicContext.MoveTo(0, 0)
	graphicContext.SetFontSize(this.FontSize)
	graphicContext.SetFontData(this.FontData)
	graphicContext.FillString(this.Text)
}*/
