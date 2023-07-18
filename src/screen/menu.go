package screen

import (
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/utils"
	"strconv"
	"strings"
)

const MenuWidth = 15

var Colors = map[int]string{
	15: "r",
	17: "g",
	19: "b",
}

const colorX = 2

func DrawMenu(s *Screen, screen [][]string) [][]string {
	ClearMenu(s, screen, MenuWidth)
	drawSymbolMenu(screen)
	drawColorMenu(s, screen)
	str := "Help " + strings.Repeat("─", MenuWidth-len("Help ")) + "┤"
	DrawString(0, s.Rows-2, str, screen)
	DrawString(1, s.Rows-1, "Press Enter", screen)

	return screen
}

func drawSymbolMenu(screen [][]string) [][]string {
	str := "Symbol " + strings.Repeat("─", MenuWidth-len("Symbol ")) + "┐"
	DrawString(0, 0, str, screen)
	for y, line := range config.Cfg.Symbols {
		for x, symbol := range line {
			utils.SetByKeys(x, y, symbol, screen)
		}
	}

	return screen
}

func drawColorMenu(s *Screen, screen [][]string) [][]string {
	str := "Color " + strings.Repeat("─", MenuWidth-len("Color ")) + "┤"
	DrawString(0, 13, str, screen)
	for y, line := range Colors {
		DrawString(4, y, strconv.Itoa(s.Cursor.Color[line]), screen)
		switch line {
		case "r":
			utils.SetByKeys(colorX, y, utils.FgRgb(s.Cursor.Color[line], 0, 0, "█"), screen)
		case "g":
			utils.SetByKeys(colorX, y, utils.FgRgb(0, s.Cursor.Color[line], 0, "█"), screen)
		case "b":
			utils.SetByKeys(colorX, y, utils.FgRgb(0, 0, s.Cursor.Color[line], "█"), screen)
		}
	}
	utils.SetByKeys(3, 21, utils.FgRgb(s.Cursor.Color["r"], s.Cursor.Color["g"], s.Cursor.Color["b"], "█"), screen)
	utils.SetByKeys(4, 21, utils.FgRgb(s.Cursor.Color["r"], s.Cursor.Color["g"], s.Cursor.Color["b"], "█"), screen)
	utils.SetByKeys(5, 21, utils.FgRgb(s.Cursor.Color["r"], s.Cursor.Color["g"], s.Cursor.Color["b"], "█"), screen)
	utils.SetByKeys(3, 22, utils.FgRgb(s.Cursor.Color["r"], s.Cursor.Color["g"], s.Cursor.Color["b"], "█"), screen)
	utils.SetByKeys(4, 22, utils.FgRgb(s.Cursor.Color["r"], s.Cursor.Color["g"], s.Cursor.Color["b"], "█"), screen)
	utils.SetByKeys(5, 22, utils.FgRgb(s.Cursor.Color["r"], s.Cursor.Color["g"], s.Cursor.Color["b"], "█"), screen)

	return screen
}

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
