package draw

import (
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/pixel"
)

func eSquare(screen Screen, xCoord int, yCoord int) {
	for i := 0; i < cursor.CC.Height; i++ {
		for j := 0; j < cursor.CC.Width; j++ {
			if j > 0 && j < cursor.CC.Width-1 && i > 0 && i < cursor.CC.Height-1 {
				continue
			}

			screen.AddPixels(
				pixel.Pixel{
					Coord:  pixel.Coord{X: xCoord + j, Y: yCoord + i},
					Color:  cursor.CC.Color,
					Symbol: cursor.CC.Symbol,
				},
			)
		}
	}
}
