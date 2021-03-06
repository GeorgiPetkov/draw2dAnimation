package main

import (
	"github.com/GeorgiPetkov/draw2dAnimation/draw2dAnimation"
	"image/color"
)

var leftIsUp bool = true
var stepSize float64 = 10.0

func MoveLeftLeg(figure draw2dAnimation.Figurer) {
	MoveLeg(figure, !leftIsUp)
}

func MoveRightLeg(figure draw2dAnimation.Figurer) {
	MoveLeg(figure, leftIsUp)
}

func MoveLeg(figure draw2dAnimation.Figurer, up bool) {
	step := stepSize
	if !up {
		step = -step
	}

	figure.SetStartPoint(draw2dAnimation.Point{figure.GetStartPoint().X, figure.GetStartPoint().Y + step})
}

func SwapLegs(figure draw2dAnimation.Figurer) {
	leftIsUp = !leftIsUp
}

func main() {
	android := draw2dAnimation.NewAndroid(100.0, 80.0, 3.0, color.RGBA{0, 255, 0, 255})
	android.SetStartPoint(draw2dAnimation.Point{100, 100})
	android.SetUpdateTranslation(draw2dAnimation.Point{5.0, 0.0})
	android.SetScale(draw2dAnimation.Point{2.0, 1.0})
	
	image := draw2dAnimation.GetTheImage()

	image.AddFigure("TheOneAndOnly", android)

	leftLeg := android.GetFigureByName("leftLeg")
	rightLeg := android.GetFigureByName("rightLeg")

	// move left leg up
	leftLeg.SetStartPoint(draw2dAnimation.Point{leftLeg.GetStartPoint().X, leftLeg.GetStartPoint().Y - stepSize})

	android.SetUpdateMethod(draw2dAnimation.NewUpdateMethod(SwapLegs))
	leftLeg.SetUpdateMethod(draw2dAnimation.NewUpdateMethod(MoveLeftLeg))
	rightLeg.SetUpdateMethod(draw2dAnimation.NewUpdateMethod(MoveRightLeg))

	for i := 0; i < 100; i += 1 {
		image.Draw()
		image.Update()
		image.SaveFrame()
		image.Clear()
	}

	draw2dAnimation.CreateVideo("../result/", "../result/out.mp4", 16.0, 25, true)
}
