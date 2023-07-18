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
			s.Cursor.Store.Symbol = symbol
			s.Cursor.Symbol = symbol
			s.SetMessage("Set " + symbol)
		}
		if c, ok := Colors[msg.Y]; ok {
			s.Cursor.Color[c] = color.MinMaxColor(s.Cursor.Color[c])
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
		if len(s.Cursor.Pixels) > 1 {
			for _, pixel := range s.Cursor.Pixels {
				s.Pixels = append(
					s.Pixels,
					Pixel{
						X: pixel.X,
						Y: pixel.Y,
						Symbol: utils.FgRgb(
							s.Cursor.Color["r"],
							s.Cursor.Color["g"],
							s.Cursor.Color["b"],
							s.Cursor.Symbol,
						),
					},
				)
			}
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
}
