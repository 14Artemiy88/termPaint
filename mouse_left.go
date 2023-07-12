package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func mouseLeft(msg tea.MouseMsg, s *screen) {
	if s.showMenu && msg.X < menuWidth {
		if symbol, ok := cfg.Symbols[msg.Y][msg.X]; ok {
			s.cursorStore = symbol
			s.cursor = symbol
		}
		if color, ok := colors[msg.Y]; ok {
			s.color[color] = minMsxColor(s.color[color])
		}
	} else if s.showFile && msg.X < s.fileListWidth {
		if file, ok := s.fileList[msg.Y]; ok {
			content, err := os.ReadFile(s.dir + file)
			if err != nil {
				s.dir += file
			} else {
				s.showFile = false
			}
			s.loadImage(string(content))
		}
	} else {
		s.pixels = append(s.pixels, pixel{X: msg.X, Y: msg.Y, symbol: fgRgb(s.color["r"], s.color["g"], s.color["b"], s.cursor)})
	}
}
