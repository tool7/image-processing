package utils

import (
	"errors"
	"image"
	"image/color"
	"math"
	"os"
	"runtime"
	"sync"
	"time"

	models "tool7/image-processing/models"
)

func NewImageLayer(operation models.ImageOperation) *models.ImageLayer {
	return &models.ImageLayer{
		Operation: operation,
		IsEnabled: true,
	}
}

func NewImageLayerCollection(img *image.RGBA) *models.ImageLayerCollection {
	return &models.ImageLayerCollection{
		InputImage:   img,
		OutputImages: make(map[int]*image.RGBA),
		Head:         nil,
		Size:         0,
	}
}

func GetKernelByType(kernelType models.KernelType) ([][]float32, error) {
	switch kernelType {
	case models.Blur:
		return [][]float32{
			{1.0 / 16.0, 2.0 / 16.0, 1.0 / 16.0},
			{2.0 / 16.0, 4.0 / 16.0, 2.0 / 16.0},
			{1.0 / 16.0, 2.0 / 16.0, 1.0 / 16.0},
		}, nil
	case models.MotionBlur:
		return [][]float32{
			{1.0 / 9.0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 1.0 / 9.0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 1.0 / 9.0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 1.0 / 9.0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 1.0 / 9.0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 1.0 / 9.0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 1.0 / 9.0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 1.0 / 9.0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 1.0 / 9.0},
		}, nil
	case models.Sharpen:
		return [][]float32{
			{+0, -1, +0},
			{-1, +5, -1},
			{+0, -1, +0},
		}, nil
	case models.EdgeDetectionHorizontal:
		return [][]float32{
			{+1, +0, -1},
			{+2, +0, -2},
			{+1, +0, -1},
		}, nil
	case models.EdgeDetectionVertical:
		return [][]float32{
			{+1, +2, +1},
			{+0, +0, +0},
			{-1, -2, -1},
		}, nil
	case models.Emboss:
		return [][]float32{
			{-2, +0, -1, +0, +0},
			{+0, -2, -1, +0, +0},
			{-1, -1, +1, +1, +1},
			{+0, +0, +1, +2, +0},
			{+0, +0, +1, +0, +2},
		}, nil
	}

	return nil, errors.New("No kernel of provided type")
}

func GetPixelColor(img image.Image, x, y int) (uint8, uint8, uint8, uint8) {
	r32, g32, b32, a32 := img.At(x, y).RGBA()
	r := uint8(r32 >> 8)
	g := uint8(g32 >> 8)
	b := uint8(b32 >> 8)
	a := uint8(a32 >> 8)
	return r, g, b, a
}

func ClipColorChannel[T models.ColorChannel](channel T) uint8 {
	if channel < 0 {
		return 0
	}
	if channel > 255 {
		return 255
	}
	return uint8(channel)
}

func GetImageFromFilePath(filePath string) (*image.RGBA, error) {
	f, err := os.Open(filePath)
	defer f.Close()

	if err != nil {
		return nil, err
	}

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	rgbaImage := image.NewRGBA(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			r, g, b, a := GetPixelColor(img, x, y)
			rgbaImage.SetRGBA(x, y, color.RGBA{r, g, b, a})
		}
	}

	return rgbaImage, nil
}

func ProcessImageConcurrently(img image.RGBA, worker func(image.Rectangle) *image.RGBA) (*image.RGBA, error) {
	numberOfThreads := runtime.NumCPU()

	var wg sync.WaitGroup
	bounds := splitImageToBounds(img, numberOfThreads)
	resultChunks := make([]*image.RGBA, 0, numberOfThreads)
	resultChannel := make(chan *image.RGBA, numberOfThreads)

	for _, imageBound := range bounds {
		wg.Add(1)

		imageBound := imageBound

		go func() {
			defer wg.Done()
			resultChannel <- worker(imageBound)
		}()
	}

	wg.Wait()

	for i := 0; i < len(bounds); i++ {
		select {
		case chunk := <-resultChannel:
			resultChunks = append(resultChunks, chunk)
		case <-time.After(time.Second * 5):
			return nil, errors.New("Timeout exceeded (5s)")
		}
	}

	result := image.NewRGBA(img.Bounds())

	for _, chunk := range resultChunks {
		minX := chunk.Bounds().Min.X
		minY := chunk.Bounds().Min.Y
		maxX := chunk.Bounds().Max.X
		maxY := chunk.Bounds().Max.Y

		for y := minY; y < maxY; y++ {
			for x := minX; x < maxX; x++ {
				r, g, b, a := GetPixelColor(chunk, x, y)
				result.SetRGBA(x, y, color.RGBA{r, g, b, a})
			}
		}
	}

	return result, nil
}

func splitImageToBounds(img image.RGBA, desiredCount int) []image.Rectangle {
	chunks := make([]image.Rectangle, 0, desiredCount)

	minX := img.Bounds().Min.X
	minY := img.Bounds().Min.Y
	maxX := img.Bounds().Max.X
	maxY := img.Bounds().Max.Y
	chunkHeight := int(
		math.Ceil(float64(maxY-minY) / float64(desiredCount)),
	)

	for y := minY; y < maxY; y += chunkHeight {
		end := y + chunkHeight

		if end > maxY {
			end = maxY
		}

		chunk := image.Rect(minX, y, maxX, end)
		chunks = append(chunks, chunk)
	}

	return chunks
}
