package menu

import (
	"github.com/14Artemiy88/termPaint/src/pixel"
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

func DrawMenu(s Screen) [][]string {
	screen := s.GetPixels()
	switch Type {
	case SymbolColor:
		drawSymbolColorMenu(screen)
	case File:
		fileMenu(s, screen)
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

	return screen
}

func ClearMenu(screen [][]string, width int) [][]string {
	for y := 0; y < size.Size.Height; y++ {
		for x := 0; x < width; x++ {
			utils.SetByKeys(x, y, " ", pixel.White, screen)
		}
		utils.SetByKeys(width, y, "│", pixel.Gray, screen)
	}

	return screen
}
