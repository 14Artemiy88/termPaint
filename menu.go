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
	for y, line := range symbols {
		for x, symbol := range line {
			board[y][x] = fgRgb(s.color.R, s.color.G, s.color.B, symbol)
		}
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
