package src

import (
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func mouseLeft(msg tea.MouseMsg, s *Screen) {
	if s.ShowMenu && msg.X < menuWidth {
		if symbol, ok := Cfg.Symbols[msg.Y][msg.X]; ok {
			s.CursorStore = symbol
			s.Cursor = symbol
		}
		if color, ok := colors[msg.Y]; ok {
			s.Color[color] = minMsxColor(s.Color[color])
		}
	} else if s.ShowFile && msg.X < s.FileListWidth {
		if file, ok := s.FileList[msg.Y]; ok {
			content, err := os.ReadFile(s.Dir + file)
			if err != nil {
				s.Dir += file
			} else {
				s.ShowFile = false
			}
			s.loadImage(string(content))
		}
	} else {
		s.Pixels = append(s.Pixels, pixel{X: msg.X, Y: msg.Y, symbol: FgRgb(s.Color["r"], s.Color["g"], s.Color["b"], s.Cursor)})
	}
}
