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
	drawSymbolMenu(s, screen)
	drawColorMenu(s, screen)
	str := "Help " + strings.Repeat("─", menuWidth-len("Help ")) + "┤"
	drawString(0, s.rows-2, str, screen)
	drawString(1, s.rows-1, "Press Enter", screen)

	return screen
}

func drawSymbolMenu(s *screen, screen [][]string) [][]string {
	str := "Symbol " + strings.Repeat("─", menuWidth-len("symbol ")) + "┐"
	drawString(0, 0, str, screen)
	for y, line := range symbols {
		for x, symbol := range line {
			screen[y][x] = fgRgb(s.color["R"], s.color["G"], s.color["B"], symbol)
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
			screen[y][colorX] = fgRgb(s.color[line], 0, 0, "█")
		case "G":
			screen[y][colorX] = fgRgb(0, s.color[line], 0, "█")
		case "B":
			screen[y][colorX] = fgRgb(0, 0, s.color[line], "█")
		}

	}
	screen[21][3] = fgRgb(s.color["R"], s.color["G"], s.color["B"], "█")
	screen[21][4] = fgRgb(s.color["R"], s.color["G"], s.color["B"], "█")
	screen[21][5] = fgRgb(s.color["R"], s.color["G"], s.color["B"], "█")
	screen[22][3] = fgRgb(s.color["R"], s.color["G"], s.color["B"], "█")
	screen[22][4] = fgRgb(s.color["R"], s.color["G"], s.color["B"], "█")
	screen[22][5] = fgRgb(s.color["R"], s.color["G"], s.color["B"], "█")

	return screen
}

func drawString(X int, Y int, val string, screen [][]string) [][]string {
	str := strings.Split(val, "")
	for k, i := range str {
		screen[Y][X+k] = i
	}

	return screen
}

func clearMenu(s *screen, screen [][]string, width int) [][]string {
	for i := 0; i < s.rows; i++ {
		for j := 0; j < width; j++ {
			screen[i][j] = " "
		}
		screen[i][width] = "│"
	}

	return screen
}
