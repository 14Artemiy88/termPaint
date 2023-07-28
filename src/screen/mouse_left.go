package screen

import (
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/pixel"
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
		draw(msg)
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

func draw(msg tea.MouseMsg) {
	symbol := utils.FgRgb(
		CC.Color["r"],
		CC.Color["g"],
		CC.Color["b"],
		CC.Symbol,
	)

	switch CC.Brush {
	case Dot:
		drawDot(pixel.Pixel{X: msg.X, Y: msg.Y, Symbol: symbol})
	case GLine:
		drawGLine(msg.X, msg.Y, symbol)
	case VLine:
		drawVLine(msg.X, msg.Y, symbol)
	case ESquare:
		drawESquare(msg.X, msg.Y, symbol)
	case FSquare:
		drawFSquare(msg.X, msg.Y, symbol)
	case ECircle:
		drawECircle(msg.X, msg.Y, symbol)
	case FCircle:
		drawFCircle(msg.X, msg.Y, symbol)
	case ContinuousLine, SmoothContinuousLine, FatContinuousLine, DoubleContinuousLine:
		drawContinuousLine(msg.X, msg.Y)
	}
}

func drawDot(p pixel.Pixel) {
	pixel.Pixels.Add(p)
}

func drawGLine(x int, y int, symbol string) {
	for i := 0; i < CC.Width; i++ {
		pixel.Pixels.Add(pixel.Pixel{X: x + i, Y: y, Symbol: symbol})
	}
}

func drawVLine(x int, y int, symbol string) {
	for i := 0; i < CC.Width; i++ {
		pixel.Pixels.Add(pixel.Pixel{X: x, Y: y + i, Symbol: symbol})
	}
}

func drawESquare(x int, y int, symbol string) {
	for i := 0; i < CC.Height; i++ {
		for j := 0; j < CC.Width; j++ {
			if j > 0 && j < CC.Width-1 && i > 0 && i < CC.Height-1 {
				continue
			}
			pixel.Pixels.Add(pixel.Pixel{X: x + j, Y: y + i, Symbol: symbol})
		}
	}
}

func drawFSquare(x int, y int, symbol string) {
	for i := 0; i < CC.Height; i++ {
		for j := 0; j < CC.Width; j++ {
			pixel.Pixels.Add(pixel.Pixel{X: x + j, Y: y + i, Symbol: symbol})
		}
	}
}

func drawECircle(X int, Y int, symbol string) {
	R := CC.Width / 2
	k := 5
	for y := -R * k; y <= R*k; y++ {
		x := int(math.Sqrt(math.Pow(float64(R), 2)-math.Pow(float64(y)/float64(k), 2)) / pixel.PixelRatio)
		ky := int(math.Round(float64(y) / float64(k)))
		pixel.Pixels.Add(
			pixel.Pixel{X: X + x, Y: Y + ky, Symbol: symbol},
			pixel.Pixel{X: X - x, Y: Y + ky, Symbol: symbol},
		)
	}
}

func drawFCircle(X int, Y int, symbol string) {
	R := CC.Width / 2
	k := 5
	for y := -R * k; y <= R*k; y++ {
		x := int(math.Sqrt(math.Pow(float64(R), 2)-math.Pow(float64(y)/float64(k), 2)) / pixel.PixelRatio)
		ky := int(math.Round(float64(y) / float64(k)))
		for i := -x; i <= x; i++ {
			pixel.Pixels.Add(pixel.Pixel{X: X + i, Y: Y + ky, Symbol: symbol})
		}
	}
}

func drawContinuousLine(X int, Y int) {
	x := X - pixel.StorePixel[1].X // -1 0 1
	y := Y - pixel.StorePixel[1].Y // -1 0 1
	if x != 0 || y != 0 {
		line := CC.Store.Symbol
		if pixel.StorePixel[1].Symbol != "" {
			if x < -1 || x > 1 || y < -1 || y > 1 {
				pixel.StorePixel = [2]pixel.Pixel{}
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

		p := pixel.Pixel{
			X: X,
			Y: Y,
			Symbol: utils.FgRgb(
				CC.Color["r"],
				CC.Color["g"],
				CC.Color["b"],
				line,
			),
		}
		pixel.Pixels.Add(p)

		var px int
		var py int
		var pr route
		if pixel.StorePixel[0].Symbol != "" {
			px = X - pixel.StorePixel[0].X
			py = Y - pixel.StorePixel[0].Y
			if px > 1 || px < -1 {
				px = 0
			}
			if py > 1 || py < -1 {
				py = 0
			}
		}
		r := getRoute[y][x]
		pr = getRoute[py][px]

		pixel.Pixels.Add(
			pixel.Pixel{
				X: pixel.StorePixel[1].X,
				Y: pixel.StorePixel[1].Y,
				Symbol: utils.FgRgb(
					CC.Color["r"],
					CC.Color["g"],
					CC.Color["b"],
					drawLine[CC.Store.Brush][pr][r],
				),
			},
		)

		pixel.StorePixel = [2]pixel.Pixel{pixel.StorePixel[1], p}
	}
}
