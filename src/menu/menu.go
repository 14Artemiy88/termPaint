package menu

import (
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/size"
	"github.com/14Artemiy88/termPaint/src/utils"
)

type menuType int

const (
	None menuType = iota
	SymbolColor
	File
	Help
	Shape
	Line
)

var Type menuType

func DrawMenu(screen [][]string) {
	switch Type {
	case SymbolColor:
		DrawSymbolColorMenu(screen)
	case File:
		FileMenu(screen, Dir)
	case Help:
		DrawHelpMenu(screen)
	case Shape:
		drawShapeMenu(screen)
	case Line:
		drawLineMenu(screen)
	}
}

func ClearMenu(screen [][]string, width int) [][]string {
	for y := 0; y < size.Size.Height; y++ {
		for x := 0; x < width; x++ {
			utils.SetByKeys(x, y, " ", color.White, screen)
		}
		utils.SetByKeys(width, y, "│", color.Gray, screen)
	}

	return screen
}
