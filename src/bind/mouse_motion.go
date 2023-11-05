package bind

import (
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/menu"
	tea "github.com/charmbracelet/bubbletea"
)

func mouseMotion(msg tea.MouseMsg, s Screen) {
	var xMin int
	switch menu.Type {
	case menu.SymbolColor:
		xMin = menu.SymbolColorWidth
	case menu.File:
		xMin = menu.FileListWidth
	case menu.Help:
		xMin = menu.HelpWidth
	case menu.Shape:
		xMin = menu.ShapeWidth
	case menu.Line:
		xMin = menu.LineWidth
	case menu.Config:
		xMin = menu.ConfigWidth
	default:
		xMin = 0
	}

	if msg.X <= xMin {
		switch menu.Type {
		case menu.SymbolColor:
			onMenu(s, msg)
		case menu.File:
			onFile(msg)
		case menu.Shape:
			onShape(msg)
		case menu.Line:
			onLine(msg)
		default:
			cursor.CC.Brush = cursor.Empty
		}
	} else {
		cursor.CC.Symbol = cursor.CC.Store.Symbol
		cursor.CC.Brush = cursor.CC.Store.Brush
	}

	if msg.X >= xMin && msg.X < s.GetWidth() {
		cursor.CC.X = msg.X
	}
	if msg.Y >= 0 && msg.Y < s.GetHeight() {
		cursor.CC.Y = msg.Y
	}
}

func onLine(msg tea.MouseMsg) {
	cursor.CC.Brush = cursor.Empty
	if _, ok := menu.LineList[msg.Y]; ok {
		cursor.CC.Brush = cursor.Pointer
	}
}

func onShape(msg tea.MouseMsg) {
	cursor.CC.Brush = cursor.Empty
	if _, ok := menu.ShapeList[msg.Y]; ok {
		cursor.CC.Brush = cursor.Pointer
	}
}

func onFile(msg tea.MouseMsg) {
	cursor.CC.Brush = cursor.Empty
	if file, ok := menu.FileList[msg.Y]; ok {
		cursor.CC.Brush = cursor.Pointer
		menu.FilePath = file
	} else {
		menu.FilePath = ""
	}
}

func onMenu(s Screen, msg tea.MouseMsg) {
	cursor.CC.Brush = cursor.Empty
	_, okSymbol := s.GetConfig().Symbols[msg.Y][msg.X]
	c, okColor := menu.Colors[msg.Y]
	if okColor && msg.X < menu.SymbolColorWidth {
		menu.Input.Lock = true
		menu.Input.Color = c
		cursor.CC.Brush = cursor.Pointer
	} else {
		menu.Input.Lock = false
		menu.Input.Value = ""
	}
	if !okSymbol && !okColor {
		cursor.CC.X = menu.SymbolColorWidth + 1
	}
}
