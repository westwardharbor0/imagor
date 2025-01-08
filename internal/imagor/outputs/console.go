package outputs

import (
	"fmt"
	"image/color"
)

// OutputConsole outputs the image data to console output.
func OutputConsole(data [][]color.RGBA) {
	for _, yVal := range data {
		for _, xVal := range yVal {
			print(
				fmt.Sprintf(
					"\033[48;2;%d;%d;%dm  \033[0m",
					xVal.R,
					xVal.G,
					xVal.B,
				),
			)
		}

		print("\n")
	}
}
