package main

import "testing"

func TestParseArgs(t *testing.T) {
	defaultOpts := Options{
		Path:    "img.png",
		Width:   0,
		Height:  0,
		Palette: "standard",
		Scale:   0.5,
		Filter:  "bilinear",
	}

	tests := []struct {
		name    string
		args    []string
		want    Options
		wantErr bool
	}{
		{
			name:    "default flags",
			args:    []string{"img.png"},
			want:    defaultOpts,
			wantErr: false,
		},
		{
			name: "short flags",
			args: []string{"-w", "80", "-H", "40", "-p", "blocks", "-i", "-s", "0.6", "-f", "nearest", "-o", "out.txt", "img.png"},
			want: Options{
				Path:    "img.png",
				Width:   80,
				Height:  40,
				Palette: "blocks",
				Invert:  true,
				Scale:   0.6,
				Filter:  "nearest",
				Output:  "out.txt",
			},
			wantErr: false,
		},
		{
			name: "long flags",
			args: []string{"--width", "100", "--height", "50", "--palette", "extended", "--ramp", " .#", "--invert", "--scale", "0.45", "--filter", "bicubic", "--output", "art.txt", "img.png"},
			want: Options{
				Path:       "img.png",
				Width:      100,
				Height:     50,
				Palette:    "extended",
				CustomRamp: " .#",
				Invert:     true,
				Scale:      0.45,
				Filter:     "bicubic",
				Output:     "art.txt",
			},
			wantErr: false,
		},
		{
			name:    "missing path",
			args:    []string{"-w", "80"},
			wantErr: true,
		},
		{
			name:    "extra positional arguments",
			args:    []string{"a.png", "b.png"},
			wantErr: true,
		},
		{
			name:    "negative width",
			args:    []string{"-w", "-5", "img.png"},
			wantErr: true,
		},
		{
			name:    "negative height",
			args:    []string{"-H", "-10", "img.png"},
			wantErr: true,
		},
		{
			name:    "negative scale",
			args:    []string{"-s", "-0.5", "img.png"},
			wantErr: true,
		},
		{
			name:    "zero scale",
			args:    []string{"-s", "0", "img.png"},
			wantErr: true,
		},
		{
			name:    "bad width value",
			args:    []string{"-w", "abc", "img.png"},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseArgs(tt.args)
			if (err != nil) != tt.wantErr {
				t.Fatalf("ParseArgs() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("ParseArgs() = %+v, want %+v", got, tt.want)
			}
			if tt.wantErr && (err == nil || err.Error() == "") {
				t.Errorf("ParseArgs() expected non-empty error message, got: %v", err)
			}
		})
	}
}
