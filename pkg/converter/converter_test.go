package converter

import (
	"image"
	"image/color"
	"testing"
)

func TestConvert(t *testing.T) {
	img := image.NewRGBA(image.Rect(0, 0, 2, 2))
	img.Set(0, 0, color.RGBA{R: 0, G: 0, B: 0, A: 255})
	img.Set(1, 0, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	img.Set(0, 1, color.RGBA{R: 255, G: 255, B: 255, A: 255})
	img.Set(1, 1, color.RGBA{R: 0, G: 0, B: 0, A: 255})

	ramp, ok := ByName("standard")
	if !ok {
		t.Fatal("failed to get standard ramp")
	}

	result := Convert(img, ramp)
	expected := " @\n@ \n"

	if result != expected {
		t.Fatalf("expected %q, got %q", expected, result)
	}
}

func TestConvertNil(t *testing.T) {
	ramp, _ := ByName("standard")
	if result := Convert(nil, ramp); result != "" {
		t.Fatalf("expected empty string for nil image, got %q", result)
	}

	img := image.NewRGBA(image.Rect(0, 0, 1, 1))
	if result := Convert(img, nil); result != "" {
		t.Fatalf("expected empty string for nil ramp, got %q", result)
	}
}
