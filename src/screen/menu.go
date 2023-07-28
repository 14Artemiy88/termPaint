package screen

import (
	"github.com/14Artemiy88/termPaint/src/utils"
	"strings"
)

type menuType int

const (
	None menuType = iota
	symbolColor
	file
	help
	shape
	line
)

var MenuType menuType

func drawMenu(s *Screen, screen [][]string) {
	switch MenuType {
	case symbolColor:
		DrawSymbolColorMenu(s, screen)
	case file:
		FileMenu(s, screen, Dir)
	case help:
		DrawHelpMenu(s, screen)
	case shape:
		drawShapeMenu(s, screen)
	case line:
		drawLineMenu(s, screen)
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
	for y := 0; y < Size.Rows; y++ {
		for x := 0; x < width; x++ {
			utils.SetByKeys(x, y, " ", screen)
		}
		utils.SetByKeys(width, y, "│", screen)
	}

	return screen
}
