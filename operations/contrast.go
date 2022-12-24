package operations

import (
	"image"
	"image/color"

	utils "tool7/image-processing/utils"
)

type ContrastOperation struct {
	factor float64
}

func NewContrastOperation(factor float64) *ContrastOperation {
	return &ContrastOperation{
		factor,
	}
}

func (this *ContrastOperation) Execute(inputImage *image.RGBA) (*image.RGBA, error) {
	worker := func(bounds image.Rectangle) *image.RGBA {
		chunkResult := image.NewRGBA(bounds)

		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				R, G, B, A := utils.GetPixelColor(inputImage, x, y)

				newR := this.factor*(float64(R)-128) + 128
				newG := this.factor*(float64(G)-128) + 128
				newB := this.factor*(float64(B)-128) + 128

				clippedR := utils.ClipColorChannel(newR)
				clippedG := utils.ClipColorChannel(newG)
				clippedB := utils.ClipColorChannel(newB)

				chunkResult.SetRGBA(x, y, color.RGBA{clippedR, clippedG, clippedB, A})
			}
		}
		return chunkResult
	}

	return utils.ProcessImageConcurrently(*inputImage, worker)
}
