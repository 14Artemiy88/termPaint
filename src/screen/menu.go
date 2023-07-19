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
)

func DrawString(X int, Y int, val string, screen [][]string) [][]string {
	str := strings.Split(val, "")
	for k, symbol := range str {
		utils.SetByKeys(X+k, Y, symbol, screen)
	}

	return screen
}

func ClearMenu(s *Screen, screen [][]string, width int) [][]string {
	for y := 0; y < s.Rows; y++ {
		for x := 0; x < width; x++ {
			utils.SetByKeys(x, y, " ", screen)
		}
		utils.SetByKeys(width, y, "│", screen)
	}

	return screen
}
