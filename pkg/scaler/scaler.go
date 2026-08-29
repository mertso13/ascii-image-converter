package scaler

import (
	"image"
	"math"

	"golang.org/x/image/draw"
)

const DefaultAspectFactor = 0.5

func CalculateDimensions(origWidth, origHeight, targetWidth, targetHeight int, aspectFactor float64) (int, int) {
	if aspectFactor <= 0 {
		aspectFactor = DefaultAspectFactor
	}

	if origWidth <= 0 || origHeight <= 0 {
		return 0, 0
	}

	if targetWidth > 0 && targetHeight > 0 {
		return targetWidth, targetHeight
	}

	calculatedWidth, calculatedHeight := 1, 1

	if targetWidth > 0 && targetHeight == 0 {
		calculatedHeight = int(math.Round(float64(targetWidth) * (float64(origHeight) / float64(origWidth)) * aspectFactor))
		calculatedWidth = targetWidth
	}

	if targetWidth == 0 && targetHeight > 0 {
		calculatedWidth = int(math.Round(float64(targetHeight) * (float64(origWidth) / float64(origHeight)) / aspectFactor))
		calculatedHeight = targetHeight
	}

	if targetWidth == 0 && targetHeight == 0 {
		calculatedHeight = int(math.Round(float64(origHeight) * aspectFactor))
		calculatedWidth = origWidth
	}

	if calculatedWidth < 1 {
		calculatedWidth = 1
	}

	if calculatedHeight < 1 {
		calculatedHeight = 1
	}

	return calculatedWidth, calculatedHeight
}

func FilterByName(name string) draw.Interpolator {
	switch name {
	case "nearest":
		return draw.NearestNeighbor
	case "bilinear":
		return draw.BiLinear
	case "catmullrom", "bicubic":
		return draw.CatmullRom
	default:
		return draw.BiLinear
	}
}

func Scale(img image.Image, targetWidth, targetHeight int, interpolator draw.Interpolator) image.Image {
	if img == nil || targetWidth <= 0 || targetHeight <= 0 {
		return nil
	}

	if interpolator == nil {
		interpolator = draw.BiLinear
	}

	destinationRectangle := image.Rect(0, 0, targetWidth, targetHeight)
	destinationImage := image.NewRGBA(destinationRectangle)
	interpolator.Scale(destinationImage, destinationImage.Bounds(), img, img.Bounds(), draw.Over, nil)

	return destinationImage
}
