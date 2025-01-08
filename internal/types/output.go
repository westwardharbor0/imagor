package types

import (
	"slices"
	"strings"
)

var (
	OutputTypeConsole OutputType = "console"
	OutputTypeTable   OutputType = "table"

	OutputTypesList OutputTypes = []OutputType{OutputTypeConsole, OutputTypeTable}
)

type OutputTypes []OutputType

// StrList converts the list of output types to string list.
func (ot OutputTypes) StrList() string {
	result := make([]string, len(ot))

	for index, value := range ot {
		result[index] = value.String()
	}

	return strings.Join(result, ",")
}

type OutputType string

func (ot OutputType) FromString(s string) OutputType {
	return (OutputType)(s)
}

// Valid checks the validation of given string as output type.
func (ot OutputType) Valid(s string) bool {
	o := ot.FromString(s)

	return slices.Contains(OutputTypesList, o)
}

func (ot OutputType) String() string {
	return (string)(ot)
}
