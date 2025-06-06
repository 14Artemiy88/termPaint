package cursor

import (
	"fmt"
	"math"

	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/utils"
)

type Cursor struct {
	X      int
	Y      int
	Brush  Type
	Width  int
	Height int
	Symbol string
	Color  pixel.Color
	Store  Store
}

var CC Cursor

type Store struct {
	Symbol string
	Brush  Type
}

type Type int

const (
	Empty Type = iota
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
	Fill
)

type Screen interface {
	GetPixels() [][]string
	GetWidth() int
	GetConfig() *config.Config
}

type Config interface {
	GetColor() pixel.Color
	GetDefaultCursor() string
}

func NewCursor(cfg *config.Config) Cursor {
	defaultCursor := cfg.GetDefaultCursor()

	return Cursor{
		Symbol: defaultCursor,
		Color:  cfg.GetColor(),
		Brush:  Dot,
		Width:  4,
		Height: 4,
		Store: Store{
			Symbol: defaultCursor,
			Brush:  Dot,
		},
	}
}

func (c *Cursor) SetCursor(cursor string) {
	c.Symbol = cursor
	c.Store.Symbol = cursor
}

func (c *Cursor) DrawCursor(s Screen) [][]string {
	clr := c.Color
	screen := s.GetPixels()

	switch c.Brush {
	case Empty:
	case Pointer:
		c.X = 1
		clr = pixel.Color{R: s.GetConfig().PointerColor["r"], G: s.GetConfig().PointerColor["g"], B: s.GetConfig().PointerColor["b"]}
		utils.SetByKeys(1, c.Y, s.GetConfig().Pointer, clr, screen)
	case Fill:
		// ToDo: вынести в конфиг 3 опции:
		// 1. показывать вообще нет заливку при этом курсоре заранее
		// 2, показывать заливку по нажатию клавиши Shift
		// 3. переключать на Dot после заливки
		changedSymbols := make(map[string]pixel.Coord)
		key := fmt.Sprintf("%d-%d", c.Y, c.X)
		changedSymbols[key] = pixel.Coord{X: c.X, Y: c.Y}
		drawFillCursor(c, clr, screen[c.Y][c.X], changedSymbols, s.GetWidth(), screen)
	case Dot,
		ContinuousLine,
		SmoothContinuousLine,
		FatContinuousLine,
		DoubleContinuousLine:
		utils.SetByKeys(c.X, c.Y, c.Symbol, clr, screen)

	case GLine:
		for i := 0; i < c.Width; i++ {
			utils.SetByKeys(c.X+i, c.Y, c.Symbol, clr, screen)
		}

	case VLine:
		for i := 0; i < c.Width; i++ {
			utils.SetByKeys(c.X, c.Y+i, c.Symbol, clr, screen)
		}

	case ESquare:
		for y := 0; y < c.Height; y++ {
			for x := 0; x < c.Width; x++ {
				if x > 0 && x < c.Width-1 && y > 0 && y < c.Height-1 {
					continue
				}

				utils.SetByKeys(c.X+x, c.Y+y, c.Symbol, clr, screen)
			}
		}

	case FSquare:
		for y := 0; y < c.Height; y++ {
			for x := 0; x < c.Width; x++ {
				utils.SetByKeys(c.X+x, c.Y+y, c.Symbol, clr, screen)
			}
		}

	case ECircle:
		R := c.Width / 2
		k := 5

		for y := -R * k; y <= R*k; y++ {
			x := int(math.Sqrt(math.Pow(float64(R), 2)-math.Pow(float64(y)/float64(k), 2)) / pixel.Ratio)
			ky := int(math.Round(float64(y) / float64(k)))
			utils.SetByKeys(c.X+x, c.Y+ky, c.Symbol, clr, screen)
			utils.SetByKeys(c.X-x, c.Y+ky, c.Symbol, clr, screen)
		}

	case FCircle:
		R := c.Width / 2
		k := 5

		for y := -R * k; y <= R*k; y++ {
			x := int(math.Sqrt(math.Pow(float64(R), 2)-math.Pow(float64(y)/float64(k), 2)) / pixel.Ratio)
			ky := int(math.Round(float64(y) / float64(k)))

			for i := -x; i <= x; i++ {
				utils.SetByKeys(c.X+i, c.Y+ky, c.Symbol, clr, screen)
			}
		}
	}

	return screen
}

func drawFillCursor(
	c *Cursor,
	clr pixel.Color,
	changedSymbol string,
	changedSymbols map[string]pixel.Coord,
	N int,
	screen [][]string,
) {
	var key string

	symbols := make(map[string]pixel.Coord)

	for _, p := range changedSymbols {
		if utils.Isset(screen, p.Y+1, p.X) && screen[p.Y+1][p.X] == changedSymbol {
			key = fmt.Sprintf("%d-%d", p.Y+1, p.X)
			symbols[key] = pixel.Coord{Y: p.Y + 1, X: p.X}
		}

		if utils.Isset(screen, p.Y-1, p.X) && screen[p.Y-1][p.X] == changedSymbol {
			key = fmt.Sprintf("%d-%d", p.Y-1, p.X)
			symbols[key] = pixel.Coord{Y: p.Y - 1, X: p.X}
		}

		if utils.Isset(screen, p.Y, p.X+1) && screen[p.Y][p.X+1] == changedSymbol {
			key = fmt.Sprintf("%d-%d", p.Y+1, p.X+1)
			symbols[key] = pixel.Coord{Y: p.Y, X: p.X + 1}
		}

		if utils.Isset(screen, p.Y, p.X-1) && screen[p.Y][p.X-1] == changedSymbol {
			key = fmt.Sprintf("%d-%d", p.Y, p.X-1)
			symbols[key] = pixel.Coord{Y: p.Y, X: p.X - 1}
		}
	}

	if len(symbols) > 0 && N > 0 {
		for _, p := range symbols {
			utils.SetByKeys(p.X, p.Y, c.Symbol, clr, screen)
		}

		N--
		drawFillCursor(c, clr, changedSymbol, symbols, N, screen)
	}
}
