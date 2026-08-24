package converter

// Luminance returns the perceived brightness of an RGB pixel on a 0–255 scale,
// using the ITU-R BT.709 weights:
//
//	Y = 0.2126*R + 0.7152*G + 0.0722*B
func Luminance(r, g, b uint8) uint8 {
	r_float, g_float, b_float := float64(r), float64(g), float64(b)
	Y := (0.2126 * r_float) + (0.7152 * g_float) + (0.0722 * b_float)
	Y_int := uint8(Y)
	return Y_int
}
