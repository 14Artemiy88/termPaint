package screen

import (
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"strconv"
)

func KeyBind(msg tea.KeyMsg, s *Screen) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyCtrlC, tea.KeyEsc:
		return s, tea.Quit

	// file
	case tea.KeyCtrlO, tea.KeyF3:
		if s.MenuType == file {
			s.MenuType = None
		} else {
			s.MenuType = file
		}

	// menu
	case tea.KeyTab, tea.KeyF2:
		if s.MenuType == symbolColor {
			s.MenuType = None
		} else {
			s.MenuType = symbolColor
		}

	// help
	case tea.KeyEnter, tea.KeyCtrlH, tea.KeyF1:
		if s.MenuType == help {
			s.MenuType = None
		} else {
			s.MenuType = help
		}

	// shape
	case tea.KeyF4:
		if s.MenuType == shape {
			s.MenuType = None
		} else {
			s.MenuType = shape
		}

	// save
	case tea.KeyCtrlS:
		s.Save = true
		s.MenuType = None
		s.Messages = []Message{}

	// del file
	case tea.KeyDelete:
		if len(s.File) > 0 {
			_ = os.Remove(s.File)
		}

	// set cursor or color
	case tea.KeyRunes:
		if s.MenuType == symbolColor && s.InputLock {
			if _, err := strconv.Atoi(string(msg.Runes)); err == nil {
				s.Input += string(msg.Runes)
			} else {
				s.SetMessage(err.Error())
			}
		} else {
			s.Cursor.Store.Symbol = string(msg.Runes)
			s.Cursor.Symbol = string(msg.Runes)
		}
	}

	return s, nil
}
