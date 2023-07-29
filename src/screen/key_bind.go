package screen

import (
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/menu"
	"github.com/14Artemiy88/termPaint/src/message"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"strconv"
)

func KeyBind(msg tea.KeyMsg, s *Screen) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyCtrlC, tea.KeyEsc:
		return s, tea.Quit

	// Help
	case tea.KeyEnter, tea.KeyCtrlH, tea.KeyF1:
		if menu.MenuType == menu.Help {
			menu.MenuType = menu.None
		} else {
			menu.MenuType = menu.Help
		}

	// menu
	case tea.KeyTab, tea.KeyF2:
		if menu.MenuType == menu.SymbolColor {
			menu.MenuType = menu.None
		} else {
			menu.MenuType = menu.SymbolColor
		}

	// file
	case tea.KeyCtrlO, tea.KeyF3:
		if menu.MenuType == menu.File {
			menu.MenuType = menu.None
		} else {
			menu.MenuType = menu.File
		}

	// shape
	case tea.KeyF4:
		if menu.MenuType == menu.Shape {
			menu.MenuType = menu.None
		} else {
			menu.MenuType = menu.Shape
		}

	// line
	case tea.KeyF5:
		if menu.MenuType == menu.Line {
			menu.MenuType = menu.None
		} else {
			menu.MenuType = menu.Line
		}

	// save
	case tea.KeyCtrlS:
		s.Save = true
		menu.MenuType = menu.None
		message.Msg = []message.Message{}

	// del file
	case tea.KeyDelete:
		if len(menu.FilePath) > 0 {
			_ = os.Remove(menu.FilePath)
		}

	// set cursor or color
	case tea.KeyRunes:
		if menu.MenuType == menu.SymbolColor && menu.Input.Lock {
			if _, err := strconv.Atoi(string(msg.Runes)); err == nil {
				menu.Input.Value += string(msg.Runes)
			} else {
				message.SetMessage(err.Error())
			}
		} else {
			cursor.CC.SetCursor(string(msg.Runes))
		}
	}

	return s, nil
}
