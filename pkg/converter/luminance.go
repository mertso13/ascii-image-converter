package converter

// Luminance returns the perceived brightness of an RGB pixel on a 0–255 scale,
// using the ITU-R BT.709 weights:
//
//	Y = 0.2126*R + 0.7152*G + 0.0722*B
func Luminance(r, g, b uint8) uint8 {
	rFloat, gFloat, bFloat := float64(r), float64(g), float64(b)
	Y := (0.2126 * rFloat) + (0.7152 * gFloat) + (0.0722 * bFloat)
	YInt := uint8(Y)
	return YInt
}
