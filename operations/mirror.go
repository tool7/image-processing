package operations

import (
	"image"
	"image/color"

	utils "tool7/image-processing/utils"
)

type VerticalMirrorOperation struct{}
type HorizontalMirrorOperation struct{}

func NewVerticalMirrorOperation() *VerticalMirrorOperation {
	return &VerticalMirrorOperation{}
}

func NewHorizontalMirrorOperation() *HorizontalMirrorOperation {
	return &HorizontalMirrorOperation{}
}

func (this *VerticalMirrorOperation) Execute(inputImage *image.RGBA) (*image.RGBA, error) {
	return CreateVerticalMirror(inputImage), nil
}

func (this *HorizontalMirrorOperation) Execute(inputImage *image.RGBA) (*image.RGBA, error) {
	return CreateHorizontalMirror(inputImage), nil
}

func CreateVerticalMirror(inputImage *image.RGBA) *image.RGBA {
	result := image.NewRGBA(inputImage.Bounds())
	minX := inputImage.Bounds().Min.X
	minY := inputImage.Bounds().Min.Y
	maxX := inputImage.Bounds().Max.X
	maxY := inputImage.Bounds().Max.Y

	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			r, g, b, a := utils.GetPixelColor(inputImage, x, y)

			result.SetRGBA(inputImage.Bounds().Max.X-x, y, color.RGBA{r, g, b, a})
		}
	}

	return result
}

func CreateHorizontalMirror(inputImage *image.RGBA) *image.RGBA {
	result := image.NewRGBA(inputImage.Bounds())
	minX := inputImage.Bounds().Min.X
	minY := inputImage.Bounds().Min.Y
	maxX := inputImage.Bounds().Max.X
	maxY := inputImage.Bounds().Max.Y

	for x := minX; x < maxX; x++ {
		for y := minY; y < maxY; y++ {
			r, g, b, a := utils.GetPixelColor(inputImage, x, y)

			result.SetRGBA(x, inputImage.Bounds().Max.Y-y, color.RGBA{r, g, b, a})
		}
	}

	return result
}
