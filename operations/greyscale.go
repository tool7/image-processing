package operations

import (
	"image"
	"image/color"

	utils "tool7/image-processing/utils"
)

type GreyscaleOperation struct{}

func NewGreyscaleOperation() *GreyscaleOperation {
	return &GreyscaleOperation{}
}

// This is one of the standard formulas for calculating "grey value" of the pixel
func getPixelGreyValue(r, g, b uint8) uint8 {
	return uint8(0.299*float64(r) + 0.587*float64(g) + 0.114*float64(b))
}

func (this *GreyscaleOperation) Execute(inputImage *image.RGBA) (*image.RGBA, error) {
	worker := func(bounds image.Rectangle) *image.RGBA {
		chunkResult := image.NewRGBA(bounds)

		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				R, G, B, A := utils.GetPixelColor(inputImage, x, y)

				grey := getPixelGreyValue(R, G, B)
				chunkResult.SetRGBA(x, y, color.RGBA{grey, grey, grey, A})
			}
		}
		return chunkResult
	}

	return utils.ProcessImageConcurrently(*inputImage, worker)
}
