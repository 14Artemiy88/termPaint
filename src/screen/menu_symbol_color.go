package screen

import (
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/utils"
	"strconv"
	"strings"
)

const MenuSymbolColorWidth = 15

type Input struct {
	lock  bool
	value string
	color string
}

var input Input

var Colors = map[int]string{
	17: "r",
	19: "g",
	21: "b",
}

const colorX = 3

func DrawSymbolColorMenu(screen [][]string) [][]string {
	ClearMenu(screen, MenuSymbolColorWidth)
	drawSymbolMenu(screen)
	drawColorMenu(screen)
	str := "Help " + strings.Repeat("─", MenuSymbolColorWidth-len("Help")-2) + "┤"
	DrawString(1, Size.Rows-2, str, screen)
	DrawString(2, Size.Rows-1, "Press Enter", screen)

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

func drawColorMenu(screen [][]string) [][]string {
	str := "Color " + strings.Repeat("─", MenuSymbolColorWidth-len("Color")-2) + "┤"
	DrawString(1, 15, str, screen)
	for y, line := range Colors {
		DrawString(colorX+2, y, strconv.Itoa(cursor.CC.Color[line]), screen)
		switch line {
		case "r":
			utils.SetByKeys(colorX, y, utils.FgRgb(cursor.CC.Color[line], 0, 0, "█"), screen)
		case "g":
			utils.SetByKeys(colorX, y, utils.FgRgb(0, cursor.CC.Color[line], 0, "█"), screen)
		case "b":
			utils.SetByKeys(colorX, y, utils.FgRgb(0, 0, cursor.CC.Color[line], "█"), screen)
		}
	}
	utils.SetByKeys(3, 23, utils.FgRgb(cursor.CC.Color["r"], cursor.CC.Color["g"], cursor.CC.Color["b"], "█"), screen)
	utils.SetByKeys(4, 23, utils.FgRgb(cursor.CC.Color["r"], cursor.CC.Color["g"], cursor.CC.Color["b"], "█"), screen)
	utils.SetByKeys(5, 23, utils.FgRgb(cursor.CC.Color["r"], cursor.CC.Color["g"], cursor.CC.Color["b"], "█"), screen)
	utils.SetByKeys(3, 24, utils.FgRgb(cursor.CC.Color["r"], cursor.CC.Color["g"], cursor.CC.Color["b"], "█"), screen)
	utils.SetByKeys(4, 24, utils.FgRgb(cursor.CC.Color["r"], cursor.CC.Color["g"], cursor.CC.Color["b"], "█"), screen)
	utils.SetByKeys(5, 24, utils.FgRgb(cursor.CC.Color["r"], cursor.CC.Color["g"], cursor.CC.Color["b"], "█"), screen)

	return screen
}
