package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mertso13/ascii-image-converter/pkg/converter"
	"github.com/mertso13/ascii-image-converter/pkg/decoder"
)

// Options holds the parsed commandline configuration for one run.
type Options struct {
	Path  string
	Width int
}

// ParseArgs parses args into Options.
func ParseArgs(args []string) (Options, error) {
	fs := flag.NewFlagSet("ascii-conv", flag.ContinueOnError)

	opts := Options{}

	fs.IntVar(&opts.Width, "w", 0, "target width in characters")
	fs.IntVar(&opts.Width, "width", 0, "target width in characters")

	err := fs.Parse(args)
	if err != nil {
		return Options{}, err
	}

	if opts.Width < 0 {
		return Options{}, fmt.Errorf("width cannot be negative: %d", opts.Width)
	}

	lenArgs := len(fs.Args())
	if lenArgs == 0 {
		return Options{}, fmt.Errorf("missing image path")
	} else if lenArgs > 1 {
		return Options{}, fmt.Errorf("extra positional arguments: %v", fs.Args()[1:])
	}

	return Options{Path: fs.Args()[0], Width: opts.Width}, nil
}

func main() {
	opts, err := ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		fmt.Fprintln(os.Stderr, "usage: ascii-conv [-w width] <image-path>")
		os.Exit(1)
	}

	img, err := decoder.DecodeFile(opts.Path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	ramp, ok := converter.ByName("standard")
	if !ok {
		fmt.Fprintln(os.Stderr, "error: standard ramp not found")
		os.Exit(1)
	}

	fmt.Print(converter.Convert(img, ramp))
}
