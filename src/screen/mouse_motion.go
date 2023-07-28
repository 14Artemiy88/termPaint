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
			onShape(msg)
		case line:
			onLine(msg)
		default:
			CC.Brush = Empty
		}
	} else {
		CC.Symbol = CC.Store.Symbol
		CC.Brush = CC.Store.Brush
	}

	if msg.X > xMin && msg.X < s.Columns {
		s.X = msg.X
	}
	if msg.Y > 0 && msg.Y < s.Rows {
		s.Y = msg.Y
	}
}

func onLine(msg tea.MouseMsg) {
	CC.Brush = Empty
	if _, ok := menuLineList[msg.Y]; ok {
		CC.Brush = Pointer
	}
}

func onShape(msg tea.MouseMsg) {
	CC.Brush = Empty
	if _, ok := shapeList[msg.Y]; ok {
		CC.Brush = Pointer
	}
}

func onFile(msg tea.MouseMsg, s *Screen) {
	CC.Brush = Empty
	if file, ok := s.FileList[msg.Y]; ok {
		CC.Brush = Pointer
		s.File = file
	} else {
		s.File = ""
	}
}

func onMenu(msg tea.MouseMsg, s *Screen) {
	CC.Brush = Empty
	_, okSymbol := config.Cfg.Symbols[msg.Y][msg.X]
	c, okColor := Colors[msg.Y]
	if okColor && msg.X < MenuSymbolColorWidth {
		input.lock = true
		input.color = c
		CC.Brush = Pointer
	} else {
		input.lock = false
		if len(input.value) > 0 {
			CC.Color[input.color] = color.SetColor(input.value)
		}
		input.value = ""
	}
	if !okSymbol && !okColor {
		s.X = MenuSymbolColorWidth + 1
	}
}
