package draw2dAnimation

import (
	"bufio"
	"fmt"
	imageLibrary "image"
	"image/draw"
	"image/png"
	"os"
)

type image struct {
	figures *figuresCollection
	canvas  draw.Image
}

func newImage() *image {
	return &image{
		newFiguresCollection(),
		imageLibrary.NewRGBA(imageLibrary.Rect(0, 0, FrameWidth, FrameHeight))}
}

func (this *image) AddFigure(name string, figure Figurer) {
	this.figures.add(name, figure)
}

func (this *image) RemoveFigure(name string) {
	this.figures.remove(name)
}

func (this *image) Update() {
	this.figures.traverse(func(figure Figurer) {
		figure.Update()
	})
}

func (this *image) Draw() {
	this.figures.traverse(func(figure Figurer) {
		figure.Draw()
	})
}

func (this *image) Clear() {
	graphicContext := GetTheImageGraphicContext()
	graphicContext.Clear()
}

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
