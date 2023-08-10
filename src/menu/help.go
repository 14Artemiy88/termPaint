package menu

import (
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/utils"
	"strings"
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
			{"key": "Enter   ", "text": "Show this help menu"},
			{"key": "Any char", "text": "Set this Symbol"},
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
	str = strings.Repeat("─", HelpWidth-2-len(m.title)) + m.end
	utils.DrawString(1, m.Y, m.title, color.Yellow, screen)
	utils.DrawString(len(m.title)+2, m.Y, str, color.Gray, screen)

	for k, str := range m.item {
		lenKey := len(str["key"])
		if lenKey > 0 {
			utils.DrawString(3, m.Y+2+k, str["key"], color.Green, screen)
		}
		utils.DrawString(lenKey+7, m.Y+2+k, str["text"], color.White, screen)
	}

	return screen
}
