package models

import (
	"image"
)

type ImageLayer struct {
	Operation ImageOperation
	IsEnabled bool
	Next      *ImageLayer
}

func (this *ImageLayer) ExecuteOperation(inputImage *image.RGBA) *image.RGBA {
	if !this.IsEnabled {
		return inputImage
	}

	outputImage, _ := this.Operation.Execute(inputImage)
	return outputImage
}

func (this *ImageLayer) Enable() {
	this.IsEnabled = true
}

func (this *ImageLayer) Disable() {
	this.IsEnabled = false
}
