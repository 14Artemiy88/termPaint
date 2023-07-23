package screen

import (
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/utils"
	"math"
)

type Cursor struct {
	Brush  cursorType
	Width  int
	Height int
	Symbol string
	Color  map[string]int
	Store  Store
}

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

func DrawCursor(s *Screen, screen [][]string) [][]string {
	symbol := utils.FgRgb(
		s.Cursor.Color["r"],
		s.Cursor.Color["g"],
		s.Cursor.Color["b"],
		s.Cursor.Symbol,
	)
	switch s.Cursor.Brush {
	case Empty:
	case Pointer:
		s.X = 1
		symbol = utils.FgRgb(
			config.Cfg.PointerColor["r"],
			config.Cfg.PointerColor["g"],
			config.Cfg.PointerColor["b"],
			config.Cfg.Pointer,
		)
		utils.SetByKeys(1, s.Y, symbol, screen)

	case Dot,
		ContinuousLine,
		SmoothContinuousLine,
		FatContinuousLine,
		DoubleContinuousLine:
		utils.SetByKeys(s.X, s.Y, symbol, screen)

	case GLine:
		for i := 0; i < s.Cursor.Width; i++ {
			utils.SetByKeys(s.X+i, s.Y, symbol, screen)
		}

	case VLine:
		for i := 0; i < s.Cursor.Width; i++ {
			utils.SetByKeys(s.X, s.Y+i, symbol, screen)
		}

	case ESquare:
		for y := 0; y < s.Cursor.Height; y++ {
			for x := 0; x < s.Cursor.Width; x++ {
				if x > 0 && x < s.Cursor.Width-1 && y > 0 && y < s.Cursor.Height-1 {
					continue
				}
				utils.SetByKeys(s.X+x, s.Y+y, symbol, screen)
			}
		}

	case FSquare:
		for y := 0; y < s.Cursor.Height; y++ {
			for x := 0; x < s.Cursor.Width; x++ {
				utils.SetByKeys(s.X+x, s.Y+y, symbol, screen)
			}
		}

	case ECircle:
		R := s.Cursor.Width / 2
		k := 5
		for y := -R * k; y <= R*k; y++ {
			x := int(math.Sqrt(math.Pow(float64(R), 2)-math.Pow(float64(y)/float64(k), 2)) / pixelRatio)
			ky := int(math.Round(float64(y) / float64(k)))
			utils.SetByKeys(s.X+x, s.Y+ky, symbol, screen)
			utils.SetByKeys(s.X-x, s.Y+ky, symbol, screen)
		}

	case FCircle:
		R := s.Cursor.Width / 2
		k := 5
		for y := -R * k; y <= R*k; y++ {
			x := int(math.Sqrt(math.Pow(float64(R), 2)-math.Pow(float64(y)/float64(k), 2)) / pixelRatio)
			ky := int(math.Round(float64(y) / float64(k)))
			for i := -x; i <= x; i++ {
				utils.SetByKeys(s.X+i, s.Y+ky, symbol, screen)
			}
		}
	}

	return screen
}
