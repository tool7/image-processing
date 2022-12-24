package operations

import (
	"image"
	"image/color"

	colorful "github.com/lucasb-eyer/go-colorful"

	utils "tool7/image-processing/utils"
)

type TintOperation struct {
	tint      color.RGBA
	intensity float64
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

				newR := utils.ClipColorChannel(int32(R) + int32(this.tint.R))
				newG := utils.ClipColorChannel(int32(G) + int32(this.tint.G))
				newB := utils.ClipColorChannel(int32(B) + int32(this.tint.B))
				newA := utils.ClipColorChannel(int32(A) + int32(this.tint.A))

				hue, normalizedSaturation, value := colorful.Color{
					R: float64(newR),
					G: float64(newG),
					B: float64(newB),
				}.Hsv()
				saturation := utils.ClipColorChannel(normalizedSaturation * 255)

				increasedSaturation := utils.ClipColorChannel(float64(saturation) * this.intensity)
				normalizedIncreasedSaturation := float64(increasedSaturation) / 255.0

				newColor := colorful.Hsv(hue, normalizedIncreasedSaturation, value)
				chunkResult.SetRGBA(x, y, color.RGBA{
					uint8(newColor.R),
					uint8(newColor.G),
					uint8(newColor.B),
					newA,
				})
			}
		}
		return chunkResult
	}

	return utils.ProcessImageConcurrently(*inputImage, worker)
}
