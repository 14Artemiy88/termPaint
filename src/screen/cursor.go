package screen

import (
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/utils"
	"math"
)

type Cursor struct {
	//X      int
	//Y      int
	Pixels []Pixel
	Brush  cursorType
	Width  int
	Symbol string
	Color  map[string]int
	Store  Store
}

type Store struct {
	Symbol string
	Brush  cursorType
}

const emptyCursor = " "

type cursorType int

const (
	Dot cursorType = iota
	Empty
	Pointer
	GLine
	VLine
	ESquare
	FSquare
	ECircle
	FCircle
)

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

	case Dot:
		s.Cursor.Pixels = []Pixel{}
		utils.SetByKeys(s.X, s.Y, symbol, screen)

	case GLine:
		s.Cursor.Pixels = []Pixel{}
		for i := 0; i < s.Cursor.Width; i++ {
			s.Cursor.Pixels = append(
				s.Cursor.Pixels,
				Pixel{X: s.X + i, Y: s.Y, Symbol: symbol},
			)
			utils.SetByKeys(s.X+i, s.Y, symbol, screen)
		}

	case VLine:
		s.Cursor.Pixels = []Pixel{}
		for i := 0; i < s.Cursor.Width; i++ {
			s.Cursor.Pixels = append(
				s.Cursor.Pixels,
				Pixel{X: s.X, Y: s.Y + i, Symbol: symbol},
			)
			utils.SetByKeys(s.X, s.Y+i, symbol, screen)
		}

	case ESquare:
		s.Cursor.Pixels = []Pixel{}
		for i := 0; i < s.Cursor.Width; i++ {
			for j := 0; j < s.Cursor.Width; j++ {
				if j > 0 && j < s.Cursor.Width-1 && i > 0 && i < s.Cursor.Width-1 {
					continue
				}
				s.Cursor.Pixels = append(
					s.Cursor.Pixels,
					Pixel{X: s.X + j, Y: s.Y + i, Symbol: symbol},
				)
				utils.SetByKeys(s.X+j, s.Y+i, symbol, screen)
			}
		}

	case FSquare:
		s.Cursor.Pixels = []Pixel{}
		for i := 0; i < s.Cursor.Width; i++ {
			for j := 0; j < s.Cursor.Width; j++ {
				s.Cursor.Pixels = append(
					s.Cursor.Pixels,
					Pixel{X: s.X + j, Y: s.Y + i, Symbol: symbol},
				)
				utils.SetByKeys(s.X+j, s.Y+i, symbol, screen)
			}
		}

	case ECircle:
		s.Cursor.Pixels = []Pixel{}
		R := s.Cursor.Width / 2
		k := 5
		for y := -R * k; y <= R*k; y++ {
			x := int(
				math.Sqrt(
					math.Pow(float64(R), 2)-math.Pow(float64(y)/float64(k), 2),
				) / .4583333333333333,
			)
			ky := int(math.Round(float64(y) / float64(k)))
			s.Cursor.Pixels = append(
				s.Cursor.Pixels,
				Pixel{X: s.X + x, Y: s.Y + ky, Symbol: symbol},
				Pixel{X: s.X - x, Y: s.Y + ky, Symbol: symbol},
			)

			utils.SetByKeys(s.X+x, s.Y+ky, symbol, screen)
			utils.SetByKeys(s.X-x, s.Y+ky, symbol, screen)
		}

	case FCircle:
		s.Cursor.Pixels = []Pixel{}
		R := s.Cursor.Width / 2
		k := 5
		for y := -R * k; y <= R*k; y++ {
			x := int(
				math.Sqrt(
					math.Pow(float64(R), 2)-math.Pow(float64(y)/float64(k), 2),
				) / .4583333333333333,
			)
			ky := int(math.Round(float64(y) / float64(k)))
			for i := -x; i <= x; i++ {
				s.Cursor.Pixels = append(
					s.Cursor.Pixels,
					Pixel{X: s.X + i, Y: s.Y + ky, Symbol: symbol},
				)
				utils.SetByKeys(s.X+i, s.Y+ky, symbol, screen)
			}
		}
	}

	return screen
}
