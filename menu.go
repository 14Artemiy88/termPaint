package main

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

func drawMenu(s *screen, screen [][]string) [][]string {
	clearMenu(s, screen, menuWidth)
	drawSymbolMenu(screen)
	drawColorMenu(s, screen)
	str := "Help " + strings.Repeat("─", menuWidth-len("Help ")) + "┤"
	drawString(0, s.rows-2, str, screen)
	drawString(1, s.rows-1, "Press Enter", screen)

	return screen
}

func drawSymbolMenu(screen [][]string) [][]string {
	str := "Symbol " + strings.Repeat("─", menuWidth-len("symbol ")) + "┐"
	drawString(0, 0, str, screen)
	for y, line := range cfg.Symbols {
		for x, symbol := range line {
			setByKeys(x, y, symbol, screen)
		}
	}

	return screen
}

func drawColorMenu(s *screen, screen [][]string) [][]string {
	str := "Color " + strings.Repeat("─", menuWidth-len("Color ")) + "┤"
	drawString(0, 13, str, screen)
	for y, line := range colors {
		drawString(4, y, strconv.Itoa(s.color[line]), screen)
		switch line {
		case "r":
			setByKeys(colorX, y, fgRgb(s.color[line], 0, 0, "█"), screen)
		case "g":
			setByKeys(colorX, y, fgRgb(0, s.color[line], 0, "█"), screen)
		case "b":
			setByKeys(colorX, y, fgRgb(0, 0, s.color[line], "█"), screen)
		}
	}
	setByKeys(3, 21, fgRgb(s.color["r"], s.color["g"], s.color["b"], "█"), screen)
	setByKeys(4, 21, fgRgb(s.color["r"], s.color["g"], s.color["b"], "█"), screen)
	setByKeys(5, 21, fgRgb(s.color["r"], s.color["g"], s.color["b"], "█"), screen)
	setByKeys(3, 22, fgRgb(s.color["r"], s.color["g"], s.color["b"], "█"), screen)
	setByKeys(4, 22, fgRgb(s.color["r"], s.color["g"], s.color["b"], "█"), screen)
	setByKeys(5, 22, fgRgb(s.color["r"], s.color["g"], s.color["b"], "█"), screen)

	return screen
}

func drawString(X int, Y int, val string, screen [][]string) [][]string {
	str := strings.Split(val, "")
	for k, symbol := range str {
		setByKeys(X+k, Y, symbol, screen)
	}

	return screen
}

func clearMenu(s *screen, screen [][]string, width int) [][]string {
	for y := 0; y < s.rows; y++ {
		for x := 0; x < width; x++ {
			setByKeys(x, y, " ", screen)
		}
		setByKeys(width, y, "│", screen)
	}

	return screen
}
