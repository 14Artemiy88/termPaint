package main

import (
	"strconv"
	"strings"
)

const menuWidth = 15
const helpWidth = 37

var symbols = map[int]map[int]string{
	1:  {3: "┌", 5: "├", 7: "┬", 9: "─", 11: "┐"},
	3:  {3: "└", 5: "┤", 7: "┴", 9: "│", 11: "┘", 13: "┼"},
	5:  {3: "┏", 5: "┣", 7: "┳", 9: "━", 11: "┓"},
	7:  {3: "┗", 5: "┫", 7: "┻", 9: "┃", 11: "┛", 13: "╋"},
	9:  {3: "╭", 5: "╮", 7: "╯", 9: "╰"},
	11: {3: "░", 5: "▒", 7: "▓", 9: "█"},
}
var colors = map[int]map[int]string{
	15: {3: "R"},
	17: {3: "G"},
	19: {3: "B"},
}

func drawMenu(s screen, screen [][]string) [][]string {
	clearMenu(s, screen, menuWidth)
	drawSymbolMenu(s, screen)
	drawColorMenu(s, screen)
	str := "Help " + strings.Repeat("─", menuWidth-2-len("Help ")) + "┤"
	drawString(0, s.rows-2, str, screen)
	drawString(1, s.rows-1, "Press Enter", screen)

	return screen
}

func drawSymbolMenu(s screen, screen [][]string) [][]string {
	str := "Symbol " + strings.Repeat("─", menuWidth-2-len("symbol ")) + "┐"
	drawString(0, 0, str, screen)
	for y, line := range symbols {
		for x, symbol := range line {
			screen[y][x] = fgRgb(s.color["R"], s.color["G"], s.color["B"], symbol)
		}
	}

	return screen
}

func drawColorMenu(s screen, screen [][]string) [][]string {
	str := "Color " + strings.Repeat("─", menuWidth-2-len("Color ")) + "┤"
	drawString(0, 13, str, screen)
	for y, line := range colors {
		for x, color := range line {
			drawString(x, y, strconv.Itoa(s.color[color]), screen)
			switch color {
			case "R":
				screen[y][x] = fgRgb(s.color[color], 0, 0, "█")
			case "G":
				screen[y][x] = fgRgb(0, s.color[color], 0, "█")
			case "B":
				screen[y][x] = fgRgb(0, 0, s.color[color], "█")
			}
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
		screen[Y][X+2+k] = i
	}

	return screen
}

func clearMenu(s screen, screen [][]string, width int) [][]string {
	for i := 0; i < s.rows; i++ {
		for j := 0; j < width; j++ {
			screen[i][j] = " "
		}
		screen[i][width] = "│"
	}

	return screen
}
