package screen

import (
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/utils"
	tea "github.com/charmbracelet/bubbletea"
	"math"
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
	symbol := utils.FgRgb(
		s.Cursor.Color["r"],
		s.Cursor.Color["g"],
		s.Cursor.Color["b"],
		s.Cursor.Symbol,
	)

	switch s.Cursor.Brush {
	case Dot:
		s.Pixels = append(s.Pixels, Pixel{X: msg.X, Y: msg.Y, Symbol: symbol})
	case GLine:
		for i := 0; i < s.Cursor.Width; i++ {
			s.Pixels = append(
				s.Pixels,
				Pixel{X: msg.X + i, Y: msg.Y, Symbol: symbol},
			)
		}

	case VLine:
		for i := 0; i < s.Cursor.Width; i++ {
			s.Pixels = append(
				s.Pixels,
				Pixel{X: msg.X, Y: msg.Y + i, Symbol: symbol},
			)
		}

	case ESquare:
		for i := 0; i < s.Cursor.Height; i++ {
			for j := 0; j < s.Cursor.Width; j++ {
				if j > 0 && j < s.Cursor.Width-1 && i > 0 && i < s.Cursor.Height-1 {
					continue
				}
				s.Pixels = append(
					s.Pixels,
					Pixel{X: msg.X + j, Y: msg.Y + i, Symbol: symbol},
				)
			}
		}

	case FSquare:
		for i := 0; i < s.Cursor.Height; i++ {
			for j := 0; j < s.Cursor.Width; j++ {
				s.Pixels = append(
					s.Pixels,
					Pixel{X: msg.X + j, Y: msg.Y + i, Symbol: symbol},
				)
			}
		}

	case ECircle:
		R := s.Cursor.Width / 2
		k := 5
		for y := -R * k; y <= R*k; y++ {
			x := int(math.Sqrt(math.Pow(float64(R), 2)-math.Pow(float64(y)/float64(k), 2)) / .4583333333333333)
			ky := int(math.Round(float64(y) / float64(k)))
			s.Pixels = append(
				s.Pixels,
				Pixel{X: msg.X + x, Y: msg.Y + ky, Symbol: symbol},
				Pixel{X: msg.X - x, Y: msg.Y + ky, Symbol: symbol},
			)
		}

	case FCircle:
		R := s.Cursor.Width / 2
		k := 5
		for y := -R * k; y <= R*k; y++ {
			x := int(math.Sqrt(math.Pow(float64(R), 2)-math.Pow(float64(y)/float64(k), 2)) / .4583333333333333)
			ky := int(math.Round(float64(y) / float64(k)))
			for i := -x; i <= x; i++ {
				s.Pixels = append(
					s.Pixels,
					Pixel{X: msg.X + i, Y: msg.Y + ky, Symbol: symbol},
				)
			}
		}
	}
}
