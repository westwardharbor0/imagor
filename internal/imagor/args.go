package imagor

import (
	"errors"
	"flag"
	"fmt"

	"github.com/westwardharbor0/imagor/internal/types"
)

const (
	OutputExtension = ".xlsx"

	DefaultOutputFile = "./output" + OutputExtension
	DefaultPTC        = 50
)

// CliArgs represents parsed arguments.
type CliArgs struct {
	ImagePath   string
	PixelToCell int
	OutputPath  string
	OutputType  string
	Verbose     bool
}

// Validate checks the required arguments and values.
func (c *CliArgs) Validate() error {
	if c.ImagePath == "" {
		return errors.New("argument `image` needs to be provided")
	}

	if !types.OutputTypeConsole.Valid(c.OutputType) {
		return fmt.Errorf("argument `output-type` has invalid value `%s`", c.OutputType)
	}

	return nil
}

// ParseArgs takes input arguments and parses them into structure.
func ParseArgs() *CliArgs {
	imagePath := flag.String("image", "", "Path to image")
	outputPath := flag.String(
		"output-file",
		DefaultOutputFile,
		"Path to result excel sheet file",
	)
	outputType := flag.String(
		"output-type",
		types.OutputTypeConsole.String(),
		"Type of output. Options: "+types.OutputTypesList.StrList(),
	)
	verboseLogging := flag.Bool("verbose", false, "Verbose logging of process")
	pixelToCell := flag.Int("ptc", DefaultPTC, "How many pixel will be represented with one cell.")

	flag.Parse()

	return &CliArgs{
		ImagePath:   *imagePath,
		PixelToCell: *pixelToCell,
		OutputPath:  *outputPath,
		OutputType:  *outputType,
		Verbose:     *verboseLogging,
	}
}
