package scaler

import (
	"image"
	"image/color"
	"testing"

	"golang.org/x/image/draw"
)

func TestCalculateDimensions(t *testing.T) {
	tests := []struct {
		name         string
		origWidth    int
		origHeight   int
		targetWidth  int
		targetHeight int
		aspectFactor float64
		wantWidth    int
		wantHeight   int
	}{
		{
			name:         "width specified, compute height with default factor",
			origWidth:    800,
			origHeight:   600,
			targetWidth:  80,
			targetHeight: 0,
			aspectFactor: 0.5,
			wantWidth:    80,
			wantHeight:   30,
		},
		{
			name:         "height specified, compute width with default factor",
			origWidth:    800,
			origHeight:   600,
			targetWidth:  0,
			targetHeight: 30,
			aspectFactor: 0.5,
			wantWidth:    80,
			wantHeight:   30,
		},
		{
			name:         "both width and height specified",
			origWidth:    800,
			origHeight:   600,
			targetWidth:  100,
			targetHeight: 50,
			aspectFactor: 0.5,
			wantWidth:    100,
			wantHeight:   50,
		},
		{
			name:         "neither specified, scale original height by aspect factor",
			origWidth:    100,
			origHeight:   100,
			targetWidth:  0,
			targetHeight: 0,
			aspectFactor: 0.5,
			wantWidth:    100,
			wantHeight:   50,
		},
		{
			name:         "zero aspect factor defaults to 0.5",
			origWidth:    800,
			origHeight:   600,
			targetWidth:  80,
			targetHeight: 0,
			aspectFactor: 0.0,
			wantWidth:    80,
			wantHeight:   30,
		},
		{
			name:         "invalid original dimensions return zero",
			origWidth:    0,
			origHeight:   600,
			targetWidth:  80,
			targetHeight: 0,
			aspectFactor: 0.5,
			wantWidth:    0,
			wantHeight:   0,
		},
		{
			name:         "minimum scaled dimension is at least 1",
			origWidth:    1000,
			origHeight:   1,
			targetWidth:  10,
			targetHeight: 0,
			aspectFactor: 0.5,
			wantWidth:    10,
			wantHeight:   1,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			gotWidth, gotHeight := CalculateDimensions(
				tc.origWidth,
				tc.origHeight,
				tc.targetWidth,
				tc.targetHeight,
				tc.aspectFactor,
			)

			if gotWidth != tc.wantWidth || gotHeight != tc.wantHeight {
				t.Errorf(
					"CalculateDimensions(%d, %d, %d, %d, %f) = (%d, %d), want (%d, %d)",
					tc.origWidth,
					tc.origHeight,
					tc.targetWidth,
					tc.targetHeight,
					tc.aspectFactor,
					gotWidth,
					gotHeight,
					tc.wantWidth,
					tc.wantHeight,
				)
			}
		})
	}
}

func TestFilterByName(t *testing.T) {
	tests := []struct {
		name       string
		filterName string
		want       draw.Interpolator
	}{
		{
			name:       "nearest filter",
			filterName: "nearest",
			want:       draw.NearestNeighbor,
		},
		{
			name:       "bilinear filter",
			filterName: "bilinear",
			want:       draw.BiLinear,
		},
		{
			name:       "catmullrom filter",
			filterName: "catmullrom",
			want:       draw.CatmullRom,
		},
		{
			name:       "bicubic alias",
			filterName: "bicubic",
			want:       draw.CatmullRom,
		},
		{
			name:       "unknown filter defaults to bilinear",
			filterName: "unknown",
			want:       draw.BiLinear,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := FilterByName(tc.filterName)
			if got != tc.want {
				t.Errorf("FilterByName(%q) = %v, want %v", tc.filterName, got, tc.want)
			}
		})
	}
}

func TestScale(t *testing.T) {
	srcImg := image.NewRGBA(image.Rect(0, 0, 100, 100))
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			srcImg.Set(x, y, color.RGBA{R: 255, G: 0, B: 0, A: 255})
		}
	}

	tests := []struct {
		name         string
		img          image.Image
		targetWidth  int
		targetHeight int
		filter       draw.Interpolator
		wantNil      bool
		wantWidth    int
		wantHeight   int
	}{
		{
			name:         "scale down valid image",
			img:          srcImg,
			targetWidth:  50,
			targetHeight: 25,
			filter:       draw.BiLinear,
			wantNil:      false,
			wantWidth:    50,
			wantHeight:   25,
		},
		{
			name:         "scale with nil filter defaults to bilinear",
			img:          srcImg,
			targetWidth:  40,
			targetHeight: 20,
			filter:       nil,
			wantNil:      false,
			wantWidth:    40,
			wantHeight:   20,
		},
		{
			name:         "nil image returns nil",
			img:          nil,
			targetWidth:  50,
			targetHeight: 25,
			filter:       draw.BiLinear,
			wantNil:      true,
		},
		{
			name:         "invalid width returns nil",
			img:          srcImg,
			targetWidth:  0,
			targetHeight: 25,
			filter:       draw.BiLinear,
			wantNil:      true,
		},
		{
			name:         "invalid height returns nil",
			img:          srcImg,
			targetWidth:  50,
			targetHeight: -1,
			filter:       draw.BiLinear,
			wantNil:      true,
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got := Scale(tc.img, tc.targetWidth, tc.targetHeight, tc.filter)
			if tc.wantNil {
				if got != nil {
					t.Fatalf("Scale(...) = %v, want nil", got)
				}
				return
			}

			if got == nil {
				t.Fatalf("Scale(...) = nil, want image")
			}

			bounds := got.Bounds()
			if bounds.Dx() != tc.wantWidth || bounds.Dy() != tc.wantHeight {
				t.Errorf("Scale bounds = (%d, %d), want (%d, %d)", bounds.Dx(), bounds.Dy(), tc.wantWidth, tc.wantHeight)
			}

			r, g, b, _ := got.At(0, 0).RGBA()
			if uint8(r>>8) != 255 || uint8(g>>8) != 0 || uint8(b>>8) != 0 {
				t.Errorf("Scale pixel color = (%d, %d, %d), want (255, 0, 0)", r>>8, g>>8, b>>8)
			}
		})
	}
}
