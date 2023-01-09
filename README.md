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

---

Thanks for checking this out! 😄
