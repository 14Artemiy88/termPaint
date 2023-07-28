package screen

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
		Y:     1,
		title: "Keys",
		end:   "┐",
		item: []string{
			"ESC      - Exit",
			"Tab      - Menu",
			"Ctrl+S   - Save in txt file",
			"Ctrl+O   - Load Image",
			"Enter    - Show this help menu",
			"Any char - Set this Symbol",
		},
	},
	{
		Y:     9,
		title: "Mouse",
		end:   "┤",
		item: []string{
			"Left   - Draw",
			"Right  - Erase",
			"Middle - Clear Screen",
		},
	},
	{
		Y:     14,
		title: "Symbol",
		end:   "┤",
		item: []string{
			"Click to select Symbol",
		},
	},
	{
		Y:     17,
		title: "Color",
		end:   "┤",
		item: []string{
			"Scroll      - Decrease/increase",
			"Click       - Set 0/255",
			"Press [0-9] - Set color",
		},
	},
	{
		Y:     22,
		title: "File",
		end:   "┤",
		item: []string{
			"Left   - Click to select file",
			"Delete - Press to delete file",
		},
	},
}

const HelpWidth = 37

func DrawHelpMenu(screen [][]string) [][]string {
	ClearMenu(screen, HelpWidth)
	for _, mi := range menu {
		mi.DrawMenuItem(screen)
	}

	return screen
}

func (m menuItem) DrawMenuItem(screen [][]string) [][]string {
	var str string
	str = m.title + " " + strings.Repeat("─", HelpWidth-2-len(m.title)) + m.end
	DrawString(1, m.Y, str, screen)
	for k, i := range m.item {
		str = i
		DrawString(3, m.Y+1+k, str, screen)
	}

	return screen
}
