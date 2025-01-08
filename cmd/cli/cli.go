package main

import (
	"flag"
	"fmt"
	"strings"

	"github.com/westwardharbor0/imagor/internal/imagor"
	"github.com/westwardharbor0/imagor/internal/imagor/outputs"
	"github.com/westwardharbor0/imagor/internal/types"
)

func main() {
	cliArgs := imagor.ParseArgs()

	if err := cliArgs.Validate(); err != nil {
		flag.Usage()
		println(strings.Repeat("--", 20))
		panic(err.Error())
	}

	imagor := imagor.NewImagor(
		cliArgs.ImagePath,
		cliArgs.PixelToCell,
	)
	proceessedImage, err := imagor.GenerateGrid()

	if err != nil {
		panic(err.Error())
	}

	switch types.OutputTypeConsole.FromString(cliArgs.OutputType) {
	case types.OutputTypeConsole:
		outputs.OutputConsole(proceessedImage)
	case types.OutputTypeTable:
		outputs.OutputExcelFile(proceessedImage)
	default:
		fmt.Printf("Unsupported output format %s", cliArgs.OutputType)
	}
}
