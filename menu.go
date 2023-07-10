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

func drawMenu(s screen, board [][]string) [][]string {
	clearMenu(s, board, menuWidth)
	drawSymbolMenu(s, board)
	drawColorMenu(s, board)
	str := "Help " + strings.Repeat("─", menuWidth-2-len("Help ")) + "┤"
	drawString(0, s.rows-2, str, board)
	drawString(1, s.rows-1, "Press Enter", board)

	return board
}

func drawSymbolMenu(s screen, board [][]string) [][]string {
	str := "Symbol " + strings.Repeat("─", menuWidth-2-len("symbol ")) + "┐"
	drawString(0, 0, str, board)
	for y, line := range symbols {
		for x, symbol := range line {
			board[y][x] = fgRgb(s.color["R"], s.color["G"], s.color["B"], symbol)
		}
	}

	return board
}

func drawColorMenu(s screen, board [][]string) [][]string {
	str := "Color " + strings.Repeat("─", menuWidth-2-len("Color ")) + "┤"
	drawString(0, 13, str, board)
	for y, line := range colors {
		for x, color := range line {
			drawString(x, y, strconv.Itoa(s.color[color]), board)
			switch color {
			case "R":
				board[y][x] = fgRgb(s.color[color], 0, 0, "█")
			case "G":
				board[y][x] = fgRgb(0, s.color[color], 0, "█")
			case "B":
				board[y][x] = fgRgb(0, 0, s.color[color], "█")
			}
		}
	}
	board[21][3] = fgRgb(s.color["R"], s.color["G"], s.color["B"], "█")
	board[21][4] = fgRgb(s.color["R"], s.color["G"], s.color["B"], "█")
	board[21][5] = fgRgb(s.color["R"], s.color["G"], s.color["B"], "█")
	board[22][3] = fgRgb(s.color["R"], s.color["G"], s.color["B"], "█")
	board[22][4] = fgRgb(s.color["R"], s.color["G"], s.color["B"], "█")
	board[22][5] = fgRgb(s.color["R"], s.color["G"], s.color["B"], "█")

	return board
}

func drawString(X int, Y int, val string, board [][]string) [][]string {
	r := strings.Split(val, "")
	for k, i := range r {
		board[Y][X+2+k] = i
	}

	return board
}

func clearMenu(s screen, board [][]string, width int) [][]string {
	for i := 0; i < s.rows; i++ {
		for j := 0; j < width; j++ {
			board[i][j] = " "
		}
		board[i][width] = "│"
	}

	return board
}
