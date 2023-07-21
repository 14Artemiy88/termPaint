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
	} else if s.MenuType == line && msg.X < MenuLineWidth {
		selectLine(msg, s)
	} else {
		draw(msg, s)
	}
}

func selectLine(msg tea.MouseMsg, s *Screen) {
	if line, ok := lineList[msg.Y]; ok {
		s.Cursor.Store.Brush = line.LineType
		s.Cursor.Store.Symbol = line.Cursor
	}
}

func selectShape(msg tea.MouseMsg, s *Screen) {
	if sh, ok := shapeList[msg.Y]; ok {
		s.Cursor.Store.Brush = sh.shapeType
	}
}

func selectColor(msg tea.MouseMsg, s *Screen) {
	if symbol, ok := config.Cfg.Symbols[msg.Y][msg.X]; ok {
		s.Cursor.setCursor(symbol)
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
		s.Pixels.add(Pixel{X: msg.X, Y: msg.Y, Symbol: symbol})
	case GLine:
		for i := 0; i < s.Cursor.Width; i++ {
			s.Pixels.add(Pixel{X: msg.X + i, Y: msg.Y, Symbol: symbol})
		}

	case VLine:
		for i := 0; i < s.Cursor.Width; i++ {
			s.Pixels.add(Pixel{X: msg.X, Y: msg.Y + i, Symbol: symbol})
		}

	case ESquare:
		for i := 0; i < s.Cursor.Height; i++ {
			for j := 0; j < s.Cursor.Width; j++ {
				if j > 0 && j < s.Cursor.Width-1 && i > 0 && i < s.Cursor.Height-1 {
					continue
				}
				s.Pixels.add(Pixel{X: msg.X + j, Y: msg.Y + i, Symbol: symbol})
			}
		}

	case FSquare:
		for i := 0; i < s.Cursor.Height; i++ {
			for j := 0; j < s.Cursor.Width; j++ {
				s.Pixels.add(Pixel{X: msg.X + j, Y: msg.Y + i, Symbol: symbol})
			}
		}

	case ECircle:
		R := s.Cursor.Width / 2
		k := 5
		for y := -R * k; y <= R*k; y++ {
			x := int(math.Sqrt(math.Pow(float64(R), 2)-math.Pow(float64(y)/float64(k), 2)) / .4583333333333333)
			ky := int(math.Round(float64(y) / float64(k)))
			s.Pixels.add(
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
				s.Pixels.add(Pixel{X: msg.X + i, Y: msg.Y + ky, Symbol: symbol})
			}
		}

	case ContinuousLine:
		var px int
		var py int
		var pr route
		if s.StorePixel[0].Symbol != "" {
			px = msg.X - s.StorePixel[0].X
			py = msg.Y - s.StorePixel[0].Y
			if px > 1 || px < -1 {
				px = 0
			}
			if py > 1 || py < -1 {
				py = 0
			}
		}
		pr = getRoute[py][px]

		var x int
		var y int
		var r route
		line := "─"
		prevLine := " "
		if s.StorePixel[1].Symbol != "" {
			x = msg.X - s.StorePixel[1].X // -1 0 1
			y = msg.Y - s.StorePixel[1].Y // -1 0 1
			if x < -1 || x > 1 || y < -1 || y > 1 {
				s.StorePixel = [2]Pixel{}
				x = 0
				y = 0
			}
			if x == 0 {
				line = "│"
				s.Cursor.setCursor("│")
			}
			if y == 0 {
				line = "─"
				s.Cursor.setCursor("─")
			}
		}
		r = getRoute[y][x]
		prevLine = drawLineList[pr][r]
		symbol = utils.FgRgb(
			s.Cursor.Color["r"],
			s.Cursor.Color["g"],
			s.Cursor.Color["b"],
			line,
		)
		pixel := Pixel{X: msg.X, Y: msg.Y, Symbol: symbol}
		prevPixel := Pixel{X: msg.X - x, Y: msg.Y - y, Symbol: prevLine}

		//s.Pixels.add(pixel)
		s.Pixels.add(prevPixel)
		//s.Pixels.add(pixel, prevPixel)
		s.StorePixel.restore(pixel, prevPixel)
	}
}

func limit(value int, min int, max int) int {
	if value <= min {
		return min
	}
	if value >= max {
		return max
	}

	return 0
}
