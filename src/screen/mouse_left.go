package screen

import (
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/utils"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func mouseLeft(msg tea.MouseMsg, s *Screen) {
	if s.ShowMenu && msg.X < MenuWidth {
		if symbol, ok := config.Cfg.Symbols[msg.Y][msg.X]; ok {
			s.CursorStore = symbol
			s.Cursor = symbol
		}
		if c, ok := Colors[msg.Y]; ok {
			s.Color[c] = color.MinMaxColor(s.Color[c])
		}
	} else if s.ShowFile && msg.X < s.FileListWidth {
		if file, ok := s.FileList[msg.Y]; ok {
			content, err := os.ReadFile(s.Dir + file)
			if err != nil {
				s.Dir += file
			} else {
				s.ShowFile = false
			}
			s.LoadImage(string(content))
		}
	} else {
		s.Pixels = append(s.Pixels, Pixel{X: msg.X, Y: msg.Y, Symbol: utils.FgRgb(s.Color["r"], s.Color["g"], s.Color["b"], s.Cursor)})
	}
}
