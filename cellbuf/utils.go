package cellbuf

import (
	"image/color"
	"strings"

	"github.com/charmbracelet/x/ansi"
)

// Height returns the height of a string.
func Height(s string) int {
	return strings.Count(s, "\n") + 1
}

func readColor(idxp *int, params []ansi.Parameter) (c ansi.Color) {
	i := *idxp
	paramsLen := len(params)
	if i > paramsLen-1 {
		return
	}
	// Note: we accept both main and subparams here
	switch params[i+1].Param(0) {
	case 2: // RGB
		if i > paramsLen-4 {
			return
		}
		c = color.RGBA{
			R: uint8(params[i+2].Param(0)), //nolint:gosec
			G: uint8(params[i+3].Param(0)), //nolint:gosec
			B: uint8(params[i+4].Param(0)), //nolint:gosec
			A: 0xff,
		}
		*idxp += 4
	case 5: // 256 colors
		if i > paramsLen-2 {
			return
		}
		c = ansi.ExtendedColor(params[i+2].Param(0)) //nolint:gosec
		*idxp += 2
	}
	return
}
