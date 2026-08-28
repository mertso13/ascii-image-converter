package main

import "testing"

func TestParseArgs(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    Options
		wantErr bool
	}{
		{name: "happy path", args: []string{"img.png"}, want: Options{Path: "img.png", Width: 0}, wantErr: false},
		{name: "short flag", args: []string{"-w", "80", "img.png"}, want: Options{Path: "img.png", Width: 80}, wantErr: false},
		{name: "long flag", args: []string{"--width", "80", "img.png"}, want: Options{Path: "img.png", Width: 80}, wantErr: false},
		{name: "missing path", args: []string{"-w", "80"}, wantErr: true},
		{name: "extra positional", args: []string{"a.png", "b.png"}, wantErr: true},
		{name: "negative width", args: []string{"-w", "-5", "img.png"}, wantErr: true},
		{name: "bad width value", args: []string{"-w", "abc", "img.png"}, wantErr: true},
		{name: "flag after path", args: []string{"img.png", "-w", "80"}, wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseArgs(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseArgs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got != tt.want {
				t.Errorf("ParseArgs() = %v, want %v", got, tt.want)
			}
			if tt.wantErr && (err == nil || err.Error() == "") {
				t.Errorf("ParseArgs() expected an error with a non-empty message, got: %v", err)
			}
		})
	}
}
