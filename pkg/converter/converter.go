package converter

import (
	"image"
	"strings"
)

// Convert returns an ASCII representation of the image using the given character ramp.
func Convert(img image.Image, ramp *Ramp) string {
	if img == nil || ramp == nil {
		return ""
	}

	bounds := img.Bounds()

	var textBuffer strings.Builder

	for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
		for x := bounds.Min.X; x < bounds.Max.X; x++ {
			r, g, b, _ := img.At(x, y).RGBA()

			r8 := uint8(r >> 8)
			g8 := uint8(g >> 8)

			b8 := uint8(b >> 8)

			pixelBrightness := Luminance(r8, g8, b8)
			mappedChar := ramp.At(pixelBrightness)

			textBuffer.WriteRune(mappedChar)
		}
		textBuffer.WriteByte('\n')
	}
	return textBuffer.String()
}
