package screen

import (
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/size"
	tea "github.com/charmbracelet/bubbletea"
)

func mouseMotion(msg tea.MouseMsg) {
	var xMin int
	switch MenuType {
	case symbolColor:
		xMin = MenuSymbolColorWidth
	case file:
		xMin = FileListWidth
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
		switch MenuType {
		case symbolColor:
			onMenu(msg)
		case file:
			onFile(msg)
		case shape:
			onShape(msg)
		case line:
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
	if _, ok := menuLineList[msg.Y]; ok {
		cursor.CC.Brush = cursor.Pointer
	}
}

func onShape(msg tea.MouseMsg) {
	cursor.CC.Brush = cursor.Empty
	if _, ok := shapeList[msg.Y]; ok {
		cursor.CC.Brush = cursor.Pointer
	}
}

func onFile(msg tea.MouseMsg) {
	cursor.CC.Brush = cursor.Empty
	if file, ok := FileList[msg.Y]; ok {
		cursor.CC.Brush = cursor.Pointer
		FilePath = file
	} else {
		FilePath = ""
	}
}

func onMenu(msg tea.MouseMsg) {
	cursor.CC.Brush = cursor.Empty
	_, okSymbol := config.Cfg.Symbols[msg.Y][msg.X]
	c, okColor := Colors[msg.Y]
	if okColor && msg.X < MenuSymbolColorWidth {
		input.lock = true
		input.color = c
		cursor.CC.Brush = cursor.Pointer
	} else {
		input.lock = false
		if len(input.value) > 0 {
			cursor.CC.Color[input.color] = color.SetColor(input.value)
		}
		input.value = ""
	}
	if !okSymbol && !okColor {
		cursor.CC.X = MenuSymbolColorWidth + 1
	}
}
