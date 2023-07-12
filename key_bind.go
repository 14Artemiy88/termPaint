package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"strconv"
)

func keyBind(msg tea.KeyMsg, s *screen) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyCtrlC, tea.KeyEsc:
		return s, tea.Quit

	case tea.KeyCtrlO:
		s.showHelp = false
		s.showMenu = false
		s.showFile = !s.showFile

	case tea.KeyTab:
		s.showHelp = false
		s.showFile = false
		s.showMenu = !s.showMenu

	case tea.KeyEnter:
		s.showMenu = false
		s.showFile = false
		s.showHelp = !s.showHelp

	case tea.KeyCtrlS:
		s.save = true
		s.showMenu = false
		s.showHelp = false
		s.showFile = false

	case tea.KeyDelete:
		if len(s.file) > 0 {
			_ = os.Remove(s.file)
		}

	case tea.KeyRunes:
		if s.showMenu && s.inputLock {
			if _, err := strconv.Atoi(string(msg.Runes)); err == nil {
				s.input += string(msg.Runes)
			}
		} else {
			s.cursorStore = string(msg.Runes)
			s.cursor = string(msg.Runes)
		}
	}

	return s, nil
}
