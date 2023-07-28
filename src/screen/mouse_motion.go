package screen

import (
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/config"
	tea "github.com/charmbracelet/bubbletea"
)

func mouseMotion(msg tea.MouseMsg, s *Screen) {
	var xMin int
	switch s.MenuType {
	case symbolColor:
		xMin = MenuSymbolColorWidth
	case file:
		xMin = s.FileListWidth
	case help:
		xMin = HelpWidth
	case shape:
		xMin = MenuShapeWidth
	case line:
		xMin = MenuLineWidth
	default:
		xMin = 0
	}

	if msg.X <= xMin {
		switch s.MenuType {
		case symbolColor:
			onMenu(msg, s)
		case file:
			onFile(msg, s)
		case shape:
			onShape(msg, s)
		case line:
			onLine(msg, s)
		default:
			s.Cursor.Brush = Empty
		}
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

func onLine(msg tea.MouseMsg, s *Screen) {
	s.Cursor.Brush = Empty
	if _, ok := menuLineList[msg.Y]; ok {
		s.Cursor.Brush = Pointer
	}
}

func onShape(msg tea.MouseMsg, s *Screen) {
	s.Cursor.Brush = Empty
	if _, ok := shapeList[msg.Y]; ok {
		s.Cursor.Brush = Pointer
	}
}

func onFile(msg tea.MouseMsg, s *Screen) {
	s.Cursor.Brush = Empty
	if file, ok := s.FileList[msg.Y]; ok {
		s.Cursor.Brush = Pointer
		s.File = file
	} else {
		s.File = ""
	}
}

func onMenu(msg tea.MouseMsg, s *Screen) {
	s.Cursor.Brush = Empty
	_, okSymbol := config.Cfg.Symbols[msg.Y][msg.X]
	c, okColor := Colors[msg.Y]
	if okColor && msg.X < MenuSymbolColorWidth {
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
		s.X = MenuSymbolColorWidth + 1
	}
}
