package draw

import (
	"fmt"

	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/pixel"
)

func fill(screen Screen, changedSymbol string, changedSymbols map[string]pixel.Coord, counter int) {
	var key string

	symbols := make(map[string]pixel.Coord)

	for _, coord := range changedSymbols {
		if screen.GetPixel(coord.Y+1, coord.X) == changedSymbol {
			key = fmt.Sprintf("%d-%d", coord.Y+1, coord.X)
			symbols[key] = pixel.Coord{Y: coord.Y + 1, X: coord.X}
		}

		if screen.GetPixel(coord.Y-1, coord.X) == changedSymbol {
			key = fmt.Sprintf("%d-%d", coord.Y-1, coord.X)
			symbols[key] = pixel.Coord{Y: coord.Y - 1, X: coord.X}
		}

		if screen.GetPixel(coord.Y, coord.X+1) == changedSymbol {
			key = fmt.Sprintf("%d-%d", coord.Y+1, coord.X+1)
			symbols[key] = pixel.Coord{Y: coord.Y, X: coord.X + 1}
		}

		if screen.GetPixel(coord.Y, coord.X-1) == changedSymbol {
			key = fmt.Sprintf("%d-%d", coord.Y, coord.X-1)
			symbols[key] = pixel.Coord{Y: coord.Y, X: coord.X - 1}
		}
	}

	if len(symbols) > 0 && counter > 0 {
		for _, p := range symbols {
			screen.AddPixels(pixel.Pixel{Coord: p, Color: cursor.CC.Color, Symbol: cursor.CC.Symbol})
		}

		counter--
		fill(screen, changedSymbol, symbols, counter)
	}
}
