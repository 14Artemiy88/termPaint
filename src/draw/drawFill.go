package draw

import (
	"fmt"
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/pixel"
)

func fill(s Screen, changedSymbol string, changedSymbols map[string]pixel.Coord, N int) {
	var key string
	symbols := make(map[string]pixel.Coord)
	for _, p := range changedSymbols {
		if s.GetPixel(p.Y+1, p.X) == changedSymbol {
			key = fmt.Sprintf("%d-%d", p.Y+1, p.X)
			symbols[key] = pixel.Coord{Y: p.Y + 1, X: p.X}
		}
		if s.GetPixel(p.Y-1, p.X) == changedSymbol {
			key = fmt.Sprintf("%d-%d", p.Y-1, p.X)
			symbols[key] = pixel.Coord{Y: p.Y - 1, X: p.X}
		}
		if s.GetPixel(p.Y, p.X+1) == changedSymbol {
			key = fmt.Sprintf("%d-%d", p.Y+1, p.X+1)
			symbols[key] = pixel.Coord{Y: p.Y, X: p.X + 1}
		}
		if s.GetPixel(p.Y, p.X-1) == changedSymbol {
			key = fmt.Sprintf("%d-%d", p.Y, p.X-1)
			symbols[key] = pixel.Coord{Y: p.Y, X: p.X - 1}
		}
	}

	if len(symbols) > 0 && N > 0 {
		for _, p := range symbols {
			s.AddPixels(pixel.Pixel{Coord: p, Color: cursor.CC.Color, Symbol: cursor.CC.Symbol})
		}
		N--
		fill(s, changedSymbol, symbols, N)
	}
}
