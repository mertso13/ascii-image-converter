package converter

import "testing"

func TestNewRamp(t *testing.T) {
	emptyRamp := NewRamp("")
	if emptyRamp != nil {
		t.Errorf("expected nil, got %v", emptyRamp)
	}
	validRamp := NewRamp("abc")
	if validRamp == nil || validRamp.Len() != 3 {
		t.Fatalf("expected length 3, got %v", validRamp)
	}
}

func TestRampAt(t *testing.T) {
	ramp := NewRamp("abc")
	tests := []struct {
		name     string
		lum      uint8
		expected rune
	}{
		{name: "a", lum: 0, expected: 'a'},
		{name: "b", lum: 128, expected: 'b'},
		{name: "c", lum: 255, expected: 'c'},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := ramp.At(tt.lum)
			if actual != tt.expected {
				t.Errorf("ramp.At(%d) = %c; want %c", tt.lum, actual, tt.expected)
			}
		})
	}
}

func TestRampInvert(t *testing.T) {
	invertedRamp := NewRamp("abc").Invert()
	if invertedRamp.At(0) != 'c' || invertedRamp.At(255) != 'a' {
		t.Fatalf("expected 'c' at 0 'a' at 255, got %v", invertedRamp)
	}
}

func TestByName(t *testing.T) {
	tests := []struct {
		name      string
		rampName  string
		wantFound bool
	}{
		{name: "standard", rampName: "standard", wantFound: true},
		{name: "extended", rampName: "extended", wantFound: true},
		{name: "minimal", rampName: "minimal", wantFound: true},
		{name: "blocks", rampName: "blocks", wantFound: true},
		{name: "invalid", rampName: "invalid", wantFound: false},
		{name: "unknown", rampName: "", wantFound: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ramp, ok := ByName(tt.rampName)
			if ok != tt.wantFound {
				t.Fatalf("ByName(%q) ok = %v; want %v", tt.rampName, ok, tt.wantFound)
			}
			if ok && ramp == nil {
				t.Fatalf("ByName(%q) returned nil ramp with ok=true", tt.rampName)
			}
		})
	}
}
