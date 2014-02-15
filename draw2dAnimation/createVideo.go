package draw2dAnimation

import (
	"fmt"
	"os/exec"
	"strings"
)

func CreateVideo(inputDir string, outputDir string, fullFileName string, inputFps float64, outputFps float64, overwrite bool) {
	commandArgs := fmt.Sprintf(
		"-r %v -i %s -c:v libx264 -vf fps=%v -pix_fmt yuv420p %s",
		inputFps,
		inputDir+FramePattern+"%03d.png",
		outputFps,
		outputDir+fullFileName)
	if overwrite {
		ExecuteCustomFFMpegCommand(commandArgs, "y")
	} else {
		ExecuteCustomFFMpegCommand(commandArgs, "N")
	}
}

func ExecuteCustomFFMpegCommand(args string, input string) {
	cmd := exec.Command(
		"cmd", "/C ffmpeg "+args)
	cmd.Stdin = strings.NewReader(input)
	error := cmd.Run()
	fmt.Println(error)
}
