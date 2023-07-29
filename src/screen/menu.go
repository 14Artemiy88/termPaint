package screen

import (
	"github.com/14Artemiy88/termPaint/src/size"
	"github.com/14Artemiy88/termPaint/src/utils"
	"strings"
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

var MenuType menuType

func drawMenu(screen [][]string) {
	switch MenuType {
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

func DrawString(X int, Y int, val string, screen [][]string) [][]string {
	str := strings.Split(val, "")
	for k, symbol := range str {
		utils.SetByKeys(X+k, Y, symbol, screen)
	}

	return screen
}

func ClearMenu(screen [][]string, width int) [][]string {
	for y := 0; y < size.Size.Rows; y++ {
		for x := 0; x < width; x++ {
			utils.SetByKeys(x, y, " ", screen)
		}
		utils.SetByKeys(width, y, "│", screen)
	}

	return screen
}
