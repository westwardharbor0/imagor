package imagor

import (
	"errors"
	"fmt"
	"image"
	"image/color"
	imageGIF "image/gif"
	imageJPEG "image/jpeg"
	imagePNG "image/png"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Imagor struct {
	path          string
	pixelCellSize int
}

func NewImagor(path string, pcs int) *Imagor {
	return &Imagor{
		path:          path,
		pixelCellSize: pcs,
	}
}

// loadImageFile loads the source image for generation.
func (img *Imagor) loadImageFile() (io.Reader, error) {
	imageFile, err := os.Open(img.path)

	if errors.Is(err, os.ErrNotExist) {
		return nil, fmt.Errorf("failed to find image file: %w", err)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to open image file: %w", err)
	}

	return imageFile, nil
}

// GenerateGrid creates a grid of average cell colors from pixels of image.
func (img *Imagor) GenerateGrid() ([][]color.RGBA, error) {
	loadedImage, err := img.loadImage()

	if err != nil {
		return nil, fmt.Errorf("failed to load image data: %w", err)
	}

	maxY := loadedImage.Bounds().Dy() / img.pixelCellSize
	result := make([][]color.RGBA, maxY)

	for y := 1; y <= maxY; y++ {
		maxX := loadedImage.Bounds().Dx() / img.pixelCellSize
		result[y-1] = make([]color.RGBA, maxX)

		for x := 1; x <= maxX; x++ {
			result[y-1][x-1] = img.averageColor(loadedImage, y*img.pixelCellSize, x*img.pixelCellSize)
		}
	}

	return result, nil
}

func (img *Imagor) loadImage() (image.Image, error) {
	var loadedImage image.Image

	extension := filepath.Ext(img.path)
	loadedFile, err := img.loadImageFile()

	if err != nil {
		return nil, fmt.Errorf("faield to load image: %w", err)
	}

	switch strings.ToLower(extension) {
	case ".jpg", ".jpeg":
		loadedImage, err = imageJPEG.Decode(loadedFile)
		if err != nil {
			return nil, fmt.Errorf("failed to open jp(e)g image data: %w", err)
		}

	case ".png":
		loadedImage, err = imagePNG.Decode(loadedFile)
		if err != nil {
			return nil, fmt.Errorf("failed to open png image data: %w", err)
		}

	case ".gif", ".jif":
		loadedImage, err = imageGIF.Decode(loadedFile)
		if err != nil {
			return nil, fmt.Errorf("failed to open gif image data: %w", err)
		}

	default:
		return nil, fmt.Errorf("unsupported file extension: %s", extension)
	}

	return loadedImage, nil
}

// averageColor creates averege cell color from pixels batch.
func (img *Imagor) averageColor(sourceImage image.Image, endY, endX int) color.RGBA {
	averageColor := NewAverageColor()

	for y := endY - img.pixelCellSize; y <= endY; y++ {
		for x := endX - img.pixelCellSize; x <= endX; x++ {
			averageColor.Add(sourceImage.At(x, y).RGBA())
		}
	}

	return averageColor.Color()
}
