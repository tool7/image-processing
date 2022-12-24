package operations

import (
	"image"
	"image/color"

	utils "tool7/image-processing/utils"
)

type SepiaOperation struct{}

func NewSepiaOperation() *SepiaOperation {
	return &SepiaOperation{}
}

// This is one of the standard formulas for calculating "sepia value" of the pixel
func getPixelSepiaValue(r, g, b uint8) (uint8, uint8, uint8) {
	newR := (float32(r) * 0.393) + (float32(g) * 0.769) + (float32(b) * 0.189)
	newG := (float32(r) * 0.349) + (float32(g) * 0.686) + (float32(b) * 0.168)
	newB := (float32(r) * 0.272) + (float32(g) * 0.534) + (float32(b) * 0.131)

	clippedR := utils.ClipColorChannel(newR)
	clippedG := utils.ClipColorChannel(newG)
	clippedB := utils.ClipColorChannel(newB)

	return clippedR, clippedG, clippedB
}

func (this *SepiaOperation) Execute(inputImage *image.RGBA) (*image.RGBA, error) {
	worker := func(bounds image.Rectangle) *image.RGBA {
		chunkResult := image.NewRGBA(bounds)

		for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
			for x := bounds.Min.X; x < bounds.Max.X; x++ {
				R, G, B, A := utils.GetPixelColor(inputImage, x, y)

				newR, newG, newB := getPixelSepiaValue(R, G, B)
				chunkResult.SetRGBA(x, y, color.RGBA{newR, newG, newB, A})
			}
		}
		return chunkResult
	}

	return utils.ProcessImageConcurrently(*inputImage, worker)
}
