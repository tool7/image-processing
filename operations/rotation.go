package operations

import (
	"errors"
	"image"
	"image/color"

	utils "tool7/image-processing/utils"
)

type RotationDegrees int

const (
	By90Deg RotationDegrees = iota
	By180Deg
	By270Deg
)

type RotationOperation struct {
	degrees RotationDegrees
}

func NewRotationOperation(degrees RotationDegrees) *RotationOperation {
	return &RotationOperation{
		degrees,
	}
}

func (this *RotationOperation) Execute(inputImage *image.RGBA) (*image.RGBA, error) {
	switch this.degrees {
	case By90Deg:
		return RotateBy90Deg(inputImage), nil
	case By180Deg:
		return RotateBy180Deg(inputImage), nil
	case By270Deg:
		return RotateBy270Deg(inputImage), nil
	}

	return nil, errors.New("Invalid RotationDegrees value")
}

func RotateBy90Deg(img *image.RGBA) *image.RGBA {
	imageMinX := img.Bounds().Min.X
	imageMinY := img.Bounds().Min.Y
	imageMaxX := img.Bounds().Max.X
	imageMaxY := img.Bounds().Max.Y

	result := image.NewRGBA(image.Rect(imageMinX, imageMinY, imageMaxY, imageMaxX))

	// Transposing the matrix
	for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
		for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
			r, g, b, a := utils.GetPixelColor(img, x, y)

			result.SetRGBA(y, x, color.RGBA{r, g, b, a})
		}
	}

	// Reversing columns
	return CreateVerticalMirror(result)
}

func RotateBy180Deg(img *image.RGBA) *image.RGBA {
	return CreateVerticalMirror(CreateHorizontalMirror(img))
}

func RotateBy270Deg(img *image.RGBA) *image.RGBA {
	return RotateBy180Deg(RotateBy90Deg(img))
}
