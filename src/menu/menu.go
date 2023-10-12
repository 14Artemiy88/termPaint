package menu

import (
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/size"
	"github.com/14Artemiy88/termPaint/src/utils"
)

type MenuType int

const (
	None MenuType = iota
	SymbolColor
	File
	Help
	Shape
	Line
	Config
)

var Type MenuType

func DrawMenu(screen [][]string) {
	switch Type {
	case SymbolColor:
		drawSymbolColorMenu(screen)
	case File:
		fileMenu(screen, Dir)
	case Help:
		drawHelpMenu(screen)
	case Shape:
		drawShapeMenu(screen)
	case Line:
		drawLineMenu(screen)
	case Config:
		drawConfigMenu(screen)
	case None:
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
