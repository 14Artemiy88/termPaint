package screen

import (
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/menu"
	"github.com/14Artemiy88/termPaint/src/message"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/utils"
	"math"
	"os"
	"path/filepath"
)

func mouseLeft(X int, Y int, s *Screen) {
	if menu.Type == menu.SymbolColor && X < menu.SymbolColorWidth {
		selectSymbol(Y)
		selectColor(X, Y)
	} else if menu.Type == menu.File && X < menu.FileListWidth {
		selectFile(Y, s)
	} else if menu.Type == menu.Shape && X < menu.ShapeWidth {
		selectShape(Y)
	} else if menu.Type == menu.Line && X < menu.LineWidth {
		selectLine(Y)
	} else {
		draw(X, Y)
	}
}

func selectLine(Y int) {
	if line, ok := menu.LineList[Y]; ok {
		cursor.CC.Store.Brush = line.LineType
		if line.LineType == cursor.Dot {
			cursor.CC.SetCursor(config.Cfg.DefaultCursor)
		} else {
			cursor.CC.SetCursor(line.Cursor)
		}
	}
}

func selectShape(Y int) {
	if sh, ok := menu.ShapeList[Y]; ok {
		cursor.CC.Store.Brush = sh.ShapeType
	}
}

func selectColor(X int, Y int) {
	if symbol, ok := config.Cfg.Symbols[Y][X]; ok {
		cursor.CC.SetCursor(symbol)
		if config.Cfg.Notifications.SetSymbol {
			message.SetMessage("Set " + symbol)
		}
	}
}

func selectSymbol(Y int) {
	if c, ok := menu.Colors[Y]; ok {
		cursor.CC.Color[c] = color.MinMaxColor(cursor.CC.Color[c])
	}
}

func selectFile(Y int, s *Screen) {
	if filePath, ok := menu.FileList[Y]; ok {
		info, err := os.Stat(menu.Dir + filePath)
		if err != nil {
			message.SetMessage(err.Error())
		}
		if info.IsDir() {
			menu.Dir += filePath
		} else {
			menu.Type = menu.None
			ext := filepath.Ext(menu.Dir + filePath)
			if ext == ".txt" {
				content, err := os.ReadFile(menu.Dir + filePath)
				if err != nil {
					message.SetMessage(err.Error())
				}
				s.LoadImage(string(content))
			}
			if ext == ".jpg" || ext == ".png" {
				s.loadFromImafe(menu.Dir + filePath)
			}
		}
	}
}

func draw(X int, Y int) {
	symbol := utils.FgRgb(
		cursor.CC.Color["r"],
		cursor.CC.Color["g"],
		cursor.CC.Color["b"],
		cursor.CC.Symbol,
	)

	switch cursor.CC.Brush {
	case cursor.Dot:
		drawDot(pixel.Pixel{X: X, Y: Y, Symbol: symbol})
	case cursor.GLine:
		drawGLine(X, Y, symbol)
	case cursor.VLine:
		drawVLine(X, Y, symbol)
	case cursor.ESquare:
		drawESquare(X, Y, symbol)
	case cursor.FSquare:
		drawFSquare(X, Y, symbol)
	case cursor.ECircle:
		drawECircle(X, Y, symbol)
	case cursor.FCircle:
		drawFCircle(X, Y, symbol)
	case cursor.ContinuousLine, cursor.SmoothContinuousLine, cursor.FatContinuousLine, cursor.DoubleContinuousLine:
		drawContinuousLine(X, Y)
	}
}

func drawDot(p pixel.Pixel) {
	pixel.Pixels.Add(p)
}

func drawGLine(x int, y int, symbol string) {
	for i := 0; i < cursor.CC.Width; i++ {
		pixel.Pixels.Add(pixel.Pixel{X: x + i, Y: y, Symbol: symbol})
	}
}

func drawVLine(x int, y int, symbol string) {
	for i := 0; i < cursor.CC.Width; i++ {
		pixel.Pixels.Add(pixel.Pixel{X: x, Y: y + i, Symbol: symbol})
	}
}

func drawESquare(x int, y int, symbol string) {
	for i := 0; i < cursor.CC.Height; i++ {
		for j := 0; j < cursor.CC.Width; j++ {
			if j > 0 && j < cursor.CC.Width-1 && i > 0 && i < cursor.CC.Height-1 {
				continue
			}
			pixel.Pixels.Add(pixel.Pixel{X: x + j, Y: y + i, Symbol: symbol})
		}
	}
}

func drawFSquare(x int, y int, symbol string) {
	for i := 0; i < cursor.CC.Height; i++ {
		for j := 0; j < cursor.CC.Width; j++ {
			pixel.Pixels.Add(pixel.Pixel{X: x + j, Y: y + i, Symbol: symbol})
		}
	}
}

func drawECircle(X int, Y int, symbol string) {
	R := cursor.CC.Width / 2
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
	R := cursor.CC.Width / 2
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
		line := cursor.CC.Store.Symbol
		if pixel.StorePixel[1].Symbol != "" {
			if x < -1 || x > 1 || y < -1 || y > 1 {
				pixel.StorePixel = [2]pixel.Pixel{}
				x = 0
				y = 0
			}
			if x == 0 {
				line = menu.GVLine[line]["v"]
				cursor.CC.SetCursor(line)
			}
			if y == 0 {
				line = menu.GVLine[line]["g"]
				cursor.CC.SetCursor(line)
			}
		}

		p := pixel.Pixel{
			X: X,
			Y: Y,
			Symbol: utils.FgRgb(
				cursor.CC.Color["r"],
				cursor.CC.Color["g"],
				cursor.CC.Color["b"],
				line,
			),
		}
		pixel.Pixels.Add(p)

		var px int
		var py int
		var pr menu.Route
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
		r := menu.GetRoute[y][x]
		pr = menu.GetRoute[py][px]

		pixel.Pixels.Add(
			pixel.Pixel{
				X: pixel.StorePixel[1].X,
				Y: pixel.StorePixel[1].Y,
				Symbol: utils.FgRgb(
					cursor.CC.Color["r"],
					cursor.CC.Color["g"],
					cursor.CC.Color["b"],
					menu.DrawLine[cursor.CC.Store.Brush][pr][r],
				),
			},
		)

		pixel.StorePixel = [2]pixel.Pixel{pixel.StorePixel[1], p}
	}
}
