package main

import "strings"

type menuItem struct {
	Y     int
	title string
	item  []string
	end   string
}

var menu = []menuItem{
	{
		Y:     0,
		title: "Keys",
		end:   "┐",
		item: []string{
			"ESC      - Exit",
			"Tab      - Menu",
			"Ctrl+S   - Save in txt file",
			"Ctrl+O   - Load Image",
			"Enter    - Show this help menu",
			"Any char - Set this symbol",
		},
	},
	{
		Y:     8,
		title: "Mouse",
		end:   "┤",
		item: []string{
			"Left   - Draw",
			"Right  - Erase",
			"Middle - Clear screen",
		},
	},
	{
		Y:     13,
		title: "Menu",
		end:   "┤",
		item: []string{
			"Symbols - Click to select symbol",
			"Colors  - Hover and scroll to",
			"          change color",
			"          or input number",
		},
	},
	{
		Y:     19,
		title: "File",
		end:   "┤",
		item: []string{
			"Left   - Click to select file",
			"Delete - Press to delete file",
		},
	},
}

const helpWidth = 37

func drawHelpMenu(s *screen, screen [][]string) [][]string {
	clearMenu(s, screen, helpWidth)
	for _, mi := range menu {
		mi.drawMenuItem(screen)
	}

	return screen
}

func (m menuItem) drawMenuItem(screen [][]string) [][]string {
	var str string
	str = m.title + " " + strings.Repeat("─", helpWidth-3-len(m.title)) + m.end
	drawString(0, m.Y, str, screen)
	for k, i := range m.item {
		str = i
		drawString(1, m.Y+1+k, str, screen)
	}

	return screen
}
