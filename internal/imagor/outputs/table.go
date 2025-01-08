package outputs

import (
	"fmt"
	"image/color"
	"slices"
	"strings"
	"time"

	"github.com/xuri/excelize/v2"
)

// OutputExcelFile outputs the given image data to xlsx file named with timestamp.
//
//nolint:funlen,gocognit,cyclop
func OutputExcelFile(data [][]color.RGBA) {
	excelFile := excelize.NewFile()
	defer func() {
		if err := excelFile.Close(); err != nil {
			panic(err.Error())
		}
	}()

	// FIXME: This can be done really easily using Mr. Zajic approach.
	for yIndex, yVal := range data {
		indexer := make([]int, 1)
		pointer := 0

		for iter := 0; iter <= len(yVal)-1; iter++ {
			if indexer[pointer]%26 == 0 && indexer[pointer] != 0 {
				var incNext bool
				for asd := len(indexer) - 1; asd >= 0; asd-- {
					if incNext {
						incNext = false
						indexer[asd]++
					}

					if indexer[asd]%26 == 0 && indexer[asd] != 0 {
						incNext = true
						indexer[asd] = 0
					}
				}

				if slices.Max(indexer) == 0 && slices.Min(indexer) == 0 {
					indexer = make([]int, len(indexer)+1)
					pointer++
				}
			}

			coor := new(strings.Builder)
			for _, c := range indexer {
				coor.WriteRune(rune('A' - 1 + (c + 1)))
			}

			coorString := coor.String()
			cellColor := yVal[iter]
			indexer[pointer]++

			style, err := createCellStyle(cellColor, excelFile)
			if err != nil {
				panic(err.Error())
			}

			cellCoor := fmt.Sprintf("%s%d", coorString, yIndex+1)
			if err = excelFile.SetCellStyle("Sheet1", cellCoor, cellCoor, style); err != nil {
				panic(err.Error())
			}

			if err = excelFile.SetColWidth("Sheet1", coorString, coorString, 3); err != nil {
				panic(err.Error())
			}
		}
	}

	if err := excelFile.SaveAs(fmt.Sprintf("img-%d.xlsx", time.Now().UTC().UnixMilli())); err != nil {
		panic(err.Error())
	}
}

func createCellStyle(cellColor color.RGBA, excelFile *excelize.File) (int, error) {
	hexColor := fmt.Sprintf("#%02x%02x%02x", cellColor.R, cellColor.G, cellColor.B)
	style, err := excelFile.NewStyle(&excelize.Style{
		Fill: excelize.Fill{
			Type:    "gradient",
			Color:   []string{hexColor, hexColor},
			Shading: 1,
		},
		Alignment: &excelize.Alignment{
			Horizontal:      "center",
			JustifyLastLine: true,
			ShrinkToFit:     true,
			WrapText:        true,
		},
	})

	return style, fmt.Errorf("failed to create cell style: %w", err)
}

/*
	import string
	def excelize(idx):
	    s = []
	    while True:
	        s.append(string.ascii_uppercase[idx % len(string.ascii_uppercase)])
	        idx = idx // len(string.ascii_uppercase)
	        if not idx:
	            break
	        idx = idx - 1
	    s.reverse()
	    return(''.join(s))
	print(excelize(1000))
*/
