package draw2dAnimation

// A global variable holding the default destination for produces files.
var DestinationFolder string = "../result/"

// A global variable holding the name of the produced frames which will be followed by their number.
var FramePattern string = "Frame"

// A global variable holding the number of the frame before the next. Used for naming of frames.
var nextFrameNumber = -1

// A global variable holding the default width of the created frames.
var FrameWidth int = 640

// A global variable holding the default height of the created frames.
var FrameHeight int = 480

// A global variable holding the number of the figure before the next. Used to set unique ID to each figure.
var nextFigureId int = -1
