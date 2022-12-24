package operations

import (
	"image"
	"image/color"

	utils "tool7/image-processing/utils"
)

type BrightnessOperation struct {
	level float64
}

func NewBrightnessOperation(level float64) *BrightnessOperation {
	return &BrightnessOperation{
		level,
	}
}

func (this *BrightnessOperation) Execute(inputImage *image.RGBA) (*image.RGBA, error) {
	worker := func(bounds image.Rectangle) *image.RGBA {
		chunkResult := image.NewRGBA(bounds)

		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				R, G, B, A := utils.GetPixelColor(inputImage, x, y)

				newR := utils.ClipColorChannel(float64(R) * this.level)
				newG := utils.ClipColorChannel(float64(G) * this.level)
				newB := utils.ClipColorChannel(float64(B) * this.level)

				chunkResult.SetRGBA(x, y, color.RGBA{newR, newG, newB, A})
			}
		}
		return chunkResult
	}

	return utils.ProcessImageConcurrently(*inputImage, worker)
}
