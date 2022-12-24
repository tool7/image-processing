package models

type KernelType int

const (
	Blur KernelType = iota
	MotionBlur
	Sharpen
	EdgeDetectionHorizontal
	EdgeDetectionVertical
	Emboss
)
