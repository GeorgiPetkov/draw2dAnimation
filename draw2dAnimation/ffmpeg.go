package draw2dAnimation

import (
	"fmt"
	"os/exec"
	"strings"
)

// Creates video using FFmpeg. Overwrite indicates whether the file should be replaces in case that such already exists in the destination folder.
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

// Executes custom FFmpeg command by given arguments excluding the prefix "ffmpeg " and input string if need.
func ExecuteCustomFFMpegCommand(args string, input string) {
	cmd := exec.Command("cmd", "/C ffmpeg "+args)
	//if input != "" {
	cmd.Stdin = strings.NewReader(input)
	//}

	error := cmd.Run()
	fmt.Println(error)
}
