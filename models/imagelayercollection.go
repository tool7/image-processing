package models

import (
	"errors"
	"image"
	"math"
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

func (this *ImageLayerCollection) Swap(indexA, indexB int) error {
	if indexA == indexB {
		return errors.New("Index A and B must be different")
	}
	ok := this.validateIndex(indexA)
	if !ok {
		return errors.New("Invalid index A")
	}
	ok = this.validateIndex(indexB)
	if !ok {
		return errors.New("Invalid index B")
	}

	var pointerToLayerA *ImageLayer
	layerA := this.Head
	for count := 0; count < indexA; count++ {
		pointerToLayerA = layerA
		layerA = layerA.Next
	}

	var pointerToLayerB *ImageLayer
	layerB := this.Head
	for count := 0; count < indexB; count++ {
		pointerToLayerB = layerB
		layerB = layerB.Next
	}

	// If "pointerToLayerA" is nil, it means that "indexA" is 0,
	// which implies that "layerA" is HEAD so HEAD must be set to "layerB".
	if pointerToLayerA == nil {
		this.Head = layerB
	} else {
		pointerToLayerA.Next = layerB
	}
	// If "pointerToLayerB" is nil, it means that "indexB" is 0,
	// which implies that "layerB" is HEAD so HEAD must be set to "layerA".
	if pointerToLayerB == nil {
		this.Head = layerA
	} else {
		pointerToLayerB.Next = layerA
	}

	layerATraget := layerA.Next
	layerBTraget := layerB.Next
	layerA.Next = layerBTraget
	layerB.Next = layerATraget

	// Updating "OutputImages" map
	minIndex := int(math.Min(float64(indexA), float64(indexB)))
	for minIndex < this.Size {
		delete(this.OutputImages, minIndex)
		minIndex++
	}

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
