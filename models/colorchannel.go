package models

type ColorChannel interface {
	~uint8 | ~uint16 | ~uint32 | ~uint64 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64
}
