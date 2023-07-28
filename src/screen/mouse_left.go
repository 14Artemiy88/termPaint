package screen

import (
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/utils"
	tea "github.com/charmbracelet/bubbletea"
	"math"
	"os"
	"path/filepath"
)

func mouseLeft(msg tea.MouseMsg, s *Screen) {
	if s.MenuType == symbolColor && msg.X < MenuSymbolColorWidth {
		selectSymbol(msg)
		selectColor(msg)
	} else if s.MenuType == file && msg.X < s.FileListWidth {
		selectFile(msg, s)
	} else if s.MenuType == shape && msg.X < MenuShapeWidth {
		selectShape(msg)
	} else if s.MenuType == line && msg.X < MenuLineWidth {
		selectLine(msg)
	} else {
		draw(msg, s)
	}
}

func selectLine(msg tea.MouseMsg) {
	if line, ok := menuLineList[msg.Y]; ok {
		CC.Store.Brush = line.LineType
		if line.LineType == Dot {
			CC.setCursor(config.Cfg.DefaultCursor)
		} else {
			CC.setCursor(line.Cursor)
		}
	}
}

func selectShape(msg tea.MouseMsg) {
	if sh, ok := shapeList[msg.Y]; ok {
		CC.Store.Brush = sh.shapeType
	}
}

func selectColor(msg tea.MouseMsg) {
	if symbol, ok := config.Cfg.Symbols[msg.Y][msg.X]; ok {
		CC.setCursor(symbol)
		if config.Cfg.Notifications.SetSymbol {
			SetMessage("Set " + symbol)
		}
	}
}

func selectSymbol(msg tea.MouseMsg) {
	if c, ok := Colors[msg.Y]; ok {
		CC.Color[c] = color.MinMaxColor(CC.Color[c])
	}
}

func selectFile(msg tea.MouseMsg, s *Screen) {
	if filePath, ok := s.FileList[msg.Y]; ok {
		info, err := os.Stat(Dir + filePath)
		if err != nil {
			SetMessage(err.Error())
		}
		if info.IsDir() {
			Dir += filePath
		} else {
			s.MenuType = None
			ext := filepath.Ext(Dir + filePath)
			if ext == ".txt" {
				content, err := os.ReadFile(Dir + filePath)
				if err != nil {
					SetMessage(err.Error())
				}
				s.LoadImage(string(content))
			}
			if ext == ".jpg" || ext == ".png" {
				s.loadFromImafe(Dir + filePath)
			}
		}
	}
}

func draw(msg tea.MouseMsg, s *Screen) {
	symbol := utils.FgRgb(
		CC.Color["r"],
		CC.Color["g"],
		CC.Color["b"],
		CC.Symbol,
	)

	switch CC.Brush {
	case Dot:
		drawDot(s, Pixel{X: msg.X, Y: msg.Y, Symbol: symbol})
	case GLine:
		drawGLine(s, msg.X, msg.Y, symbol)
	case VLine:
		drawVLine(s, msg.X, msg.Y, symbol)
	case ESquare:
		drawESquare(s, msg.X, msg.Y, symbol)
	case FSquare:
		drawFSquare(s, msg.X, msg.Y, symbol)
	case ECircle:
		drawECircle(s, msg.X, msg.Y, symbol)
	case FCircle:
		drawFCircle(s, msg.X, msg.Y, symbol)
	case ContinuousLine, SmoothContinuousLine, FatContinuousLine, DoubleContinuousLine:
		drawContinuousLine(s, msg.X, msg.Y)
	}
}

func drawDot(s *Screen, pixel Pixel) {
	s.Pixels.add(pixel)
}

func drawGLine(s *Screen, x int, y int, symbol string) {
	for i := 0; i < CC.Width; i++ {
		s.Pixels.add(Pixel{X: x + i, Y: y, Symbol: symbol})
	}
}

func drawVLine(s *Screen, x int, y int, symbol string) {
	for i := 0; i < CC.Width; i++ {
		s.Pixels.add(Pixel{X: x, Y: y + i, Symbol: symbol})
	}
}

func drawESquare(s *Screen, x int, y int, symbol string) {
	for i := 0; i < CC.Height; i++ {
		for j := 0; j < CC.Width; j++ {
			if j > 0 && j < CC.Width-1 && i > 0 && i < CC.Height-1 {
				continue
			}
			s.Pixels.add(Pixel{X: x + j, Y: y + i, Symbol: symbol})
		}
	}
}

func drawFSquare(s *Screen, x int, y int, symbol string) {
	for i := 0; i < CC.Height; i++ {
		for j := 0; j < CC.Width; j++ {
			s.Pixels.add(Pixel{X: x + j, Y: y + i, Symbol: symbol})
		}
	}
}

func drawECircle(s *Screen, X int, Y int, symbol string) {
	R := CC.Width / 2
	k := 5
	for y := -R * k; y <= R*k; y++ {
		x := int(math.Sqrt(math.Pow(float64(R), 2)-math.Pow(float64(y)/float64(k), 2)) / pixelRatio)
		ky := int(math.Round(float64(y) / float64(k)))
		s.Pixels.add(
			Pixel{X: X + x, Y: Y + ky, Symbol: symbol},
			Pixel{X: X - x, Y: Y + ky, Symbol: symbol},
		)
	}
}

func drawFCircle(s *Screen, X int, Y int, symbol string) {
	R := CC.Width / 2
	k := 5
	for y := -R * k; y <= R*k; y++ {
		x := int(math.Sqrt(math.Pow(float64(R), 2)-math.Pow(float64(y)/float64(k), 2)) / pixelRatio)
		ky := int(math.Round(float64(y) / float64(k)))
		for i := -x; i <= x; i++ {
			s.Pixels.add(Pixel{X: X + i, Y: Y + ky, Symbol: symbol})
		}
	}
}

func drawContinuousLine(s *Screen, X int, Y int) {
	x := X - s.StorePixel[1].X // -1 0 1
	y := Y - s.StorePixel[1].Y // -1 0 1
	if x != 0 || y != 0 {
		line := CC.Store.Symbol
		if s.StorePixel[1].Symbol != "" {
			if x < -1 || x > 1 || y < -1 || y > 1 {
				s.StorePixel = [2]Pixel{}
				x = 0
				y = 0
			}
			if x == 0 {
				line = gvLine[line]["v"]
				CC.setCursor(line)
			}
			if y == 0 {
				line = gvLine[line]["g"]
				CC.setCursor(line)
			}
		}

		pixel := Pixel{
			X: X,
			Y: Y,
			Symbol: utils.FgRgb(
				CC.Color["r"],
				CC.Color["g"],
				CC.Color["b"],
				line,
			),
		}
		s.Pixels.add(pixel)

		var px int
		var py int
		var pr route
		if s.StorePixel[0].Symbol != "" {
			px = X - s.StorePixel[0].X
			py = Y - s.StorePixel[0].Y
			if px > 1 || px < -1 {
				px = 0
			}
			if py > 1 || py < -1 {
				py = 0
			}
		}
		r := getRoute[y][x]
		pr = getRoute[py][px]

		s.Pixels.add(
			Pixel{
				X: s.StorePixel[1].X,
				Y: s.StorePixel[1].Y,
				Symbol: utils.FgRgb(
					CC.Color["r"],
					CC.Color["g"],
					CC.Color["b"],
					drawLine[CC.Store.Brush][pr][r],
				),
			},
		)

		s.StorePixel = [2]Pixel{s.StorePixel[1], pixel}
	}
}
