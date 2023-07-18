package screen

import (
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/utils"
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
		s.Cursor.Symbol = " "
	} else {
		s.Cursor.Symbol = s.Cursor.Store
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
		s.Cursor.Symbol = utils.FgRgb(
			config.Cfg.PointerColor["r"],
			config.Cfg.PointerColor["g"],
			config.Cfg.PointerColor["b"],
			config.Cfg.Pointer,
		)
		s.File = file
	} else {
		s.Cursor.Symbol = " "
		s.File = ""
	}
}

func onMenu(msg tea.MouseMsg, s *Screen) {
	s.Cursor.Symbol = " "
	_, okSymbol := config.Cfg.Symbols[msg.Y][msg.X]
	c, okColor := Colors[msg.Y]
	//if okSymbol {
	//	s.X = msg.X - 1
	//	s.Cursor = FgRgb(Cfg.PointerColor["r"], Cfg.PointerColor["g"], Cfg.PointerColor["b"], Cfg.Pointer)
	//}
	if okColor && msg.X < MenuWidth {
		s.InputLock = true
		s.InputColor = c
		s.X = 0
		s.Cursor.Symbol = utils.FgRgb(
			config.Cfg.PointerColor["r"],
			config.Cfg.PointerColor["g"],
			config.Cfg.PointerColor["b"],
			config.Cfg.Pointer,
		)
	} else {
		s.InputLock = false
		if len(s.Input) > 0 {
			s.Cursor.Color[s.InputColor] = color.SetColor(s.Input)
		}
		s.Input = ""
	}
	if !okSymbol && !okColor {
		s.X = MenuWidth + 1
		s.Cursor.Symbol = " "
	}
}
