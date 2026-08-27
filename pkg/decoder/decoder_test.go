package decoder

import (
	"bytes"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestDecode(t *testing.T) {
	rect := image.Rect(0, 0, 2, 2)
	testImage := image.NewRGBA(rect)
	buffer := new(bytes.Buffer)
	encodedImage := png.Encode(buffer, testImage)
	if encodedImage != nil {
		t.Fatalf("failed to encode test image: %v", encodedImage)
	}

	decodedImage, err := Decode(buffer)
	if err != nil {
		t.Fatalf("Decode returned an error: %v", err)
	}
	if decodedImage.Bounds() != rect {
		t.Fatalf("expected bounds %v, got %v", rect, decodedImage.Bounds())
	}
}

func TestDecodeInvalid(t *testing.T) {
	invalidImage := strings.NewReader("not an image")
	_, err := Decode(invalidImage)
	if err == nil {
		t.Fatalf("Decode did not return an error for invalid input")
	}
}

func TestDecodeFile(t *testing.T) {
	tempFilePath := filepath.Join(t.TempDir(), "test.png")
	rect := image.Rect(0, 0, 2, 2)
	tempImage := image.NewRGBA(rect)
	tempFile, err := os.Create(tempFilePath)
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	err = png.Encode(tempFile, tempImage)
	tempFile.Close()
	if err != nil {
		t.Fatalf("failed to encode image to temp file: %v", err)
	}

	decodedImage, err := DecodeFile(tempFilePath)
	if err != nil {
		t.Fatalf("DecodeFile returned an error: %v", err)
	}
	if decodedImage.Bounds() != rect {
		t.Fatalf("expected bounds %v, got %v", rect, decodedImage.Bounds())
	}
}

func TestDecodeFileNonExistent(t *testing.T) {
	someRandomPath := filepath.Join(t.TempDir(), "nonexistent.png")
	_, err := DecodeFile(someRandomPath)
	if err == nil {
		t.Fatalf("DecodeFile didn't return an error for non-existent file")
	}
}
