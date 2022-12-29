package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/png"

	"tool7/image-processing/models"
	"tool7/image-processing/operations"
	"tool7/image-processing/utils"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type App struct {
	ctx                  context.Context
	originalImage        *image.RGBA
	imageLayerCollection *models.ImageLayerCollection
}

type ProcessedImage struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Base64 string `json:"base64"`
}

type ImageOperationType int

const (
	Brightness ImageOperationType = iota + 1
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

type ImageColor struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
	A uint8 `json:"a"`
}

type ImageOperation struct {
	Type  ImageOperationType `json:"type"`
	Level float64            `json:"level,omitempty"`
	Tint  ImageColor         `json:"tint,omitempty"`
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

	imageLayerCollection := utils.NewImageLayerCollection(a.originalImage)
	a.imageLayerCollection = imageLayerCollection

	return true
}

func (a *App) ProcessImage() ProcessedImage {
	var buff bytes.Buffer

	if a.imageLayerCollection.Size > 0 {
		processedImage, err := a.imageLayerCollection.ExecuteLayersFrom(0)
		if err != nil {
			panic(err)
		}
		png.Encode(&buff, processedImage)
	} else {
		png.Encode(&buff, a.originalImage)
	}

	rawBase64String := base64.StdEncoding.EncodeToString(buff.Bytes())
	base64String := "data:image/png;base64,"
	base64String += rawBase64String

	return ProcessedImage{
		Width:  a.originalImage.Bounds().Max.X - a.originalImage.Bounds().Min.X,
		Height: a.originalImage.Bounds().Max.Y - a.originalImage.Bounds().Min.Y,
		Base64: base64String,
	}
}

func (a *App) AppendImageOperation(operation ImageOperation) bool {
	switch operation.Type {
	case Brightness:
		brightnessOperation := operations.NewBrightnessOperation(operation.Level)
		brightnessLayer := utils.NewImageLayer(brightnessOperation)
		a.imageLayerCollection.Append(brightnessLayer)
		break
	case Contrast:
		contrastOperation := operations.NewContrastOperation(operation.Level)
		contrastLayer := utils.NewImageLayer(contrastOperation)
		a.imageLayerCollection.Append(contrastLayer)
		break
	case Saturation:
		saturationOperation := operations.NewSaturationOperation(operation.Level)
		saturationLayer := utils.NewImageLayer(saturationOperation)
		a.imageLayerCollection.Append(saturationLayer)
		break
	case Tint:
		applyTintOperation := operations.NewTintOperation(color.RGBA{
			R: operation.Tint.R,
			G: operation.Tint.G,
			B: operation.Tint.B,
			A: operation.Tint.A,
		}, operation.Level)

		tintLayer := utils.NewImageLayer(applyTintOperation)
		a.imageLayerCollection.Append(tintLayer)
		break
	case Greyscale:
		greyscaleOperation := operations.NewGreyscaleOperation()
		greyscaleLayer := utils.NewImageLayer(greyscaleOperation)
		a.imageLayerCollection.Append(greyscaleLayer)
		break
	case Negative:
		negativeOperation := operations.NewNegativeOperation()
		negativeLayer := utils.NewImageLayer(negativeOperation)
		a.imageLayerCollection.Append(negativeLayer)
		break
	case Sepia:
		sepiaOperation := operations.NewSepiaOperation()
		sepiaLayer := utils.NewImageLayer(sepiaOperation)
		a.imageLayerCollection.Append(sepiaLayer)
		break
	case Emboss:
		break
	case EdgesVertical:
		break
	case EdgesHorizontal:
		break
	case MirrorVertical:
		break
	case MirrorHorizontal:
		break
	case RotateBy90:
		break
	case RotateBy180:
		break
	case RotateBy270:
		break
	}

	fmt.Println("=================================== AppendImageOperation")
	fmt.Println(a.imageLayerCollection.Size)

	return true
}
