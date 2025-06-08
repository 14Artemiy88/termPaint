package menu

import (
	"github.com/14Artemiy88/termPaint/src/pixel"
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

func DrawMenu(screen Screen) {
	switch Type {
	case SymbolColor:
		drawSymbolColorMenu(screen)
	case File:
		fileMenu(screen)
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

func ClearMenu(screen Screen, pixels [][]string, width int) [][]string {
	white := pixel.GetConstColor(pixel.White)
	gray := pixel.GetConstColor(pixel.Gray)

	for y := 0; y < screen.GetHeight(); y++ {
		for x := 0; x < width; x++ {
			utils.SetByKeys(x, y, " ", white, pixels)
		}

		utils.SetByKeys(width, y, "│", gray, pixels)
	}

	return pixels
}
