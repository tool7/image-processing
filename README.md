# Image Processing in Go

Image processing desktop app built with Wails (https://wails.io/) and Vue.js framework.
All image operations are implemented on Go "backend" which than communicates the changes to Vue frontend.

---

#### Setup

Required dependencies are Go (https://go.dev/), Node (16.2.0+) (https://nodejs.org/en/download/) and Wails CLI (https://wails.io/docs/gettingstarted/installation).
After installing those:

1. Navigate to `frontend` folder and install Vue dependencies with `npm install` command
2. Navigate back to root folder and install Go dependencies with `go get .` command
3. Run project with `wails dev` command

You can also use `wails build` command for compiling the project to executable.

#### Screenshots

<img src="https://github.com/tool7/image-processing/blob/main/screenshots/screenshot-1.png" width="400" height="300">
<img src="https://github.com/tool7/image-processing/blob/main/screenshots/screenshot-2.png" width="400" height="300">
<img src="https://github.com/tool7/image-processing/blob/main/screenshots/screenshot-3.png" width="400" height="300">
<img src="https://github.com/tool7/image-processing/blob/main/screenshots/screenshot-4.png" width="400" height="300">
<img src="https://github.com/tool7/image-processing/blob/main/screenshots/screenshot-5.png" width="400" height="300">
<img src="https://github.com/tool7/image-processing/blob/main/screenshots/screenshot-6.png" width="400" height="300">

#### Available operations

Implemented operations include `brightness`, `contrast`, `saturation`, `tint`, `greyscale`, `negative`, `sepia`, `box blur`, `motion blur`, `sharpen`, `emboss`, horizontal and vertical `edge detection`, and `outline`.
Basic image transformations are also implemented: `rotation` (90°, 180°, -90°) and `mirroring` (horizontal and vertical).

#### Limitations

App cannot handle processing images that are too large, so for a good performance please select images that are Full HD (1080p) or below that.

#### Potential optimizations

Each operation could be processed on the so called “preview” image, a much smaller copy of the original image. The only part that needs to be done on the original image is exporting of PNG image. The problem that needs to be solved is with kernel operations - e.g. Blur operation would have a much larger effect on the smaller “preview” image compared to the result of processing original image (when exporting PNG) - reason being hardcoded kernel sizes.

---

Thanks for checking this out! 😄
