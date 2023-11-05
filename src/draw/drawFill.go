package draw

import (
	"fmt"
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/utils"
)

func fill(s Screen, clr pixel.Color, changedSymbol string, changedSymbols map[string]pixel.Coord, N int) {
	var key string
	symbols := make(map[string]pixel.Coord)
	pixels := s.GetPixels()
	for _, p := range changedSymbols {
		if utils.Isset(pixels, p.Y+1, p.X) && s.GetPixel(p.Y+1, p.X) == changedSymbol {
			key = fmt.Sprintf("%d-%d", p.Y+1, p.X)
			symbols[key] = pixel.Coord{Y: p.Y + 1, X: p.X}
		}
		if utils.Isset(pixels, p.Y-1, p.X) && s.GetPixel(p.Y-1, p.X) == changedSymbol {
			key = fmt.Sprintf("%d-%d", p.Y-1, p.X)
			symbols[key] = pixel.Coord{Y: p.Y - 1, X: p.X}
		}
		if utils.Isset(pixels, p.Y, p.X+1) && s.GetPixel(p.Y, p.X+1) == changedSymbol {
			key = fmt.Sprintf("%d-%d", p.Y+1, p.X+1)
			symbols[key] = pixel.Coord{Y: p.Y, X: p.X + 1}
		}
		if utils.Isset(pixels, p.Y, p.X-1) && s.GetPixel(p.Y, p.X-1) == changedSymbol {
			key = fmt.Sprintf("%d-%d", p.Y, p.X-1)
			symbols[key] = pixel.Coord{Y: p.Y, X: p.X - 1}
		}
	}

	if len(symbols) > 0 && N > 0 {
		for _, p := range symbols {
			s.AddPixels(pixel.Pixel{Coord: p, Color: clr, Symbol: cursor.CC.Symbol})
		}
		N--
		fill(nil, clr, changedSymbol, symbols, N)
	}
}
