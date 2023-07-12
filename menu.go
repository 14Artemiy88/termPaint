package main

import (
	"strconv"
	"strings"
)

const menuWidth = 15
const pointer = "❯"

var symbols = map[int]map[int]string{
	1:  {symbolX: "┌", symbolX + 2: "├", symbolX + 4: "┬", symbolX + 6: "─", symbolX + 7: "┐"},
	3:  {symbolX: "└", symbolX + 2: "┤", symbolX + 4: "┴", symbolX + 6: "│", symbolX + 7: "┘", symbolX + 9: "┼"},
	5:  {symbolX: "┏", symbolX + 2: "┣", symbolX + 4: "┳", symbolX + 6: "━", symbolX + 7: "┓"},
	7:  {symbolX: "┗", symbolX + 2: "┫", symbolX + 4: "┻", symbolX + 6: "┃", symbolX + 7: "┛", symbolX + 9: "╋"},
	9:  {symbolX: "╭", symbolX + 2: "╮", symbolX + 4: "╯", symbolX + 6: "╰"},
	11: {symbolX: "░", symbolX + 2: "▒", symbolX + 4: "▓", symbolX + 6: "█"},
}
var colors = map[int]string{
	15: "R",
	17: "G",
	19: "B",
}

const colorX = 2
const symbolX = 2

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
	for y, line := range symbols {
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

		drawString(3, y, strconv.Itoa(s.color[line]), screen)
		switch line {
		case "R":
			setByKeys(colorX, y, fgRgb(s.color[line], 0, 0, "█"), screen)
		case "G":
			setByKeys(colorX, y, fgRgb(0, s.color[line], 0, "█"), screen)
		case "B":
			setByKeys(colorX, y, fgRgb(0, 0, s.color[line], "█"), screen)
		}

	}
	setByKeys(3, 21, fgRgb(s.color["R"], s.color["G"], s.color["B"], "█"), screen)
	setByKeys(4, 21, fgRgb(s.color["R"], s.color["G"], s.color["B"], "█"), screen)
	setByKeys(5, 21, fgRgb(s.color["R"], s.color["G"], s.color["B"], "█"), screen)
	setByKeys(3, 22, fgRgb(s.color["R"], s.color["G"], s.color["B"], "█"), screen)
	setByKeys(4, 22, fgRgb(s.color["R"], s.color["G"], s.color["B"], "█"), screen)
	setByKeys(5, 22, fgRgb(s.color["R"], s.color["G"], s.color["B"], "█"), screen)

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
