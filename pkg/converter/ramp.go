package converter

import (
	"math"
)

type Ramp struct {
	chars []rune
}

func NewRamp(chars string) *Ramp {
	if len(chars) == 0 {
		return nil
	}

	ramp := Ramp{
		chars: []rune(chars),
	}

	return &ramp
}

func (r *Ramp) Len() int {
	return len(r.chars)
}

func (r *Ramp) At(lum uint8) rune {
	lumFloat := float64(lum)
	index := math.Floor(lumFloat/255*(float64(r.Len()-1)) + 0.5)
	indexInt := int(index)
	return r.chars[indexInt]
}

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
