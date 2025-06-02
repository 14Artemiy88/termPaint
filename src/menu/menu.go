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

func DrawMenu(s Screen) {
	switch Type {
	case SymbolColor:
		drawSymbolColorMenu(s)
	case File:
		fileMenu(s)
	case Help:
		drawHelpMenu(s)
	case Shape:
		drawShapeMenu(s)
	case Line:
		drawLineMenu(s)
	case Config:
		drawConfigMenu(s)
	case None:
	}
}

func ClearMenu(s Screen, screen [][]string, width int) [][]string {
	white := pixel.GetConstColor("white")
	gray := pixel.GetConstColor("gray")

	for y := 0; y < s.GetHeight(); y++ {
		for x := 0; x < width; x++ {
			utils.SetByKeys(x, y, " ", white, screen)
		}

		utils.SetByKeys(width, y, "│", gray, screen)
	}

	return screen
}
