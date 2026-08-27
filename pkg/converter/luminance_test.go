package converter

import "testing"

func TestLuminance(t *testing.T) {
	tests := []struct {
		name     string
		r        uint8
		g        uint8
		b        uint8
		expected uint8
	}{
		{name: "black", r: 0, g: 0, b: 0, expected: 0},
		{name: "white", r: 255, g: 255, b: 255, expected: 255},
		{name: "pure red", r: 255, g: 0, b: 0, expected: 54},
		{name: "pure green", r: 0, g: 255, b: 0, expected: 182},
		{name: "pure blue", r: 0, g: 0, b: 255, expected: 18},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Luminance(tt.r, tt.g, tt.b)
			if actual != tt.expected {
				t.Errorf("Luminance(%d, %d, %d) = %d; want %d", tt.r, tt.g, tt.b, actual, tt.expected)
			}
		})
	}
}
