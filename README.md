draw2dAnimation
===============

A Go package for creating 2d animations using primitives from the [draw2d library](https://code.google.com/p/draw2d/) and [FFmpeg](http://ffmpeg.org/).

Install
=======

        go get github.com/GeorgiPetkov/draw2dAnimation/draw2dAnimation

In order to use FFmpeg features you need to get ffmpeg.exe from [the official site](http://ffmpeg.zeranoe.com/builds/win64/static/ffmpeg-20140221-git-76dd01e-win64-static.7z) and place it in the same directory as the running .go file.
Note: FFmpeg features are supported only for Windows for now.

Features
========

* Create video from images set by pattern "Name%03d.png" and having the same dimensions.
* Use custom commands from FFmpeg.
* Create images.
* Images can update their objects in rich variety of ways.
* Use defined figures.
* Extend or create new figures.

Usage
=================

        import (
			"github.com/GeorgiPetkov/draw2dAnimation/draw2dAnimation"
		)


The main idea
-------------

Images contain a set of figures to be drawn and saved as frame. Each figure can be updated by rotation, translation or by given custom method (in order to simulate movement for the next frame). By custom method you can update any of the figures fields including depth(layer of the image), position, rotation, rate of translations or rotations, should the figure be filled with it's fill color and changing any other custom or given field.
You can create frames from the images and concatenate them in a video.

Extending or adding new figure
------------------------------

In order to extend the functionality you can easily add new figure types by using the base structs Figure and ComposedFigure or any of the rest presented figures which are a good example of how to do this. To extend a selected type you need to have anonymous field with pointer to that type and any custom field you may need. You have to provide a few constructors for that type, calling the constructor of the base type and calling the SetSubClass() method with 'this'. At that point all that is left is to implement the Visualize() method. In this method you should define the appearance of the object according to position (0, 0). The base class takes care of using it according to all type of updates that you have provided, setting colors and calling Stroke() or FillStroke() depending on the figure's state.
For example of a new figure see [heart.go](https://github.com/GeorgiPetkov/draw2dAnimation/blob/master/draw2dAnimation/heart.go).
For example of a new composed figure see [android.go](https://github.com/GeorgiPetkov/draw2dAnimation/blob/master/draw2dAnimation/android.go).
For example of extending an already existing figure see [rectangle.go](https://github.com/GeorgiPetkov/draw2dAnimation/blob/master/draw2dAnimation/rectangle.go) and [roundedRectangle.go](https://github.com/GeorgiPetkov/draw2dAnimation/blob/master/draw2dAnimation/roundedRectangle.go).

Using FFmpeg
------------

To use functionality of ffmpeg, different from the presented, use the ExecuteCustomFFMpegCommand() function passing the hole line excluding the prefix "ffmpeg " and an optional input (empty string if not required). For example of this see [ffmpeg.go](https://github.com/GeorgiPetkov/draw2dAnimation/blob/master/draw2dAnimation/ffmpeg.go), CreateVideoWithFrameStartNumber() function.

Using fonts
-----------

If you want to use a font for the Text figure you must provide it in ../resource/font/ according to the running .go file and the format described in function fontFileName() in [font.go](https://code.google.com/p/draw2d/source/browse/draw2d/font.go).

LICENSE
=======

The MIT License (MIT)

Copyright (c) 2013 Georgi Petkov

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
the Software, and to permit persons to whom the Software is furnished to do so,
subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
