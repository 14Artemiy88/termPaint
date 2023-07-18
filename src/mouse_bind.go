package src

import (
	tea "github.com/charmbracelet/bubbletea"
)

func mouseBind(msg tea.MouseMsg, s *Screen) {
	switch msg.Type {
	case tea.MouseMotion:
		mouseMotion(msg, s)

	case tea.MouseLeft:
		mouseLeft(msg, s)

	case tea.MouseRight:
		s.Pixels = append(s.Pixels, pixel{X: msg.X, Y: msg.Y, symbol: " "})

	case tea.MouseMiddle:
		s.Pixels = []pixel{}

	case tea.MouseWheelDown:
		if color, ok := colors[msg.Y]; ok {
			s.Color[color] = Decrease(s.Color[color])
		}

	case tea.MouseWheelUp:
		if color, ok := colors[msg.Y]; ok {
			s.Color[color] = Increase(s.Color[color])
		}
	}
}
