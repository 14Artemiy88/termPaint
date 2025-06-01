package draw

import (
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/pixel"
)

func fSquare(s Screen, x int, y int) {
	for i := 0; i < cursor.CC.Height; i++ {
		for j := 0; j < cursor.CC.Width; j++ {
			s.AddPixels(
				pixel.Pixel{
					Coord:  pixel.Coord{X: x + j, Y: y + i},
					Color:  cursor.CC.Color,
					Symbol: cursor.CC.Symbol,
				},
			)
		}
	}
}
