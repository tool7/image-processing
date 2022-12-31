package models

import (
	"errors"
	"image"
)

// Linked list data structure for handling image layers
type ImageLayerCollection struct {
	InputImage   *image.RGBA
	OutputImages map[int]*image.RGBA
	Head         *ImageLayer
	Size         int
}

func (this *ImageLayerCollection) validateIndex(index int) bool {
	if index < 0 || index >= this.Size {
		return false
	}
	return true
}

func (this *ImageLayerCollection) At(index int) (*ImageLayer, error) {
	ok := this.validateIndex(index)
	if !ok {
		return nil, errors.New("Invalid index")
	}

	current := this.Head
	for count := 0; count < index; count++ {
		current = current.Next
	}

	return current, nil
}

func (this *ImageLayerCollection) Append(imageLayer *ImageLayer) {
	if this.Head == nil {
		this.Head = imageLayer
		this.Size++
		return
	}

	current := this.Head

	for current.Next != nil {
		current = current.Next
	}

	current.Next = imageLayer
	this.Size++
}

func (this *ImageLayerCollection) InsertAt(imageLayer *ImageLayer, index int) error {
	if this.Size == index {
		this.Append(imageLayer)
		return nil
	}

	ok := this.validateIndex(index)
	if !ok {
		return errors.New("Invalid index")
	}

	if index == 0 {
		oldHead := this.Head
		this.Head = imageLayer
		this.Head.Next = oldHead

		this.Size++
		return nil
	}

	var previous *ImageLayer
	current := this.Head

	for count := 0; count < index; count++ {
		previous = current
		current = current.Next
	}

	imageLayer.Next = current
	previous.Next = imageLayer

	this.Size++
	return nil
}

func (this *ImageLayerCollection) RemoveAt(index int) error {
	ok := this.validateIndex(index)
	if !ok {
		return errors.New("Invalid index")
	}

	var previous *ImageLayer
	current := this.Head

	if index == 0 {
		this.Head = current.Next
	} else {
		for count := 0; count < index; count++ {
			previous = current
			current = current.Next
		}

		previous.Next = current.Next
	}

	for index < this.Size {
		delete(this.OutputImages, index)
		index++
	}

	current.Next = nil
	this.Size--
	return nil
}

func (this *ImageLayerCollection) ExecuteLayersFrom(index int) (*image.RGBA, error) {
	ok := this.validateIndex(index)
	if !ok {
		return nil, errors.New("Invalid index")
	}

	current := this.Head

	for count := 0; count < index; count++ {
		current = current.Next
	}

	var currentLayerInputImage *image.RGBA

	if index > 0 && this.OutputImages[index-1] != nil {
		currentLayerInputImage = this.OutputImages[index-1]
	} else {
		currentLayerInputImage = this.InputImage
	}

	for current != nil {
		outputImage := current.ExecuteOperation(currentLayerInputImage)

		currentLayerInputImage = outputImage
		this.OutputImages[index] = outputImage

		current = current.Next
		index++
	}

	return currentLayerInputImage, nil
}
