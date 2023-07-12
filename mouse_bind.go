package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func mouseBind(msg tea.MouseMsg, s *screen) {
	switch msg.Type {
	case tea.MouseMotion:
		mouseMotion(msg, s)

	case tea.MouseLeft:
		mouseLeft(msg, s)

	case tea.MouseRight:
		s.pixels = append(s.pixels, pixel{X: msg.X, Y: msg.Y, symbol: " "})

	case tea.MouseMiddle:
		s.pixels = []pixel{}

	case tea.MouseWheelDown:
		if color, ok := colors[msg.Y]; ok {
			s.color[color] = decrease(s.color[color])
		}

	case tea.MouseWheelUp:
		if color, ok := colors[msg.Y]; ok {
			s.color[color] = increase(s.color[color])
		}
	}
}
