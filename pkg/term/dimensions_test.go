package term

import (
	"testing"
)

func TestDefaultConstants(t *testing.T) {
	if DefaultWidth != 80 {
		t.Errorf("expected DefaultWidth to be 80, got %d", DefaultWidth)
	}

	if DefaultHeight != 24 {
		t.Errorf("expected DefaultHeight to be 24, got %d", DefaultHeight)
	}
}

func TestGetSizeInvalidDescriptor(t *testing.T) {
	_, _, err := GetSize(-1)
	if err == nil {
		t.Error("expected error for invalid file descriptor, got nil")
	}
}

func TestGetSizeOrDefaultInvalidDescriptor(t *testing.T) {
	width, height := GetSizeOrDefault(-1)

	if width != DefaultWidth {
		t.Errorf("expected width %d, got %d", DefaultWidth, width)
	}

	if height != DefaultHeight {
		t.Errorf("expected height %d, got %d", DefaultHeight, height)
	}
}
