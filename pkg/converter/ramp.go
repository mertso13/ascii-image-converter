package converter

import (
	"math"
)

// Ramp maps pixel brightness to characters. Its characters must be ordered
// darkest first.
type Ramp struct {
	chars []rune
}

// NewRamp returns a Ramp using chars, ordered darkest to lightest,
// or nil if chars is empty.
func NewRamp(chars string) *Ramp {
	if len(chars) == 0 {
		return nil
	}

	ramp := Ramp{
		chars: []rune(chars),
	}

	return &ramp
}

// Len returns the number of characters in the ramp.
func (r *Ramp) Len() int {
	return len(r.chars)
}

// At returns the character for luminance lum. Black maps to the first
// character, white to the last.
func (r *Ramp) At(lum uint8) rune {
	lumFloat := float64(lum)
	index := math.Floor(lumFloat/255*(float64(r.Len()-1)) + 0.5)
	indexInt := int(index)
	return r.chars[indexInt]
}

// Invert returns a copy of the ramp with reversed order, for use on
// light terminal backgrounds.
func (r *Ramp) Invert() *Ramp {
	var rReversed []rune
	for i := len(r.chars) - 1; i >= 0; i-- {
		rReversed = append(rReversed, r.chars[i])
	}

	r2 := Ramp{
		chars: rReversed,
	}

	return &r2
}

const (
	rampStandard = " .:-=+*#%@"
	rampExtended = "$@B%8&WM#*oahkbdpqwmZO0QLCJUYXzcvunxrjft/\\|()1{}[]?-_+~<>i!lI;:,\"^`'. "
	rampMinimal  = " .oO"
	rampBlocks   = " ░▒▓█"
)

// ByName returns the named built-in ramp: "standard", "extended",
// "minimal", or "blocks". ok is false if name is unknown.
func ByName(name string) (r *Ramp, ok bool) {
	switch name {
	case "standard":
		return NewRamp(rampStandard), true
	case "extended":
		return NewRamp(rampExtended), true
	case "minimal":
		return NewRamp(rampMinimal), true
	case "blocks":
		return NewRamp(rampBlocks), true
	default:
		return nil, false
	}
}
