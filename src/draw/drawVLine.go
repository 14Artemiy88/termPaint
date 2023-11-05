package draw

import (
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/pixel"
)

func vLine(s Screen, x int, y int, clr pixel.Color, symbol string) {
	for i := 0; i < cursor.CC.Width; i++ {
		s.AddPixels(pixel.Pixel{Coord: pixel.Coord{X: x, Y: y + i}, Color: clr, Symbol: symbol})
	}
}
