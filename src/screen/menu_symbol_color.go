package screen

import (
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/utils"
	"strconv"
	"strings"
)

const MenuSymbolColorWidth = 15

var Colors = map[int]string{
	17: "r",
	19: "g",
	21: "b",
}

const colorX = 3

func DrawSymbolColorMenu(s *Screen, screen [][]string) [][]string {
	ClearMenu(s, screen, MenuSymbolColorWidth)
	drawSymbolMenu(screen)
	drawColorMenu(s, screen)
	str := "Help " + strings.Repeat("─", MenuSymbolColorWidth-len("Help")-2) + "┤"
	DrawString(1, s.Rows-2, str, screen)
	DrawString(2, s.Rows-1, "Press Enter", screen)

	return screen
}

func drawSymbolMenu(screen [][]string) [][]string {
	str := "Symbol " + strings.Repeat("─", MenuSymbolColorWidth-len("Symbol")-2) + "┐"
	DrawString(1, 1, str, screen)
	for y, line := range config.Cfg.Symbols {
		for x, symbol := range line {
			utils.SetByKeys(x, y, symbol, screen)
		}
	}

	return screen
}

func drawColorMenu(s *Screen, screen [][]string) [][]string {
	str := "Color " + strings.Repeat("─", MenuSymbolColorWidth-len("Color")-2) + "┤"
	DrawString(1, 15, str, screen)
	for y, line := range Colors {
		DrawString(colorX+2, y, strconv.Itoa(s.Cursor.Color[line]), screen)
		switch line {
		case "r":
			utils.SetByKeys(colorX, y, utils.FgRgb(s.Cursor.Color[line], 0, 0, "█"), screen)
		case "g":
			utils.SetByKeys(colorX, y, utils.FgRgb(0, s.Cursor.Color[line], 0, "█"), screen)
		case "b":
			utils.SetByKeys(colorX, y, utils.FgRgb(0, 0, s.Cursor.Color[line], "█"), screen)
		}
	}
	utils.SetByKeys(3, 23, utils.FgRgb(s.Cursor.Color["r"], s.Cursor.Color["g"], s.Cursor.Color["b"], "█"), screen)
	utils.SetByKeys(4, 23, utils.FgRgb(s.Cursor.Color["r"], s.Cursor.Color["g"], s.Cursor.Color["b"], "█"), screen)
	utils.SetByKeys(5, 23, utils.FgRgb(s.Cursor.Color["r"], s.Cursor.Color["g"], s.Cursor.Color["b"], "█"), screen)
	utils.SetByKeys(3, 24, utils.FgRgb(s.Cursor.Color["r"], s.Cursor.Color["g"], s.Cursor.Color["b"], "█"), screen)
	utils.SetByKeys(4, 24, utils.FgRgb(s.Cursor.Color["r"], s.Cursor.Color["g"], s.Cursor.Color["b"], "█"), screen)
	utils.SetByKeys(5, 24, utils.FgRgb(s.Cursor.Color["r"], s.Cursor.Color["g"], s.Cursor.Color["b"], "█"), screen)

	return screen
}