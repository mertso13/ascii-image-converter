package decoder

import (
	"fmt"
	"image"
	"io"
	"os"

	_ "golang.org/x/image/bmp"
	_ "golang.org/x/image/webp"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

func Decode(r io.Reader) (image.Image, error) {
	decodedImage, _, err := image.Decode(r)
	if err != nil {
		return nil, fmt.Errorf("decoding image: %w", err)
	}
	return decodedImage, nil
}

func DecodeFile(path string) (image.Image, error) {
	imageFile, err := os.OpenFile(path, os.O_RDONLY, 0)
	if err != nil {
		return nil, fmt.Errorf("reading image file: %w", err)
	}

	decodedImage, err := Decode(imageFile)

	defer imageFile.Close()

	return decodedImage, err
}
