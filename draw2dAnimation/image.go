package draw2dAnimation

import (
	"bufio"
	"fmt"
	imageLibrary "image"
	"image/draw"
	"image/png"
	"os"
)

// The image struct contains a set of all figures and take operations over them like drawing, updating, adding and deleting figures. Can save the result as a .png file.
type image struct {
	figures *figuresCollection
	canvas  draw.Image
}

// Default constructor for the struct starting with empty collection of figures and setting the canvas reference.
func newImage() *image {
	return &image{
		newFiguresCollection(),
		imageLibrary.NewRGBA(imageLibrary.Rect(0, 0, FrameWidth, FrameHeight))}
}

// Adds figure with string key to the collection.
func (this *image) AddFigure(name string, figure Figurer) {
	this.figures.add(name, figure)
}

// Removes figure by string key from the collection.
func (this *image) RemoveFigure(name string) {
	this.figures.remove(name)
}

// Gets the figure corresponding to the given string key in the collection or nil if not found.
func (this *image) getByName(name string) Figurer {
	return this.figures.getByName(name)
}

// Updates all contained figures.
func (this *image) Update() {
	this.figures.traverse(func(figure Figurer) {
		figure.Update()
	})
}

// Draws all contained figures in their order by depth then by order of creation.
func (this *image) Draw() {
	this.figures.traverse(func(figure Figurer) {
		figure.Draw()
	})
}

// Clears the canvas of the image.
func (this *image) Clear() {
	graphicContext := GetTheImageGraphicContext()
	graphicContext.Clear()
}

// Saves the result as a .png file using the DestinationFolder and FramePattern global variables.
func (this *image) SaveFrame() {
	nextFrameNumber++

	file, err := os.Create(
		fmt.Sprintf("%s%s%03d.png", DestinationFolder, FramePattern, nextFrameNumber))
	if err != nil {
		os.Exit(1)
	}

	defer file.Close()
	buffer := bufio.NewWriter(file)
	err = png.Encode(buffer, this.canvas)
	if err != nil {
		os.Exit(1)
	}

	err = buffer.Flush()
	if err != nil {
		os.Exit(1)
	}
}
