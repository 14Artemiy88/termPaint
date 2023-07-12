package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func mouseLeft(msg tea.MouseMsg, s *screen) {
	if s.showMenu && msg.X < menuWidth {
		if symbol, ok := symbols[msg.Y][msg.X]; ok {
			s.cursorStore = symbol
			s.cursor = symbol
		}
		if color, ok := colors[msg.Y]; ok {
			s.color[color] = minMsxColor(s.color[color])
		}
	} else if s.showFile && msg.X < s.fileListWidth {
		s.showFile = false
		if file, ok := s.fileList[msg.Y]; ok {
			content, err := os.ReadFile(file)
			if err != nil {
				s.setMessage(err.Error())
			}
			s.load(string(content))
		}
	} else {
		s.pixels = append(s.pixels, pixel{X: msg.X, Y: msg.Y, symbol: fgRgb(s.color["R"], s.color["G"], s.color["B"], s.cursor)})
	}
}
