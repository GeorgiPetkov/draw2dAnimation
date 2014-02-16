package draw2dAnimation

import (
	"code.google.com/p/draw2d/draw2d"
)

// singletons
var imageInstance *image = nil
var graphicContextInstance *draw2d.ImageGraphicContext = nil

// Gets the singleton image.
func GetTheImage() *image {
	if imageInstance == nil {
		imageInstance = newImage()
		imageInstance.Clear()
	}

	return imageInstance
}

// Gets the singleton ImageGraphicContext.
func GetTheImageGraphicContext() *draw2d.ImageGraphicContext {
	if graphicContextInstance == nil {
		graphicContextInstance = draw2d.NewGraphicContext(GetTheImage().canvas)
	}

	return graphicContextInstance
}
