package operations

import (
	"image"
	"image/color"

	utils "tool7/image-processing/utils"
)

type TintOperation struct {
	Tint      color.RGBA
	Intensity float64
}

func NewTintOperation(tint color.RGBA, intensity float64) *TintOperation {
	return &TintOperation{
		tint,
		intensity,
	}
}

func (this *TintOperation) Execute(inputImage *image.RGBA) (*image.RGBA, error) {
	worker := func(bounds image.Rectangle) *image.RGBA {
		chunkResult := image.NewRGBA(bounds)

		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				R, G, B, A := utils.GetPixelColor(inputImage, x, y)

				tintR := float64(this.Tint.R) * this.Intensity
				tintG := float64(this.Tint.G) * this.Intensity
				tintB := float64(this.Tint.B) * this.Intensity

				newR := utils.ClipColorChannel(int32(R) + int32(tintR))
				newG := utils.ClipColorChannel(int32(G) + int32(tintG))
				newB := utils.ClipColorChannel(int32(B) + int32(tintB))

				chunkResult.SetRGBA(x, y, color.RGBA{newR, newG, newB, A})
			}
		}
		return chunkResult
	}

	return utils.ProcessImageConcurrently(*inputImage, worker)
}
