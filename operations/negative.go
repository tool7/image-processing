package operations

import (
	"image"
	"image/color"

	utils "tool7/image-processing/utils"
)

type NegativeOperation struct{}

func NewNegativeOperation() *NegativeOperation {
	return &NegativeOperation{}
}

func (this *NegativeOperation) Execute(inputImage *image.RGBA) (*image.RGBA, error) {
	worker := func(bounds image.Rectangle) *image.RGBA {
		chunkResult := image.NewRGBA(bounds)

		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				R, G, B, A := utils.GetPixelColor(inputImage, x, y)

				newR := 255 - R
				newG := 255 - G
				newB := 255 - B

				chunkResult.SetRGBA(x, y, color.RGBA{newR, newG, newB, A})
			}
		}
		return chunkResult
	}

	return utils.ProcessImageConcurrently(*inputImage, worker)
}
