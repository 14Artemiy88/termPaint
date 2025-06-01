package draw

import (
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/pixel"
)

func gLine(s Screen, x int, y int) {
	for i := 0; i < cursor.CC.Width; i++ {
		s.AddPixels(
			pixel.Pixel{
				Coord:  pixel.Coord{X: x + i, Y: y},
				Color:  cursor.CC.Color,
				Symbol: cursor.CC.Symbol,
			},
		)
	}
}
