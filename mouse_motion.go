package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func mouseMotion(msg tea.MouseMsg, s *screen) {
	xMin := 0
	if s.showMenu {
		xMin = menuWidth
	}
	if s.showHelp {
		xMin = helpWidth
	}
	if s.showFile {
		xMin = s.fileListWidth
	}
	if s.showMenu && msg.X <= xMin {
		onMenu(msg, s)
	} else if s.showFile && msg.X <= xMin {
		onFile(msg, s)

	} else if msg.X <= xMin {
		s.X = xMin + 1
		s.cursor = " "
	} else {
		s.cursor = s.cursorStore
	}

	if msg.X > xMin && msg.X < s.columns {
		s.X = msg.X
	}
	if msg.Y > 0 && msg.Y < s.rows {
		s.Y = msg.Y
	}
}

func onFile(msg tea.MouseMsg, s *screen) {
	if file, ok := s.fileList[msg.Y]; ok {
		s.X = 0
		s.cursor = fgRgb(170, 170, 170, pointer)
		s.file = file
	} else {
		s.cursor = " "
		s.file = ""
	}
}

func onMenu(msg tea.MouseMsg, s *screen) {
	s.cursor = " "
	_, okS := symbols[msg.Y][msg.X]
	color, okC := colors[msg.Y][msg.X]
	if okS {
		s.X = msg.X - 1
		s.cursor = fgRgb(170, 170, 170, pointer)
	}
	if okC {
		s.inputLock = true
		s.inputColor = color
		s.X = msg.X - 2
		s.cursor = fgRgb(170, 170, 170, pointer)
	} else {
		s.inputLock = false
		if len(s.input) > 0 {
			s.color[s.inputColor] = setColor(s.input)
		}
		s.input = ""
	}
	if !okS && !okC {
		s.X = menuWidth + 1
		s.cursor = " "
	}
}
