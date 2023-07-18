package src

import (
	"strings"
)

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
			"Middle - Clear Screen",
		},
	},
	{
		Y:     13,
		title: "Symbol",
		end:   "┤",
		item: []string{
			"Click to select symbol",
		},
	},
	{
		Y:     16,
		title: "Color",
		end:   "┤",
		item: []string{
			"Scroll      - Decrease/increase",
			"Click       - Set 0/255",
			"Press [0-9] - Set color",
		},
	},
	{
		Y:     21,
		title: "File",
		end:   "┤",
		item: []string{
			"Left   - Click to select file",
			"Delete - Press to delete file",
		},
	},
}

const helpWidth = 37

func drawHelpMenu(s *Screen, screen [][]string) [][]string {
	ClearMenu(s, screen, helpWidth)
	for _, mi := range menu {
		mi.drawMenuItem(screen)
	}

	return screen
}

func (m menuItem) drawMenuItem(screen [][]string) [][]string {
	var str string
	str = m.title + " " + strings.Repeat("─", helpWidth-1-len(m.title)) + m.end
	DrawString(0, m.Y, str, screen)
	for k, i := range m.item {
		str = i
		DrawString(2, m.Y+1+k, str, screen)
	}

	return screen
}
