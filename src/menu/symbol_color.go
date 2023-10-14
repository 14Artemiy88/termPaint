package menu

import (
	"github.com/14Artemiy88/termPaint/src/color"
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

func drawSymbolColorMenu(screen [][]string) [][]string {
	ClearMenu(screen, SymbolColorWidth)
	drawSymbolMenu(screen)
	drawColorMenu(screen)
	title := "Help"
	str := strings.Repeat("─", SymbolColorWidth-len(title)-2) + "┤"
	utils.DrawString(1, size.Size.Height-3, title, color.Yellow, screen)
	utils.DrawString(len(title)+2, size.Size.Height-3, str, color.Gray, screen)
	utils.DrawString(2, size.Size.Height-1, "Press", color.White, screen)
	utils.DrawString(len("Press")+3, size.Size.Height-1, "Ctrl+H", color.Green, screen)

	return screen
}

func drawSymbolMenu(screen [][]string) [][]string {
	title := "Symbol"
	str := strings.Repeat("─", SymbolColorWidth-len(title)-2) + "┐"
	utils.DrawString(1, 1, title, color.Yellow, screen)
	utils.DrawString(len(title)+2, 1, str, color.Gray, screen)
	for y, line := range config.Cfg.Symbols {
		for x, symbol := range line {
			utils.SetByKeys(x, y, symbol, color.White, screen)
		}
	}

	return screen
}

func drawColorMenu(screen [][]string) [][]string {
	title := "Color"
	str := strings.Repeat("─", SymbolColorWidth-len(title)-2) + "┤"
	utils.DrawString(1, 15, title, color.Yellow, screen)
	utils.DrawString(len(title)+2, 15, str, color.Gray, screen)
	for y, line := range Colors {
		switch line {
		case "r":
			utils.DrawString(colorX+2, y, strconv.Itoa(cursor.CC.Color.R), color.White, screen)
			utils.SetByKeys(colorX, y, utils.FgRgb(color.Color{R: cursor.CC.Color.R}, "█"), color.White, screen)
		case "g":
			utils.DrawString(colorX+2, y, strconv.Itoa(cursor.CC.Color.G), color.White, screen)
			utils.SetByKeys(colorX, y, utils.FgRgb(color.Color{G: cursor.CC.Color.G}, "█"), color.White, screen)
		case "b":
			utils.DrawString(colorX+2, y, strconv.Itoa(cursor.CC.Color.B), color.White, screen)
			utils.SetByKeys(colorX, y, utils.FgRgb(color.Color{B: cursor.CC.Color.B}, "█"), color.White, screen)
		}
	}
	utils.SetByKeys(3, 23, "█", cursor.CC.Color, screen)
	utils.SetByKeys(4, 23, "█", cursor.CC.Color, screen)
	utils.SetByKeys(5, 23, "█", cursor.CC.Color, screen)
	utils.SetByKeys(3, 24, "█", cursor.CC.Color, screen)
	utils.SetByKeys(4, 24, "█", cursor.CC.Color, screen)
	utils.SetByKeys(5, 24, "█", cursor.CC.Color, screen)

	return screen
}
