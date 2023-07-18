package src

import (
	tea "github.com/charmbracelet/bubbletea"
)

func mouseMotion(msg tea.MouseMsg, s *Screen) {
	xMin := 0
	if s.ShowMenu {
		xMin = menuWidth
	}
	if s.ShowHelp {
		xMin = helpWidth
	}
	if s.ShowFile {
		xMin = s.FileListWidth
	}
	if s.ShowMenu && msg.X <= xMin {
		onMenu(msg, s)
	} else if s.ShowFile && msg.X <= xMin {
		onFile(msg, s)

	} else if msg.X <= xMin {
		s.X = xMin + 1
		s.Cursor = " "
	} else {
		s.Cursor = s.CursorStore
	}

	if msg.X > xMin && msg.X < s.Columns {
		s.X = msg.X
	}
	if msg.Y > 0 && msg.Y < s.Rows {
		s.Y = msg.Y
	}
}

func onFile(msg tea.MouseMsg, s *Screen) {
	if file, ok := s.FileList[msg.Y]; ok {
		s.X = 0
		s.Cursor = FgRgb(Cfg.PointerColor["r"], Cfg.PointerColor["g"], Cfg.PointerColor["b"], Cfg.Pointer)
		s.File = file
	} else {
		s.Cursor = " "
		s.File = ""
	}
}

func onMenu(msg tea.MouseMsg, s *Screen) {
	s.Cursor = " "
	_, okSymbol := Cfg.Symbols[msg.Y][msg.X]
	color, okColor := colors[msg.Y]
	//if okSymbol {
	//	s.X = msg.X - 1
	//	s.Cursor = FgRgb(Cfg.PointerColor["r"], Cfg.PointerColor["g"], Cfg.PointerColor["b"], Cfg.Pointer)
	//}
	if okColor && msg.X < menuWidth {
		s.InputLock = true
		s.InputColor = color
		s.X = 0
		s.Cursor = FgRgb(Cfg.PointerColor["r"], Cfg.PointerColor["g"], Cfg.PointerColor["b"], Cfg.Pointer)
	} else {
		s.InputLock = false
		if len(s.Input) > 0 {
			s.Color[s.InputColor] = setColor(s.Input)
		}
		s.Input = ""
	}
	if !okSymbol && !okColor {
		s.X = menuWidth + 1
		s.Cursor = " "
	}
}
