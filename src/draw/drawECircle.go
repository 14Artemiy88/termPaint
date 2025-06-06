package draw

import (
	"math"

	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/pixel"
)

func eCircle(s Screen, X int, Y int) {
	R := cursor.CC.Width / 2
	k := 5

	for y := -R * k; y <= R*k; y++ {
		x := int(math.Sqrt(math.Pow(float64(R), 2)-math.Pow(float64(y)/float64(k), 2)) / pixel.Ratio)
		ky := int(math.Round(float64(y) / float64(k)))
		s.AddPixels(
			pixel.Pixel{Coord: pixel.Coord{X: X + x, Y: Y + ky}, Color: cursor.CC.Color, Symbol: cursor.CC.Symbol},
			pixel.Pixel{Coord: pixel.Coord{X: X - x, Y: Y + ky}, Color: cursor.CC.Color, Symbol: cursor.CC.Symbol},
		)
	}
}
