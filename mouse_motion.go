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
		s.cursor = fgRgb(cfg.PointerColor["r"], cfg.PointerColor["g"], cfg.PointerColor["b"], cfg.Pointer)
		s.file = file
	} else {
		s.cursor = " "
		s.file = ""
	}
}

func onMenu(msg tea.MouseMsg, s *screen) {
	s.cursor = " "
	_, okSymbol := cfg.Symbols[msg.Y][msg.X]
	color, okColor := colors[msg.Y]
	//if okSymbol {
	//	s.X = msg.X - 1
	//	s.cursor = fgRgb(cfg.PointerColor["r"], cfg.PointerColor["g"], cfg.PointerColor["b"], cfg.Pointer)
	//}
	if okColor && msg.X < menuWidth {
		s.inputLock = true
		s.inputColor = color
		s.X = 0
		s.cursor = fgRgb(cfg.PointerColor["r"], cfg.PointerColor["g"], cfg.PointerColor["b"], cfg.Pointer)
	} else {
		s.inputLock = false
		if len(s.input) > 0 {
			s.color[s.inputColor] = setColor(s.input)
		}
		s.input = ""
	}
	if !okSymbol && !okColor {
		s.X = menuWidth + 1
		s.cursor = " "
	}
}
