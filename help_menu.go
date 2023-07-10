package main

import "strings"

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

func drawHelpMenu(s screen, board [][]string) [][]string {
	clearMenu(s, board, helpWidth)
	for _, mi := range menu {
		mi.drawMenuItem(board)
	}

	return board
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
