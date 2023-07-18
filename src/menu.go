package src

import (
	"strconv"
	"strings"
)

const menuWidth = 15

var colors = map[int]string{
	15: "r",
	17: "g",
	19: "b",
}

const colorX = 2

func drawMenu(s *Screen, screen [][]string) [][]string {
	ClearMenu(s, screen, menuWidth)
	drawSymbolMenu(screen)
	drawColorMenu(s, screen)
	str := "Help " + strings.Repeat("─", menuWidth-len("Help ")) + "┤"
	DrawString(0, s.Rows-2, str, screen)
	DrawString(1, s.Rows-1, "Press Enter", screen)

	return screen
}

func drawSymbolMenu(screen [][]string) [][]string {
	str := "Symbol " + strings.Repeat("─", menuWidth-len("symbol ")) + "┐"
	DrawString(0, 0, str, screen)
	for y, line := range Cfg.Symbols {
		for x, symbol := range line {
			SetByKeys(x, y, symbol, screen)
		}
	}

	return screen
}

func drawColorMenu(s *Screen, screen [][]string) [][]string {
	str := "Color " + strings.Repeat("─", menuWidth-len("Color ")) + "┤"
	DrawString(0, 13, str, screen)
	for y, line := range colors {
		DrawString(4, y, strconv.Itoa(s.Color[line]), screen)
		switch line {
		case "r":
			SetByKeys(colorX, y, FgRgb(s.Color[line], 0, 0, "█"), screen)
		case "g":
			SetByKeys(colorX, y, FgRgb(0, s.Color[line], 0, "█"), screen)
		case "b":
			SetByKeys(colorX, y, FgRgb(0, 0, s.Color[line], "█"), screen)
		}
	}
	SetByKeys(3, 21, FgRgb(s.Color["r"], s.Color["g"], s.Color["b"], "█"), screen)
	SetByKeys(4, 21, FgRgb(s.Color["r"], s.Color["g"], s.Color["b"], "█"), screen)
	SetByKeys(5, 21, FgRgb(s.Color["r"], s.Color["g"], s.Color["b"], "█"), screen)
	SetByKeys(3, 22, FgRgb(s.Color["r"], s.Color["g"], s.Color["b"], "█"), screen)
	SetByKeys(4, 22, FgRgb(s.Color["r"], s.Color["g"], s.Color["b"], "█"), screen)
	SetByKeys(5, 22, FgRgb(s.Color["r"], s.Color["g"], s.Color["b"], "█"), screen)

	return screen
}

func DrawString(X int, Y int, val string, screen [][]string) [][]string {
	str := strings.Split(val, "")
	for k, symbol := range str {
		SetByKeys(X+k, Y, symbol, screen)
	}

	return screen
}

func ClearMenu(s *Screen, screen [][]string, width int) [][]string {
	for y := 0; y < s.Rows; y++ {
		for x := 0; x < width; x++ {
			SetByKeys(x, y, " ", screen)
		}
		SetByKeys(width, y, "│", screen)
	}

	return screen
}
