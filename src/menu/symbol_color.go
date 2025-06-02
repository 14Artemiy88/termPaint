package menu

import (
	"strconv"
	"strings"

	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/utils"
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

func drawSymbolColorMenu(s Screen) {
	screen := s.GetPixels()
	ClearMenu(s, screen, SymbolColorWidth)
	drawSymbolMenu(s, screen)
	drawColorMenu(screen)

	title := "Help"
	str := strings.Repeat("─", SymbolColorWidth-len(title)-2) + "┤"
	height := s.GetHeight()
	utils.DrawString(1, height-3, title, pixel.GetConstColor("yellow"), screen)
	utils.DrawString(len(title)+2, height-3, str, pixel.GetConstColor("gray"), screen)
	utils.DrawString(2, height-1, "Press", pixel.GetConstColor("white"), screen)
	utils.DrawString(len("Press")+3, height-1, "Ctrl+H", pixel.GetConstColor("green"), screen)
}

func drawSymbolMenu(s Screen, screen [][]string) [][]string {
	white := pixel.GetConstColor("white")
	title := "Symbol"
	str := strings.Repeat("─", SymbolColorWidth-len(title)-2) + "┐"
	utils.DrawString(1, 1, title, pixel.GetConstColor("yellow"), screen)
	utils.DrawString(len(title)+2, 1, str, pixel.GetConstColor("gray"), screen)

	for y, line := range s.GetConfig().Symbols {
		for x, symbol := range line {
			utils.SetByKeys(x, y, symbol, white, screen)
		}
	}

	return screen
}

func drawColorMenu(screen [][]string) [][]string {
	white := pixel.GetConstColor("white")
	title := "Color"
	str := strings.Repeat("─", SymbolColorWidth-len(title)-2) + "┤"
	utils.DrawString(1, 15, title, pixel.GetConstColor("yellow"), screen)
	utils.DrawString(len(title)+2, 15, str, pixel.GetConstColor("gray"), screen)

	for y, line := range Colors {
		switch line {
		case "r":
			utils.DrawString(colorX+2, y, strconv.Itoa(cursor.CC.Color.R), white, screen)
			utils.SetByKeys(colorX, y, utils.FgRgb(pixel.Color{R: cursor.CC.Color.R}, "█"), white, screen)
		case "g":
			utils.DrawString(colorX+2, y, strconv.Itoa(cursor.CC.Color.G), white, screen)
			utils.SetByKeys(colorX, y, utils.FgRgb(pixel.Color{G: cursor.CC.Color.G}, "█"), white, screen)
		case "b":
			utils.DrawString(colorX+2, y, strconv.Itoa(cursor.CC.Color.B), white, screen)
			utils.SetByKeys(colorX, y, utils.FgRgb(pixel.Color{B: cursor.CC.Color.B}, "█"), white, screen)
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
