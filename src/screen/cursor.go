package screen

import (
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/utils"
	"math"
)

type Cursor struct {
	X      int
	Y      int
	Brush  cursorType
	Width  int
	Height int
	Symbol string
	Color  map[string]int
	Store  Store
}

var CC Cursor

type Store struct {
	Symbol string
	Brush  cursorType
}

type cursorType int

const (
	Empty cursorType = iota
	Pointer
	Dot
	GLine
	VLine
	ESquare
	FSquare
	ECircle
	FCircle
	ContinuousLine
	SmoothContinuousLine
	FatContinuousLine
	DoubleContinuousLine
)

func (c *Cursor) setCursor(cursor string) {
	c.Symbol = cursor
	c.Store.Symbol = cursor
}

func (c *Cursor) DrawCursor(screen [][]string) [][]string {
	symbol := utils.FgRgb(
		CC.Color["r"],
		CC.Color["g"],
		CC.Color["b"],
		CC.Symbol,
	)
	switch CC.Brush {
	case Empty:
	case Pointer:
		c.X = 1
		symbol = utils.FgRgb(
			config.Cfg.PointerColor["r"],
			config.Cfg.PointerColor["g"],
			config.Cfg.PointerColor["b"],
			config.Cfg.Pointer,
		)
		utils.SetByKeys(1, c.Y, symbol, screen)

	case Dot,
		ContinuousLine,
		SmoothContinuousLine,
		FatContinuousLine,
		DoubleContinuousLine:
		utils.SetByKeys(c.X, c.Y, symbol, screen)

	case GLine:
		for i := 0; i < CC.Width; i++ {
			utils.SetByKeys(c.X+i, c.Y, symbol, screen)
		}

	case VLine:
		for i := 0; i < CC.Width; i++ {
			utils.SetByKeys(c.X, c.Y+i, symbol, screen)
		}

	case ESquare:
		for y := 0; y < CC.Height; y++ {
			for x := 0; x < CC.Width; x++ {
				if x > 0 && x < CC.Width-1 && y > 0 && y < CC.Height-1 {
					continue
				}
				utils.SetByKeys(c.X+x, c.Y+y, symbol, screen)
			}
		}

	case FSquare:
		for y := 0; y < CC.Height; y++ {
			for x := 0; x < CC.Width; x++ {
				utils.SetByKeys(c.X+x, c.Y+y, symbol, screen)
			}
		}

	case ECircle:
		R := CC.Width / 2
		k := 5
		for y := -R * k; y <= R*k; y++ {
			x := int(math.Sqrt(math.Pow(float64(R), 2)-math.Pow(float64(y)/float64(k), 2)) / pixelRatio)
			ky := int(math.Round(float64(y) / float64(k)))
			utils.SetByKeys(c.X+x, c.Y+ky, symbol, screen)
			utils.SetByKeys(c.X-x, c.Y+ky, symbol, screen)
		}

	case FCircle:
		R := CC.Width / 2
		k := 5
		for y := -R * k; y <= R*k; y++ {
			x := int(math.Sqrt(math.Pow(float64(R), 2)-math.Pow(float64(y)/float64(k), 2)) / pixelRatio)
			ky := int(math.Round(float64(y) / float64(k)))
			for i := -x; i <= x; i++ {
				utils.SetByKeys(c.X+i, c.Y+ky, symbol, screen)
			}
		}
	}

	return screen
}
