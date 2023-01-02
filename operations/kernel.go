package operations

import (
	"image"
	"image/color"
	"math"

	models "tool7/image-processing/models"
	utils "tool7/image-processing/utils"
)

type KernelOperation struct {
	KernelType models.KernelType
	KernelSize models.KernelSize
}

func NewKernelOperation(kernelType models.KernelType, kernelSize models.KernelSize) *KernelOperation {
	return &KernelOperation{
		kernelType,
		kernelSize,
	}
}

func (this *KernelOperation) Execute(inputImage *image.RGBA) (*image.RGBA, error) {
	kernel := models.GenerateKernel(this.KernelType, this.KernelSize)

	worker := func(bounds image.Rectangle) *image.RGBA {
		return applyKernel(inputImage, bounds, kernel)
	}

	return utils.ProcessImageConcurrently(*inputImage, worker)
}

func applyKernel(inputImage *image.RGBA, bounds image.Rectangle, kernel [][]float32) *image.RGBA {
	result := image.NewRGBA(bounds)

	minY := bounds.Min.Y
	maxY := bounds.Max.Y
	minX := bounds.Min.X
	maxX := bounds.Max.X

	kernelLength := len(kernel)
	kernelCenter := int(math.Floor(float64(kernelLength) / 2))

	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			_, _, _, A := utils.GetPixelColor(inputImage, x, y)

			var sumR float32 = 0
			var sumG float32 = 0
			var sumB float32 = 0

			for j := -kernelCenter; j <= kernelCenter; j++ {
				for i := -kernelCenter; i <= kernelCenter; i++ {
					offsetX := x + i
					offsetY := y + j
					kernelX := i + kernelCenter
					kernelY := j + kernelCenter

					R, G, B, _ := utils.GetPixelColor(inputImage, offsetX, offsetY)

					sumR += kernel[kernelY][kernelX] * float32(R)
					sumG += kernel[kernelY][kernelX] * float32(G)
					sumB += kernel[kernelY][kernelX] * float32(B)
				}
			}

			newR := utils.ClipColorChannel(sumR)
			newG := utils.ClipColorChannel(sumG)
			newB := utils.ClipColorChannel(sumB)

			result.SetRGBA(x, y, color.RGBA{newR, newG, newB, A})
		}
	}

	return result
}
