package screen

import (
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/config"
	tea "github.com/charmbracelet/bubbletea"
)

func MouseBind(msg tea.MouseMsg, s *Screen) {
	switch msg.Type {
	case tea.MouseMotion:
		mouseMotion(msg, s)

	case tea.MouseLeft:
		mouseLeft(msg, s)

	case tea.MouseRight:
		s.Pixels = append(s.Pixels, Pixel{X: msg.X, Y: msg.Y, Symbol: " "})

	case tea.MouseMiddle:
		s.Pixels = []Pixel{}

	case tea.MouseWheelDown:
		if c, ok := Colors[msg.Y]; ok {
			s.Cursor.Color[c] = color.Decrease(s.Cursor.Color[c])
		}
		if s.Cursor.Symbol != emptyCursor && s.Cursor.Symbol != config.Cfg.Pointer && s.Cursor.Width > 1 {
			s.Cursor.Width--
		}

	case tea.MouseWheelUp:
		if c, ok := Colors[msg.Y]; ok {
			s.Cursor.Color[c] = color.Increase(s.Cursor.Color[c])
		}
		if s.Cursor.Symbol != emptyCursor && s.Cursor.Symbol != config.Cfg.Pointer {
			s.Cursor.Width++
		}
	}
}
