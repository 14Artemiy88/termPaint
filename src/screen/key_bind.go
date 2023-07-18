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
	case tea.KeyCtrlO:
		s.ShowHelp = false
		s.ShowMenu = false
		s.ShowFile = !s.ShowFile

	// menu
	case tea.KeyTab:
		s.ShowHelp = false
		s.ShowFile = false
		s.ShowMenu = !s.ShowMenu

	// help
	case tea.KeyEnter, tea.KeyCtrlH:
		s.ShowMenu = false
		s.ShowFile = false
		s.ShowHelp = !s.ShowHelp

	// save
	case tea.KeyCtrlS:
		s.Save = true
		s.ShowMenu = false
		s.ShowHelp = false
		s.ShowFile = false
		s.Messages = []Message{}

	// del file
	case tea.KeyDelete:
		if len(s.File) > 0 {
			_ = os.Remove(s.File)
		}

	case tea.KeyCtrlZ:
		s.Cursor.Brush++
		if s.Cursor.Brush > 8 {
			s.Cursor.Brush = 0
		}

	// set cursor or color
	case tea.KeyRunes:
		if s.ShowMenu && s.InputLock {
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
