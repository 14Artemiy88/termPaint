package draw

import (
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/pixel"
)

func eSquare(s Screen, x int, y int) {
	for i := 0; i < cursor.CC.Height; i++ {
		for j := 0; j < cursor.CC.Width; j++ {
			if j > 0 && j < cursor.CC.Width-1 && i > 0 && i < cursor.CC.Height-1 {
				continue
			}

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
