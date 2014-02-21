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

// Adds figure with string key to the contained collection.
func (this *image) AddFigure(name string, figure Figurer) {
	this.figures.add(name, figure)
}

// Removes figure by string key from the contained collection.
func (this *image) RemoveFigure(name string) {
	this.figures.remove(name)
}

// Removes all figures from the contained collection passing the given filter.
func (this *image) RemoveByFilter(filter func(Figurer) bool) {
	this.figures.removeByFilter(filter)
}

// Gets the figure corresponding to the given string key in the contained collection or nil if not found.
func (this *image) GetByName(name string) Figurer {
	return this.figures.getByName(name)
}

// Gets the string keys and the figures in the contained collection passing the given filter.
func (this *image) GetByFilter(filter func(Figurer) bool) map[string]Figurer {
	return this.figures.getByFilter(filter)
}

// Updates all figures in the contained collection.
func (this *image) Update() {
	this.figures.traverse(func(figure Figurer) {
		figure.Update()
	})
}

// Draws all figures in the contained collection in their order by depth then by order of creation.
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
