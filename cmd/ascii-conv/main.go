package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/mertso13/ascii-image-converter/pkg/converter"
	"github.com/mertso13/ascii-image-converter/pkg/decoder"
	"github.com/mertso13/ascii-image-converter/pkg/scaler"
	"github.com/mertso13/ascii-image-converter/pkg/term"
)

type Options struct {
	Path       string
	Width      int
	Height     int
	Palette    string
	CustomRamp string
	Invert     bool
	Scale      float64
	Filter     string
	Output     string
}

func ParseArgs(args []string) (Options, error) {
	fs := flag.NewFlagSet("ascii-conv", flag.ContinueOnError)

	opts := Options{}

	fs.IntVar(&opts.Width, "w", 0, "target width")
	fs.IntVar(&opts.Width, "width", 0, "target width")
	fs.IntVar(&opts.Height, "H", 0, "target height")
	fs.IntVar(&opts.Height, "height", 0, "target height")
	fs.StringVar(&opts.Palette, "p", "standard", "palette name")
	fs.StringVar(&opts.Palette, "palette", "standard", "palette name")
	fs.StringVar(&opts.CustomRamp, "r", "", "custom character ramp")
	fs.StringVar(&opts.CustomRamp, "ramp", "", "custom character ramp")
	fs.BoolVar(&opts.Invert, "i", false, "invert character ramp")
	fs.BoolVar(&opts.Invert, "invert", false, "invert character ramp")
	fs.Float64Var(&opts.Scale, "s", scaler.DefaultAspectFactor, "font aspect ratio factor")
	fs.Float64Var(&opts.Scale, "scale", scaler.DefaultAspectFactor, "font aspect ratio factor")
	fs.StringVar(&opts.Filter, "f", "bilinear", "resampling filter")
	fs.StringVar(&opts.Filter, "filter", "bilinear", "resampling filter")
	fs.StringVar(&opts.Output, "o", "", "output file path")
	fs.StringVar(&opts.Output, "output", "", "output file path")

	err := fs.Parse(args)
	if err != nil {
		return Options{}, err
	}

	if opts.Width < 0 {
		return Options{}, fmt.Errorf("width must be positive or zero: %d", opts.Width)
	}

	if opts.Height < 0 {
		return Options{}, fmt.Errorf("height must be positive or zero: %d", opts.Height)
	}

	if opts.Scale <= 0 {
		return Options{}, fmt.Errorf("scale must be greater than zero: %f", opts.Scale)
	}

	positionalArguments := fs.Args()
	if len(positionalArguments) == 0 {
		return Options{}, fmt.Errorf("missing image path")
	}

	if len(positionalArguments) > 1 {
		return Options{}, fmt.Errorf("extra positional arguments: %v", positionalArguments[1:])
	}

	opts.Path = positionalArguments[0]

	return opts, nil
}

func main() {
	opts, err := ParseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	img, err := decoder.DecodeFile(opts.Path)
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

	targetWidth := opts.Width
	targetHeight := opts.Height

	if targetWidth == 0 && targetHeight == 0 {
		terminalWidth, _ := term.GetSizeOrDefault(int(os.Stdout.Fd()))
		targetWidth = terminalWidth
	}

	origWidth := img.Bounds().Dx()
	origHeight := img.Bounds().Dy()

	scaledWidth, scaledHeight := scaler.CalculateDimensions(origWidth, origHeight, targetWidth, targetHeight, opts.Scale)
	scaledImage := scaler.Scale(img, scaledWidth, scaledHeight, scaler.FilterByName(opts.Filter))

	var ramp *converter.Ramp
	if opts.CustomRamp != "" {
		ramp = converter.NewRamp(opts.CustomRamp)
	} else {
		var ok bool
		ramp, ok = converter.ByName(opts.Palette)
		if !ok {
			fmt.Fprintf(os.Stderr, "error: unknown palette %q\n", opts.Palette)
			os.Exit(1)
		}
	}

	if opts.Invert {
		ramp = ramp.Invert()
	}

	asciiArt := converter.Convert(scaledImage, ramp)

	if opts.Output != "" {
		err = os.WriteFile(opts.Output, []byte(asciiArt), 0644)
		if err != nil {
			fmt.Fprintln(os.Stderr, "error writing output file:", err)
			os.Exit(1)
		}
		return
	}

	fmt.Print(asciiArt)
}
