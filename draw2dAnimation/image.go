package draw2dAnimation

import (
	"bufio"
	"fmt"
	imageLibrary "image"
	"image/draw"
	"image/png"
	"image/color"
	"os"
)

// The image struct contains a set of all figures and take operations over them like drawing, updating, adding, deleting or filtering figures. Can save the result as a .png file. The default clear color is white.
type image struct {
	figures *figuresCollection
	canvas  draw.Image
	ClearColor color.Color
}

// Clears the image using the setted clear color.
func (this *image) Clear() {
	width, height := this.canvas.Bounds().Dx(), this.canvas.Bounds().Dy()
	this.ClearRectangle(0, 0, width, height)
}

// Clears a rectangle area of the image using the setted clear color.
func (this *image) ClearRectangle(x1, y1, x2, y2 int) {
	imageColor := imageLibrary.NewUniform(this.ClearColor)
	draw.Draw(this.canvas, imageLibrary.Rect(x1, y1, x2, y2), imageColor, imageLibrary.ZP, draw.Over)
}

// Default constructor for the struct starting with empty collection of figures and setting the canvas reference.
func newImage() *image {
	return &image{
		newFiguresCollection(),
		imageLibrary.NewRGBA(imageLibrary.Rect(0, 0, FrameWidth, FrameHeight)),
		color.RGBA{255, 255, 255, 255}}
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
