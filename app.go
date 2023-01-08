package main

import (
	"bytes"
	"context"
	"encoding/base64"
	"image"
	"image/color"
	"image/png"
	"os"
	"strings"

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

type Base64Image struct {
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
	BoxBlur
	MotionBlur
	Sharpen
	Emboss
	EdgesHorizontal
	EdgesVertical
	Outline
)

type TintRGB struct {
	R uint8 `json:"r"`
	G uint8 `json:"g"`
	B uint8 `json:"b"`
}

type ImageOperation struct {
	Type       ImageOperationType `json:"type"`
	Level      float64            `json:"level,omitempty"`
	Tint       TintRGB            `json:"tint,omitempty"`
	KernelSize models.KernelSize  `json:"kernelSize,omitempty"`
	IsEnabled  bool               `json:"isEnabled"`
}

func NewApp() *App {
	return &App{}
}

func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) initProject(img *image.RGBA) {
	a.originalImage = img

	imageLayerCollection := utils.NewImageLayerCollection(a.originalImage)
	a.imageLayerCollection = imageLayerCollection
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

	a.initProject(img)
	return true
}

func (a *App) GetOriginalImage() Base64Image {
	var buff bytes.Buffer
	png.Encode(&buff, a.originalImage)

	rawBase64String := base64.StdEncoding.EncodeToString(buff.Bytes())
	base64String := "data:image/png;base64,"
	base64String += rawBase64String

	return Base64Image{
		Width:  a.originalImage.Bounds().Max.X - a.originalImage.Bounds().Min.X,
		Height: a.originalImage.Bounds().Max.Y - a.originalImage.Bounds().Min.Y,
		Base64: base64String,
	}
}

func (a *App) SetOriginalImage(imageBase64 string) {
	reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(imageBase64))
	decodedImage, _, err := image.Decode(reader)

	if err != nil {
		panic(err)
	}

	rgbaImage := image.NewRGBA(decodedImage.Bounds())
	for y := decodedImage.Bounds().Min.Y; y < decodedImage.Bounds().Max.Y; y++ {
		for x := decodedImage.Bounds().Min.X; x < decodedImage.Bounds().Max.X; x++ {
			r, g, b, a := utils.GetPixelColor(decodedImage, x, y)
			rgbaImage.SetRGBA(x, y, color.RGBA{r, g, b, a})
		}
	}

	a.initProject(rgbaImage)
}

func (a *App) GetUserSelectedProjectFileContent() string {
	filePath, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Project File (.goimp extension)",
		Filters: []runtime.FileFilter{
			{
				DisplayName: "Files (*.goimp)",
				Pattern:     "*.goimp",
			},
		},
	})
	if err != nil {
		panic("Error on project file selection")
	}
	if filePath == "" {
		return ""
	}

	byteData, err := os.ReadFile(filePath)
	if err != nil {
		panic("Error reading project file")
	}

	return string(byteData)
}

func (a *App) ProcessImage(indexToExecuteFrom int) Base64Image {
	var buff bytes.Buffer

	if a.imageLayerCollection.Size > 0 {
		processedImage, err := a.imageLayerCollection.ExecuteLayersFrom(indexToExecuteFrom)
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

	return Base64Image{
		Width:  a.originalImage.Bounds().Max.X - a.originalImage.Bounds().Min.X,
		Height: a.originalImage.Bounds().Max.Y - a.originalImage.Bounds().Min.Y,
		Base64: base64String,
	}
}

func (a *App) ResetAppState() {
	a.originalImage = nil
	a.imageLayerCollection = nil
}

func (a *App) AppendImageOperation(operation ImageOperation) error {
	imageLayer, err := CreateImageLayerWithOperation(operation)
	if err != nil {
		panic(err)
	}

	imageLayer.IsEnabled = operation.IsEnabled
	a.imageLayerCollection.Append(imageLayer)

	return nil
}

func (a *App) RemoveImageOperationAtIndex(index int) error {
	return a.imageLayerCollection.RemoveAt(index)
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
	case BoxBlur, MotionBlur, Sharpen, Emboss:
		kernelOperation, ok := imageLayer.Operation.(*operations.KernelOperation)
		if !ok {
			panic("Failed to cast to KernelOperation")
		}
		kernelOperation.KernelSize = operation.KernelSize
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

func (a *App) MoveImageOperation(oldIndex, newIndex int) error {
	if oldIndex == newIndex {
		return nil
	}

	imageLayer, err := a.imageLayerCollection.At(oldIndex)
	if err != nil {
		panic(err)
	}

	a.imageLayerCollection.RemoveAt(oldIndex)
	a.imageLayerCollection.InsertAt(imageLayer, newIndex)

	return nil
}

func (a *App) ToggleImageOperation(index int) error {
	imageLayer, err := a.imageLayerCollection.At(index)
	if err != nil {
		panic(err)
	}

	if imageLayer.IsEnabled {
		imageLayer.Disable()
	} else {
		imageLayer.Enable()
	}

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
	case BoxBlur:
		boxBlurOperation := operations.NewKernelOperation(models.BoxBlur, operation.KernelSize)
		return utils.NewImageLayer(boxBlurOperation), nil
	case MotionBlur:
		motionBlurOperation := operations.NewKernelOperation(models.MotionBlur, operation.KernelSize)
		return utils.NewImageLayer(motionBlurOperation), nil
	case Sharpen:
		sharpenOperation := operations.NewKernelOperation(models.Sharpen, operation.KernelSize)
		return utils.NewImageLayer(sharpenOperation), nil
	case Emboss:
		embossOperation := operations.NewKernelOperation(models.Emboss, operation.KernelSize)
		return utils.NewImageLayer(embossOperation), nil
	case EdgesHorizontal:
		horizontalEdgeDetectionOperation := operations.NewKernelOperation(models.EdgeDetectionHorizontal, operation.KernelSize)
		return utils.NewImageLayer(horizontalEdgeDetectionOperation), nil
	case EdgesVertical:
		verticalEdgeDetectionOperation := operations.NewKernelOperation(models.EdgeDetectionVertical, operation.KernelSize)
		return utils.NewImageLayer(verticalEdgeDetectionOperation), nil
	case Outline:
		outlineOperation := operations.NewKernelOperation(models.Outline, operation.KernelSize)
		return utils.NewImageLayer(outlineOperation), nil
	}

	return nil, errors.New("Failed to create ImageLayer with provided ImageOperation")
}

func (a *App) RotateImageBy90Deg() error {
	rotationOperation := operations.NewRotationOperation(operations.By90Deg)
	rotatedImage, err := rotationOperation.Execute(a.originalImage)

	if err != nil {
		return err
	}

	a.originalImage = rotatedImage
	a.imageLayerCollection.InputImage = rotatedImage

	return nil
}

func (a *App) MirrorImageVertically() error {
	verticalMirrorOperation := operations.NewVerticalMirrorOperation()
	mirroredImage, err := verticalMirrorOperation.Execute(a.originalImage)

	if err != nil {
		return err
	}

	a.originalImage = mirroredImage
	a.imageLayerCollection.InputImage = mirroredImage

	return nil
}

func (a *App) MirrorImageHorizontally() error {
	horizontalMirrorOperation := operations.NewHorizontalMirrorOperation()
	mirroredImage, err := horizontalMirrorOperation.Execute(a.originalImage)

	if err != nil {
		return err
	}

	a.originalImage = mirroredImage
	a.imageLayerCollection.InputImage = mirroredImage

	return nil
}
