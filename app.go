package main

import (
	"bytes"
	"context"
	"encoding/base64"
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

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) OpenImageFileSelector() {
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Image File (PNG or JPG)",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Images (*.png;*.jpg)",
				Pattern:     "*.png;*.jpg",
			},
		},
	})

	if err != nil || filePath == "" {
		panic("Error on image file selection")
	}

	img, err := utils.GetImageFromFilePath(filePath)
	if err != nil {
		panic(err)
	}

	a.originalImage = img
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
