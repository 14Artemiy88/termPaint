package main

import (
	"strconv"
	"strings"
)

const menuWidth = 15
const helpWidth = 37

var symbols = map[string]pixel{
	"3,1": {X: 3, Y: 1, symbol: "┌"}, "5,1": {X: 5, Y: 1, symbol: "├"}, "7,1": {X: 7, Y: 1, symbol: "┬"}, "9,1": {X: 9, Y: 1, symbol: "─"}, "11,1": {X: 11, Y: 1, symbol: "┐"},
	"3,3": {X: 3, Y: 3, symbol: "└"}, "5,3": {X: 5, Y: 3, symbol: "┤"}, "7,3": {X: 7, Y: 3, symbol: "┴"}, "9,3": {X: 9, Y: 3, symbol: "│"}, "11,3": {X: 11, Y: 3, symbol: "┘"}, "13,3": {X: 13, Y: 3, symbol: "┼"},
	"3,5": {X: 3, Y: 5, symbol: "┏"}, "5,5": {X: 5, Y: 5, symbol: "┣"}, "7,5": {X: 7, Y: 5, symbol: "┳"}, "9,5": {X: 9, Y: 5, symbol: "━"}, "11,5": {X: 11, Y: 5, symbol: "┓"},
	"3,7": {X: 3, Y: 7, symbol: "┗"}, "5,7": {X: 5, Y: 7, symbol: "┫"}, "7,7": {X: 7, Y: 7, symbol: "┻"}, "9,7": {X: 9, Y: 7, symbol: "┃"}, "11,7": {X: 11, Y: 7, symbol: "┛"}, "13,7": {X: 13, Y: 7, symbol: "╋"},
	"3,9": {X: 3, Y: 9, symbol: "╭"}, "5,9": {X: 5, Y: 9, symbol: "╮"}, "7,9": {X: 7, Y: 9, symbol: "╯"}, "9,9": {X: 9, Y: 9, symbol: "╰"},
	"3,11": {X: 3, Y: 11, symbol: "░"}, "5,11": {X: 5, Y: 11, symbol: "▒"}, "7,11": {X: 7, Y: 11, symbol: "▓"}, "9,11": {X: 9, Y: 11, symbol: "█"},
}

var colors = map[string]pixel{
	"3,15": {X: 3, Y: 15, symbol: "R"},
	"3,17": {X: 3, Y: 17, symbol: "G"},
	"3,19": {X: 3, Y: 19, symbol: "B"},
}

type menuItem struct {
	Y     int
	width int
	title string
	item  []string
	end   string
}

var menu = []menuItem{
	{
		Y:     0,
		title: "Keys",
		width: helpWidth,
		end:   "┐",
		item: []string{
			"ESC      - Exit",
			"Tab      - Menu",
			"Ctrl+S   - Save in txt file",
			"Enter    - Show this help menu",
			"Any char - Set this symbol",
		},
	},
	{
		Y:     7,
		title: "Mouse",
		width: helpWidth,
		end:   "┤",
		item: []string{
			"Left   - Draw",
			"Right  - Erase",
			"Middle - Clear screen",
		},
	},
	{
		Y:     12,
		title: "Menu",
		width: helpWidth,
		end:   "┤",
		item: []string{
			"Symbols - Click to select symbol",
			"Colors  - Hover and scroll to",
			"          change color",
			"          or input number",
		},
	},
}

func (m menuItem) drawMenuItem(board [][]string) [][]string {
	var str string
	str = m.title + " " + strings.Repeat("─", helpWidth-3-len(m.title)) + m.end
	drawString(0, m.Y, str, board)
	for k, i := range m.item {
		str = i
		drawString(1, m.Y+1+k, str, board)
	}

	return board
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
	for _, symbol := range symbols {
		board[symbol.Y][symbol.X] = fgRgb(s.color.R, s.color.G, s.color.B, symbol.symbol)
	}

	return board
}

func drawColorMenu(s screen, board [][]string) [][]string {
	str := "Color " + strings.Repeat("─", menuWidth-2-len("Color ")) + "┤"
	drawString(0, 13, str, board)
	for _, c := range colors {
		switch c.symbol {
		case "R":
			drawString(c.X, c.Y, strconv.Itoa(s.color.R), board)
			board[c.Y][c.X] = fgRgb(s.color.R, 0, 0, "█")
		case "G":
			drawString(c.X, c.Y, strconv.Itoa(s.color.G), board)
			board[c.Y][c.X] = fgRgb(0, s.color.G, 0, "█")
		case "B":
			drawString(c.X, c.Y, strconv.Itoa(s.color.B), board)
			board[c.Y][c.X] = fgRgb(0, 0, s.color.B, "█")
		}

	}
	board[21][3] = fgRgb(s.color.R, s.color.G, s.color.B, "█")
	board[21][4] = fgRgb(s.color.R, s.color.G, s.color.B, "█")
	board[21][5] = fgRgb(s.color.R, s.color.G, s.color.B, "█")
	board[22][3] = fgRgb(s.color.R, s.color.G, s.color.B, "█")
	board[22][4] = fgRgb(s.color.R, s.color.G, s.color.B, "█")
	board[22][5] = fgRgb(s.color.R, s.color.G, s.color.B, "█")

	return board
}

func drawString(X int, Y int, val string, board [][]string) [][]string {
	r := strings.Split(val, "")
	for k, i := range r {
		board[Y][X+2+k] = i
	}

	return board
}

func drawHelpMenu(s screen, board [][]string) [][]string {
	clearMenu(s, board, helpWidth)
	for _, mi := range menu {
		mi.drawMenuItem(board)
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
