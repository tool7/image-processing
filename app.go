package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"image"
	"image/color"
	"image/png"

	"tool7/image-processing/models"
	"tool7/image-processing/operations"
	"tool7/image-processing/utils"

	"github.com/pkg/errors"
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

type TintRGB struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
}

type ImageOperation struct {
	Type  ImageOperationType `json:"type"`
	Level float64            `json:"level,omitempty"`
	Tint  TintRGB            `json:"tint,omitempty"`
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

func (a *App) ResetAppState() {
	imageLayerCollection := utils.NewImageLayerCollection(a.originalImage)
	a.imageLayerCollection = imageLayerCollection
}

func (a *App) AppendImageOperation(operation ImageOperation) error {
	imageLayer, err := CreateImageLayerWithOperation(operation)
	if err != nil {
		panic(err)
	}

	a.imageLayerCollection.Append(imageLayer)

	return nil
}

func (a *App) RemoveImageOperationAtIndex(index int) error {
	a.imageLayerCollection.RemoveAt(index)

	return nil
}

func (a *App) UpdateImageOperationAtIndex(index int, operation ImageOperation) error {
	imageLayer, err := a.imageLayerCollection.At(index)
	if err != nil {
		panic(err)
	}

	switch operation.Type {
	case Brightness:
		brightnessOperation, ok := imageLayer.Operation.(*operations.BrightnessOperation)
		if !ok {
			panic("Failed to cast to BrightnessOperation")
		}
		brightnessOperation.Level = operation.Level
		break
	case Contrast:
		contrastOperation, ok := imageLayer.Operation.(*operations.ContrastOperation)
		if !ok {
			panic("Failed to cast to ContrastOperation")
		}
		contrastOperation.Factor = operation.Level
		break
	case Saturation:
		saturationOperation, ok := imageLayer.Operation.(*operations.SaturationOperation)
		if !ok {
			panic("Failed to cast to SaturationOperation")
		}
		saturationOperation.Level = operation.Level
		break
	case Tint:
		tintOperation, ok := imageLayer.Operation.(*operations.TintOperation)
		if !ok {
			panic("Failed to cast to TintOperation")
		}
		tintOperation.Intensity = operation.Level
		tintOperation.Tint = color.RGBA{
			R: operation.Tint.R,
			G: operation.Tint.G,
			B: operation.Tint.B,
			A: 255,
		}
		break
	}

	return nil
}

func (a *App) ReplaceImageOperationAtIndex(index int, operation ImageOperation) error {
	imageLayer, err := CreateImageLayerWithOperation(operation)
	if err != nil {
		panic(err)
	}

	a.imageLayerCollection.RemoveAt(index)
	a.imageLayerCollection.InsertAt(imageLayer, index)

	return nil
}

func CreateImageLayerWithOperation(operation ImageOperation) (*models.ImageLayer, error) {
	switch operation.Type {
	case Brightness:
		brightnessOperation := operations.NewBrightnessOperation(operation.Level)
		return utils.NewImageLayer(brightnessOperation), nil
	case Contrast:
		contrastOperation := operations.NewContrastOperation(operation.Level)
		return utils.NewImageLayer(contrastOperation), nil
	case Saturation:
		saturationOperation := operations.NewSaturationOperation(operation.Level)
		return utils.NewImageLayer(saturationOperation), nil
	case Tint:
		applyTintOperation := operations.NewTintOperation(color.RGBA{
			R: operation.Tint.R,
			G: operation.Tint.G,
			B: operation.Tint.B,
			A: 255,
		}, operation.Level)

		return utils.NewImageLayer(applyTintOperation), nil
	case Greyscale:
		greyscaleOperation := operations.NewGreyscaleOperation()
		return utils.NewImageLayer(greyscaleOperation), nil
	case Negative:
		negativeOperation := operations.NewNegativeOperation()
		return utils.NewImageLayer(negativeOperation), nil
	case Sepia:
		sepiaOperation := operations.NewSepiaOperation()
		return utils.NewImageLayer(sepiaOperation), nil
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

	return nil, errors.New("Failed to create ImageLayer with provided ImageOperation")
}
