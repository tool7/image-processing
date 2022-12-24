package operations

import (
	"image"
	"image/color"
	utils "tool7/image-processing/utils"

	colorful "github.com/lucasb-eyer/go-colorful"
)

type SaturationOperation struct {
	level float64
}

func NewSaturationOperation(level float64) *SaturationOperation {
	return &SaturationOperation{
		level,
	}
}

func (this *SaturationOperation) Execute(inputImage *image.RGBA) (*image.RGBA, error) {
	worker := func(bounds image.Rectangle) *image.RGBA {
		chunkResult := image.NewRGBA(bounds)

		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				R, G, B, A := utils.GetPixelColor(inputImage, x, y)

				hue, normalizedSaturation, value := colorful.Color{
					R: float64(R),
					G: float64(G),
					B: float64(B)}.Hsv()
				saturation := utils.ClipColorChannel(normalizedSaturation * 255)

				increasedSaturation := utils.ClipColorChannel(float64(saturation) * this.level)
				normalizedIncreasedSaturation := float64(increasedSaturation) / 255.0

				newColor := colorful.Hsv(hue, normalizedIncreasedSaturation, value)
				chunkResult.SetRGBA(x, y, color.RGBA{
					uint8(newColor.R),
					uint8(newColor.G),
					uint8(newColor.B),
					A,
				})
			}
		}
		return chunkResult
	}

	return utils.ProcessImageConcurrently(*inputImage, worker)
}
