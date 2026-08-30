package term

import (
	"golang.org/x/term"
)

const (
	DefaultWidth  = 80
	DefaultHeight = 24
)

func GetSize(fileDescriptor int) (int, int, error) {
	return term.GetSize(fileDescriptor)
}

func GetSizeOrDefault(fileDescriptor int) (int, int) {
	width, height, err := GetSize(fileDescriptor)
	if err != nil {
		return DefaultWidth, DefaultHeight
	}

	return width, height
}
