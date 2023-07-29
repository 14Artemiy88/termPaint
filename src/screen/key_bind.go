package screen

import (
	"github.com/14Artemiy88/termPaint/src/cursor"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"strconv"
)

func KeyBind(msg tea.KeyMsg, s *Screen) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyCtrlC, tea.KeyEsc:
		return s, tea.Quit

	// help
	case tea.KeyEnter, tea.KeyCtrlH, tea.KeyF1:
		if MenuType == help {
			MenuType = None
		} else {
			MenuType = help
		}

	// menu
	case tea.KeyTab, tea.KeyF2:
		if MenuType == symbolColor {
			MenuType = None
		} else {
			MenuType = symbolColor
		}

	// file
	case tea.KeyCtrlO, tea.KeyF3:
		if MenuType == File {
			MenuType = None
		} else {
			MenuType = File
		}

	// shape
	case tea.KeyF4:
		if MenuType == shape {
			MenuType = None
		} else {
			MenuType = shape
		}

	// line
	case tea.KeyF5:
		if MenuType == line {
			MenuType = None
		} else {
			MenuType = line
		}

	// save
	case tea.KeyCtrlS:
		s.Save = true
		MenuType = None
		Msg = []Message{}

	// del file
	case tea.KeyDelete:
		if len(FilePath) > 0 {
			_ = os.Remove(FilePath)
		}

	// set cursor or color
	case tea.KeyRunes:
		if MenuType == symbolColor && input.lock {
			if _, err := strconv.Atoi(string(msg.Runes)); err == nil {
				input.value += string(msg.Runes)
			} else {
				SetMessage(err.Error())
			}
		} else {
			cursor.CC.SetCursor(string(msg.Runes))
		}
	}

	return s, nil
}
