package models

import "image"

type ImageOperation interface {
	Execute(*image.RGBA) (*image.RGBA, error)
}
