package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/png"

	"tool7/image-processing/utils"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx           context.Context
	originalImage *image.RGBA
}

type ProcessedImage struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Base64 string `json:"base64"`
}

type ImageOperationType int

const (
	Brightness ImageOperationType = iota
	Contrast
	Saturation
	Tint
	Greyscale
	Negative
	Sepia
	Emboss
	EdgesVertical
	EdgesHorizontal
	MirrorVertical
	MirrorHorizontal
	RotateBy90
	RotateBy180
	RotateBy270
)

func (o ImageOperationType) String() string {
	return [...]string{
		"brightness",
		"contrast",
		"saturation",
		"tint",
		"greyscale",
		"negative",
		"sepia",
		"emboss",
		"edges-vertical",
		"edges-horizontal",
		"mirror-vertical",
		"mirror-horizontal",
		"rotate-90",
		"rotate-180",
		"rotate-270",
	}[o]
}

type ImageOperation struct {
	Type ImageOperationType `json:"type"`
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) OpenImageFileSelector() bool {
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Image File (PNG or JPG)",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Images (*.png;*.jpg)",
				Pattern:     "*.png;*.jpg",
			},
		},
	})

	if err != nil {
		panic("Error on image file selection")
	}

	if filePath == "" {
		return false
	}

	img, err := utils.GetImageFromFilePath(filePath)
	if err != nil {
		panic(err)
	}

	a.originalImage = img
	return true
}

func (a *App) ProccessImage() ProcessedImage {
	var buff bytes.Buffer
	png.Encode(&buff, a.originalImage)

	rawBase64String := base64.StdEncoding.EncodeToString(buff.Bytes())
	base64String := "data:image/png;base64,"
	base64String += rawBase64String

	return ProcessedImage{
		Width:  a.originalImage.Bounds().Max.X - a.originalImage.Bounds().Min.X,
		Height: a.originalImage.Bounds().Max.Y - a.originalImage.Bounds().Min.Y,
		Base64: base64String,
	}
}

func (a *App) AddImageOperation(operation ImageOperation) bool {
	fmt.Println(operation.Type.String())

	return true
}
