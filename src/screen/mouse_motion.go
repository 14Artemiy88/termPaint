package screen

import (
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/menu"
	"github.com/14Artemiy88/termPaint/src/size"
	tea "github.com/charmbracelet/bubbletea"
)

func mouseMotion(msg tea.MouseMsg) {
	var xMin int
	switch menu.MenuType {
	case menu.SymbolColor:
		xMin = menu.MenuSymbolColorWidth
	case menu.File:
		xMin = menu.FileListWidth
	case menu.Help:
		xMin = menu.HelpWidth
	case menu.Shape:
		xMin = menu.MenuShapeWidth
	case menu.Line:
		xMin = menu.MenuLineWidth
	default:
		xMin = 0
	}

	if msg.X <= xMin {
		switch menu.MenuType {
		case menu.SymbolColor:
			onMenu(msg)
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

	if msg.X > xMin && msg.X < size.Size.Columns {
		cursor.CC.X = msg.X
	}
	if msg.Y > 0 && msg.Y < size.Size.Rows {
		cursor.CC.Y = msg.Y
	}
}

func onLine(msg tea.MouseMsg) {
	cursor.CC.Brush = cursor.Empty
	if _, ok := menu.MenuLineList[msg.Y]; ok {
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

func onMenu(msg tea.MouseMsg) {
	cursor.CC.Brush = cursor.Empty
	_, okSymbol := config.Cfg.Symbols[msg.Y][msg.X]
	c, okColor := menu.Colors[msg.Y]
	if okColor && msg.X < menu.MenuSymbolColorWidth {
		menu.Input.Lock = true
		menu.Input.Color = c
		cursor.CC.Brush = cursor.Pointer
	} else {
		menu.Input.Lock = false
		if len(menu.Input.Value) > 0 {
			cursor.CC.Color[menu.Input.Color] = color.SetColor(menu.Input.Value)
		}
		menu.Input.Value = ""
	}
	if !okSymbol && !okColor {
		cursor.CC.X = menu.MenuSymbolColorWidth + 1
	}
}
