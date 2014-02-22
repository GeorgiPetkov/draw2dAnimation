package draw2dAnimation

import (
	"fmt"
	"os/exec"
	"strings"
)

// Creates video using FFmpeg and all of the frames with the global FramePattern in inputDir. Overwrite indicates whether the file should be replaces in case that such already exists in the destination folder.
func CreateVideo(inputDir string, outputDir string, fullFileName string, inputFps float64, outputFps float64, overwrite bool) {
	CreateVideoWithFrameStartNumber(inputDir, outputDir, fullFileName, inputFps, outputFps, -1, overwrite)
}

// Creates video using FFmpeg and all of the frames with the global FramePattern in inputDir starting from StartNumber. Overwrite indicates whether the file should be replaces in case that such already exists in the destination folder.
// Information source: https://trac.ffmpeg.org/wiki/Create%20a%20video%20slideshow%20from%20images
func CreateVideoWithFrameStartNumber(inputDir string, outputDir string, fullFileName string, inputFps float64, outputFps float64, startNumber int, overwrite bool) {
	startNumberArgument := ""
	if startNumber >= 0 {
		startNumberArgument = fmt.Sprintf("-start_number %d", startNumber)
	}

	commandArgs := fmt.Sprintf(
		"-r %v %s -i %s -c:v libx264 -vf fps=%v -pix_fmt yuv420p %s",
		inputFps,
		startNumberArgument,
		inputDir+FramePattern+"%03d.png",
		outputFps,
		outputDir+fullFileName)

	input := "N"
	if overwrite {
		input = "y"
	}

	ExecuteCustomFFMpegCommand(commandArgs, input)
}

// Executes custom FFmpeg command by given arguments excluding the command name and input string if need. The ffmpeg.exe should be found in the %PATH% variable or in the current directory.
func ExecuteCustomFFMpegCommand(args string, input string) {
	cmd := exec.Command("cmd", "/C ffmpeg " + args)
	if input != "" {
		cmd.Stdin = strings.NewReader(input)
	}

	error := cmd.Run()
	if error != nil {
		panic(fmt.Sprintf("Error occured while trying to execute FFmpeg command. %v Please check ffmpeg.exe existance in PATH or the current directory and the arguments. Also make sure the command doesn't require additional input that was not provided.", error))
	}
}
