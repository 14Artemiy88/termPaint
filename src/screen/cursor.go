package screen

import (
	"github.com/14Artemiy88/termPaint/src/utils"
	"math"
)

type Cursor struct {
	X      int
	Y      int
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
	switch s.Cursor.Brush {
	case Empty:
	case Pointer:
	case Dot:
		s.Cursor.Pixels = []Pixel{}
		utils.SetByKeys(
			s.X,
			s.Y,
			utils.FgRgb(
				s.Cursor.Color["r"],
				s.Cursor.Color["g"],
				s.Cursor.Color["b"],
				s.Cursor.Symbol,
			),
			screen,
		)

	case GLine:
		s.Cursor.Pixels = []Pixel{}
		for i := 0; i < s.Cursor.Width; i++ {
			s.Cursor.Pixels = append(
				s.Cursor.Pixels,
				Pixel{
					X: s.X + i,
					Y: s.Y,
					Symbol: utils.FgRgb(
						s.Cursor.Color["r"],
						s.Cursor.Color["g"],
						s.Cursor.Color["b"],
						s.Cursor.Symbol,
					),
				},
			)
			utils.SetByKeys(
				s.X+i,
				s.Y,
				utils.FgRgb(
					s.Cursor.Color["r"],
					s.Cursor.Color["g"],
					s.Cursor.Color["b"],
					s.Cursor.Symbol,
				),
				screen,
			)
		}
	case VLine:
		s.Cursor.Pixels = []Pixel{}
		for i := 0; i < s.Cursor.Width; i++ {
			s.Cursor.Pixels = append(
				s.Cursor.Pixels,
				Pixel{
					X: s.X,
					Y: s.Y + i,
					Symbol: utils.FgRgb(
						s.Cursor.Color["r"],
						s.Cursor.Color["g"],
						s.Cursor.Color["b"],
						s.Cursor.Symbol,
					),
				},
			)
			utils.SetByKeys(
				s.X,
				s.Y+i,
				utils.FgRgb(
					s.Cursor.Color["r"],
					s.Cursor.Color["g"],
					s.Cursor.Color["b"],
					s.Cursor.Symbol,
				),
				screen,
			)
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
					Pixel{
						X: s.X + j,
						Y: s.Y + i,
						Symbol: utils.FgRgb(
							s.Cursor.Color["r"],
							s.Cursor.Color["g"],
							s.Cursor.Color["b"],
							s.Cursor.Symbol,
						),
					},
				)
				utils.SetByKeys(
					s.X+j,
					s.Y+i,
					utils.FgRgb(
						s.Cursor.Color["r"],
						s.Cursor.Color["g"],
						s.Cursor.Color["b"],
						s.Cursor.Symbol,
					),
					screen,
				)
			}
		}
	case FSquare:
		s.Cursor.Pixels = []Pixel{}
		for i := 0; i < s.Cursor.Width; i++ {
			for j := 0; j < s.Cursor.Width; j++ {
				s.Cursor.Pixels = append(
					s.Cursor.Pixels,
					Pixel{
						X: s.X + j,
						Y: s.Y + i,
						Symbol: utils.FgRgb(
							s.Cursor.Color["r"],
							s.Cursor.Color["g"],
							s.Cursor.Color["b"],
							s.Cursor.Symbol,
						),
					},
				)
				utils.SetByKeys(
					s.X+j,
					s.Y+i,
					utils.FgRgb(
						s.Cursor.Color["r"],
						s.Cursor.Color["g"],
						s.Cursor.Color["b"],
						s.Cursor.Symbol,
					),
					screen,
				)
			}
		}
	case ECircle:
		s.Cursor.Pixels = []Pixel{}
		for y := -s.Cursor.Width * 10; y <= s.Cursor.Width*10; y++ {
			x := int(
				math.Sqrt(
					math.Pow(float64(s.Cursor.Width), 2)-math.Pow(float64(y)/10.0, 2),
				) / .4583333333333333,
			)

			s.Cursor.Pixels = append(
				s.Cursor.Pixels,
				Pixel{
					X: s.X + x,
					Y: s.Y + y/10,
					Symbol: utils.FgRgb(
						s.Cursor.Color["r"],
						s.Cursor.Color["g"],
						s.Cursor.Color["b"],
						s.Cursor.Symbol,
					),
				},
			)

			s.Cursor.Pixels = append(
				s.Cursor.Pixels,
				Pixel{
					X: s.X - x,
					Y: s.Y + y/10,
					Symbol: utils.FgRgb(
						s.Cursor.Color["r"],
						s.Cursor.Color["g"],
						s.Cursor.Color["b"],
						s.Cursor.Symbol,
					),
				},
			)
			utils.SetByKeys(
				s.X+x,
				s.Y+y/10,
				utils.FgRgb(
					s.Cursor.Color["r"],
					s.Cursor.Color["g"],
					s.Cursor.Color["b"],
					s.Cursor.Symbol,
				),
				screen,
			)

			utils.SetByKeys(
				s.X-x,
				s.Y+y/10,
				utils.FgRgb(
					s.Cursor.Color["r"],
					s.Cursor.Color["g"],
					s.Cursor.Color["b"],
					s.Cursor.Symbol,
				),
				screen,
			)
		}
	case FCircle:
		s.Cursor.Pixels = []Pixel{}
		R := s.Cursor.Width / 2
		for y := -R; y <= R; y++ {
			x := int(
				math.Sqrt(
					math.Pow(float64(R), 2)-math.Pow(float64(y), 2),
				) / .4583333333333333,
			)
			for i := -x; i < x; i++ {
				s.Cursor.Pixels = append(
					s.Cursor.Pixels,
					Pixel{
						X: s.X + i,
						Y: s.Y + y,
						Symbol: utils.FgRgb(
							s.Cursor.Color["r"],
							s.Cursor.Color["g"],
							s.Cursor.Color["b"],
							s.Cursor.Symbol,
						),
					},
				)
				utils.SetByKeys(
					s.X+i,
					s.Y+y,
					utils.FgRgb(
						s.Cursor.Color["r"],
						s.Cursor.Color["g"],
						s.Cursor.Color["b"],
						s.Cursor.Symbol,
					),
					screen,
				)
			}
		}
	}

	return screen
}
