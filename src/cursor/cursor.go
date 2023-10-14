package cursor

import (
	"fmt"
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/coord"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/size"
	"github.com/14Artemiy88/termPaint/src/utils"
	"math"
)

type Cursor struct {
	X      int
	Y      int
	Brush  Type
	Width  int
	Height int
	Symbol string
	Color  map[string]int
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

func (c *Cursor) SetCursor(cursor string) {
	c.Symbol = cursor
	c.Store.Symbol = cursor
}

func (c *Cursor) DrawCursor(screen [][]string) [][]string {
	clr := color.Color{R: c.Color["r"], G: c.Color["g"], B: c.Color["b"]}
	switch c.Brush {
	case Empty:
	case Pointer:
		c.X = 1
		clr = color.Color{R: config.Cfg.PointerColor["r"], G: config.Cfg.PointerColor["g"], B: config.Cfg.PointerColor["b"]}
		utils.SetByKeys(1, c.Y, config.Cfg.Pointer, clr, screen)
	case Fill:
		changedSymbols := make(map[string]coord.Coord)
		key := fmt.Sprintf("%d-%d", c.Y, c.X)
		changedSymbols[key] = coord.Coord{X: c.X, Y: c.Y}
		drawFillCursor(c, clr, screen[c.Y][c.X], changedSymbols, size.Size.Width, screen)
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
	clr color.Color,
	changedSymbol string,
	changedSymbols map[string]coord.Coord,
	N int,
	screen [][]string,
) {
	var key string
	symbols := make(map[string]coord.Coord)
	for _, p := range changedSymbols {
		if utils.Isset(screen, p.Y+1, p.X) && screen[p.Y+1][p.X] == changedSymbol {
			key = fmt.Sprintf("%d-%d", p.Y+1, p.X)
			symbols[key] = coord.Coord{Y: p.Y + 1, X: p.X}
		}
		if utils.Isset(screen, p.Y-1, p.X) && screen[p.Y-1][p.X] == changedSymbol {
			key = fmt.Sprintf("%d-%d", p.Y-1, p.X)
			symbols[key] = coord.Coord{Y: p.Y - 1, X: p.X}
		}
		if utils.Isset(screen, p.Y, p.X+1) && screen[p.Y][p.X+1] == changedSymbol {
			key = fmt.Sprintf("%d-%d", p.Y+1, p.X+1)
			symbols[key] = coord.Coord{Y: p.Y, X: p.X + 1}
		}
		if utils.Isset(screen, p.Y, p.X-1) && screen[p.Y][p.X-1] == changedSymbol {
			key = fmt.Sprintf("%d-%d", p.Y, p.X-1)
			symbols[key] = coord.Coord{Y: p.Y, X: p.X - 1}
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
