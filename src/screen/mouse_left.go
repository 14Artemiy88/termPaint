package screen

import (
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/utils"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func mouseLeft(msg tea.MouseMsg, s *Screen) {
	if s.MenuType == symbolColor && msg.X < MenuSymbolColorWidth {
		selectSymbol(msg, s)
		selectColor(msg, s)
	} else if s.MenuType == file && msg.X < s.FileListWidth {
		selectFile(msg, s)
	} else if s.MenuType == shape && msg.X < MenuShapeWidth {
		selectShape(msg, s)
	} else {
		draw(msg, s)
	}
}

func selectShape(msg tea.MouseMsg, s *Screen) {
	if sh, ok := shapeList[msg.Y]; ok {
		s.Cursor.Store.Brush = sh.shapeType
		s.Cursor.Brush = sh.shapeType
	}
}

func selectColor(msg tea.MouseMsg, s *Screen) {
	if symbol, ok := config.Cfg.Symbols[msg.Y][msg.X]; ok {
		s.Cursor.Store.Symbol = symbol
		s.Cursor.Symbol = symbol
		if config.Cfg.Notifications.SetSymbol {
			s.SetMessage("Set " + symbol)
		}
	}
}

func selectSymbol(msg tea.MouseMsg, s *Screen) {
	if c, ok := Colors[msg.Y]; ok {
		s.Cursor.Color[c] = color.MinMaxColor(s.Cursor.Color[c])
	}
}

func selectFile(msg tea.MouseMsg, s *Screen) {
	if file, ok := s.FileList[msg.Y]; ok {
		content, err := os.ReadFile(s.Dir + file)
		if err != nil {
			s.Dir += file
		} else {
			s.MenuType = None
		}
		s.LoadImage(string(content))
	}
}

func draw(msg tea.MouseMsg, s *Screen) {
	if s.Cursor.Brush != Dot && len(s.Cursor.Pixels) > 1 {
		s.Pixels = append(
			s.Pixels,
			s.Cursor.Pixels...,
		)
	} else {
		s.Pixels = append(
			s.Pixels,
			Pixel{
				X: msg.X,
				Y: msg.Y,
				Symbol: utils.FgRgb(
					s.Cursor.Color["r"],
					s.Cursor.Color["g"],
					s.Cursor.Color["b"],
					s.Cursor.Symbol,
				),
			},
		)
	}
}
