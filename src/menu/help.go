package menu

import (
	"strings"

	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/utils"
)

type menuItem struct {
	Y     int
	title string
	item  []map[string]string
	end   string
}

var menu = []menuItem{
	{
		Y:     1,
		title: "Keys",
		end:   "┐",
		item: []map[string]string{
			{"key": "ESC     ", "text": "Exit"},
			{"key": "Tab     ", "text": "Menu"},
			{"key": "Ctrl+S  ", "text": "Save in txt file"},
			{"key": "Ctrl+O  ", "text": "Load Image"},
			{"key": "Ctrl+H  ", "text": "Show this help menu"},
			{"key": "Any char", "text": "Set as a Symbol"},
		},
	},
	{
		Y:     10,
		title: "Mouse",
		end:   "┤",
		item: []map[string]string{
			{"key": "Left  ", "text": "Draw"},
			{"key": "Right ", "text": "Erase"},
			{"key": "Middle", "text": "Clear Screen"},
		},
	},
	{
		Y:     16,
		title: "Symbol",
		end:   "┤",
		item: []map[string]string{
			{"key": "", "text": "Click to select Symbol"},
		},
	},
	{
		Y:     20,
		title: "Color",
		end:   "┤",
		item: []map[string]string{
			{"key": "Scroll     ", "text": "Decrease/increase"},
			{"key": "Click      ", "text": "Set 0/255"},
			{"key": "Press [0-9]", "text": "Set color"},
		},
	},
	{
		Y:     26,
		title: "FilePath",
		end:   "┤",
		item: []map[string]string{
			{"key": "Left  ", "text": "Click to select file"},
			{"key": "Delete", "text": "Press to delete file"},
		},
	},
}

const HelpWidth = 40

func drawHelpMenu(s Screen) {
	screen := s.GetPixels()
	ClearMenu(s, screen, HelpWidth)

	for _, mi := range menu {
		mi.DrawMenuItem(screen)
	}
}

func (m menuItem) DrawMenuItem(screen [][]string) [][]string {
	str := strings.Repeat("─", HelpWidth-2-len(m.title)) + m.end
	utils.DrawString(1, m.Y, m.title, pixel.GetConstColor("yellow"), screen)
	utils.DrawString(len(m.title)+2, m.Y, str, pixel.GetConstColor("gray"), screen)

	green := pixel.GetConstColor("green")
	white := pixel.GetConstColor("White")

	for k, str := range m.item {
		lenKey := len(str["key"])
		if lenKey > 0 {
			utils.DrawString(3, m.Y+2+k, str["key"], green, screen)
		}

		utils.DrawString(16, m.Y+2+k, str["text"], white, screen)
	}

	return screen
}
