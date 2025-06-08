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

func drawSymbolColorMenu(screen Screen) {
	pixels := screen.GetPixels()
	ClearMenu(screen, pixels, SymbolColorWidth)
	drawSymbolMenu(screen, pixels)
	drawColorMenu(pixels)

	title := "Help"
	str := strings.Repeat("─", SymbolColorWidth-len(title)-2) + "┤"
	height := screen.GetHeight()
	utils.DrawString(1, height-3, title, pixel.GetConstColor("yellow"), pixels)
	utils.DrawString(len(title)+2, height-3, str, pixel.GetConstColor("gray"), pixels)
	utils.DrawString(2, height-1, "Press", pixel.GetConstColor("white"), pixels)
	utils.DrawString(len("Press")+3, height-1, "Ctrl+H", pixel.GetConstColor("green"), pixels)
}

func drawSymbolMenu(screen Screen, pixels [][]string) [][]string {
	white := pixel.GetConstColor("white")
	title := "Symbol"
	str := strings.Repeat("─", SymbolColorWidth-len(title)-2) + "┐"
	utils.DrawString(1, 1, title, pixel.GetConstColor("yellow"), pixels)
	utils.DrawString(len(title)+2, 1, str, pixel.GetConstColor("gray"), pixels)

	for y, line := range screen.GetConfig().Symbols {
		for x, symbol := range line {
			utils.SetByKeys(x, y, symbol, white, pixels)
		}
	}

	return pixels
}

func drawColorMenu(screen [][]string) [][]string {
	white := pixel.GetConstColor("white")
	title := "Color"
	str := strings.Repeat("─", SymbolColorWidth-len(title)-2) + "┤"
	utils.DrawString(1, 15, title, pixel.GetConstColor("yellow"), screen)
	utils.DrawString(len(title)+2, 15, str, pixel.GetConstColor("gray"), screen)

	for yCoord, line := range Colors {
		switch line {
		case "r":
			utils.DrawString(colorX+2, yCoord, strconv.Itoa(cursor.CC.Color.R), white, screen)
			utils.SetByKeys(colorX, yCoord, utils.FgRgb(pixel.Color{R: cursor.CC.Color.R}, "█"), white, screen)
		case "g":
			utils.DrawString(colorX+2, yCoord, strconv.Itoa(cursor.CC.Color.G), white, screen)
			utils.SetByKeys(colorX, yCoord, utils.FgRgb(pixel.Color{G: cursor.CC.Color.G}, "█"), white, screen)
		case "b":
			utils.DrawString(colorX+2, yCoord, strconv.Itoa(cursor.CC.Color.B), white, screen)
			utils.SetByKeys(colorX, yCoord, utils.FgRgb(pixel.Color{B: cursor.CC.Color.B}, "█"), white, screen)
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
