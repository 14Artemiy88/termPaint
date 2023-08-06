package menu

import (
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/size"
	"github.com/14Artemiy88/termPaint/src/utils"
	"strconv"
	"strings"
)

const SymbolColorWidth = 15

type InputStruct struct {
	Lock  bool
	Value string
	Color string
}

var Input InputStruct

var Colors = map[int]string{
	17: "r",
	19: "g",
	21: "b",
}

const colorX = 3

func DrawSymbolColorMenu(screen [][]string) [][]string {
	ClearMenu(screen, SymbolColorWidth)
	drawSymbolMenu(screen)
	drawColorMenu(screen)
	str := "Help " + strings.Repeat("─", SymbolColorWidth-len("Help")-2) + "┤"
	utils.DrawString(1, size.Size.Height-2, str, screen)
	utils.DrawString(2, size.Size.Height-1, "Press Ctrl+H", screen)

	return screen
}

func drawSymbolMenu(screen [][]string) [][]string {
	str := "Symbol " + strings.Repeat("─", SymbolColorWidth-len("Symbol")-2) + "┐"
	utils.DrawString(1, 1, str, screen)
	for y, line := range config.Cfg.Symbols {
		for x, symbol := range line {
			utils.SetByKeys(x, y, symbol, screen)
		}
	}

	return screen
}

func drawColorMenu(screen [][]string) [][]string {
	str := "Color " + strings.Repeat("─", SymbolColorWidth-len("Color")-2) + "┤"
	utils.DrawString(1, 15, str, screen)
	for y, line := range Colors {
		utils.DrawString(colorX+2, y, strconv.Itoa(cursor.CC.Color[line]), screen)
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
