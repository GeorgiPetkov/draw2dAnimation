package draw2dAnimation

import (
	"code.google.com/p/draw2d/draw2d"
)

const (
	TopMargin int = iota
	BottomMargin
	LeftMargin
	RightMargin
)

type TextWithFrame struct {
	*ComposedFigure
	Margins []float64
}

// Constructor accepting initialized base class and creating text with rectangular frame and equal margins for all sides.
func NewTextWithFrame(fontData draw2d.FontData, fontSize float64, text string, margin float64,
	base *ComposedFigure) *TextWithFrame {
	return NewTextWithFrameCustomMargins(fontData, fontSize, text, []float64{margin, margin, margin, margin}, base)
}

// Constructor accepting initialized base class and creating text with rectangular frame and custom margin for each sides.
func NewTextWithFrameCustomMargins(fontData draw2d.FontData, fontSize float64, text string, margins []float64,
	base *ComposedFigure) *TextWithFrame {
	return newTextWithFrame(fontData, fontSize, text, margins, false, 0.0, base)
}

// Constructor accepting initialized base class and radius and creating text with rectangular frame with round edges and equal margins for all sides.
func NewTextWithRoundFrame(fontData draw2d.FontData, fontSize float64, text string, margin float64, radius float64,
	base *ComposedFigure) *TextWithFrame {
	return NewTextWithRoundFrameCustomMargins(fontData, fontSize, text, []float64{margin, margin, margin, margin}, radius, base)
}

// Constructor accepting initialized base class and radius and creating text with rectangular frame with round edges and custom margin for each sides.
func NewTextWithRoundFrameCustomMargins(fontData draw2d.FontData, fontSize float64, text string, margins []float64, radius float64,
	base *ComposedFigure) *TextWithFrame {
	return newTextWithFrame(fontData, fontSize, text, margins, true, radius, base)
}

// Called by constructors to set initial state of the figure.
func newTextWithFrame(
	fontData draw2d.FontData, fontSize float64, text string, margins []float64, roundedFrame bool, radius float64,
	base *ComposedFigure) *TextWithFrame {
	textWithFrame := &TextWithFrame{base, margins}
	
	textFigure := NewText6(fontData, fontSize, text, 1, Point{margins[LeftMargin], margins[TopMargin]}, 0.0)
	textWithFrame.AddFigure("text", textFigure)
	
	graphicContext := GetTheImageGraphicContext()
	graphicContext.SetFontSize(fontSize)
	graphicContext.SetFontData(fontData)
	left, top, right, bottom := graphicContext.GetStringBounds(text)
	
	width := right - left + margins[LeftMargin] + margins[RightMargin]
	height:= bottom - top + margins[TopMargin] + margins[BottomMargin]
	
	var frame Rectangler
	if roundedFrame {
		frame = NewRoundRectangle(radius, width, height, textWithFrame.GetLineWidth())
	} else {
		frame = NewRectangle(width, height, base.GetLineWidth())
	}
	
	textWithFrame.AddFigure("frame", frame)
	
	return textWithFrame
}