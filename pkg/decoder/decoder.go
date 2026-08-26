package decoder

import (
	"fmt"
	"image"
	"io"
	"os"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

// Decode decodes an image from r. PNG, JPEG, and GIF are supported;
// the format is detected from the bytes themselves.
func Decode(r io.Reader) (image.Image, error) {
	decodedImage, _, err := image.Decode(r)
	if err != nil {
		return nil, fmt.Errorf("decoding image: %w", err)
	}
	return decodedImage, nil
}

// DecodeFile opens path and decodes it. The file is closed before
// DecodeFile returns, success or failure.
func DecodeFile(path string) (image.Image, error) {
	imageFile, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, fmt.Errorf("reading image file: %w", err)
	}

	decodedImage, err := Decode(imageFile)

	defer imageFile.Close()

	return decodedImage, err
}
