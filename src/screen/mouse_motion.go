package screen

import (
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/config"
	tea "github.com/charmbracelet/bubbletea"
)

func mouseMotion(msg tea.MouseMsg, s *Screen) {
	xMin := 0
	if s.ShowMenu {
		xMin = MenuWidth
	}
	if s.ShowHelp {
		xMin = HelpWidth
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
	} else {
		s.Cursor.Symbol = s.Cursor.Store.Symbol
		s.Cursor.Brush = s.Cursor.Store.Brush
	}

	if msg.X > xMin && msg.X < s.Columns {
		s.X = msg.X
	}
	if msg.Y > 0 && msg.Y < s.Rows {
		s.Y = msg.Y
	}
}

func onFile(msg tea.MouseMsg, s *Screen) {
	s.Cursor.Symbol = emptyCursor
	s.Cursor.Brush = Dot
	if file, ok := s.FileList[msg.Y]; ok {
		s.Cursor.Brush = Pointer
		s.File = file
	} else {
		s.File = ""
	}
}

func onMenu(msg tea.MouseMsg, s *Screen) {
	s.Cursor.Symbol = emptyCursor
	s.Cursor.Brush = Dot
	_, okSymbol := config.Cfg.Symbols[msg.Y][msg.X]
	c, okColor := Colors[msg.Y]
	if okColor && msg.X < MenuWidth {
		s.InputLock = true
		s.InputColor = c
		s.Cursor.Brush = Pointer
	} else {
		s.InputLock = false
		if len(s.Input) > 0 {
			s.Cursor.Color[s.InputColor] = color.SetColor(s.Input)
		}
		s.Input = ""
	}
	if !okSymbol && !okColor {
		s.X = MenuWidth + 1
	}
}
