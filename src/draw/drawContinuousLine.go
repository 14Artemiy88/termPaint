package draw

import (
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/menu"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/utils"
)

func continuousLine(s Screen, X int, Y int, clr pixel.Color) {
	x := X - pixel.StorePixel[1].Coord.X // -1 0 1
	y := Y - pixel.StorePixel[1].Coord.Y // -1 0 1
	if x != 0 || y != 0 {
		line := cursor.CC.Store.Symbol
		if pixel.StorePixel[1].Symbol != "" {
			if x < -1 || x > 1 || y < -1 || y > 1 {
				pixel.StorePixel = [2]pixel.Pixel{}
				x = 0
				y = 0
			}
			if x == 0 {
				line = menu.GVLine[line]["v"]
				cursor.CC.SetCursor(line)
			}
			if y == 0 {
				line = menu.GVLine[line]["g"]
				cursor.CC.SetCursor(line)
			}
		}

		p := pixel.Pixel{
			Coord: pixel.Coord{
				X: X,
				Y: Y,
			},
			Color: clr,
			Symbol: utils.FgRgb(
				cursor.CC.Color,
				line,
			),
		}
		s.AddPixels(p)

		var px int
		var py int
		var pr menu.Route
		if pixel.StorePixel[0].Symbol != "" {
			px = X - pixel.StorePixel[0].Coord.X
			py = Y - pixel.StorePixel[0].Coord.Y
			if px > 1 || px < -1 {
				px = 0
			}
			if py > 1 || py < -1 {
				py = 0
			}
		}
		r := menu.GetRoute[y][x]
		pr = menu.GetRoute[py][px]

		s.AddPixels(
			pixel.Pixel{
				Coord: pixel.Coord{
					X: pixel.StorePixel[1].Coord.X,
					Y: pixel.StorePixel[1].Coord.Y,
				},
				Symbol: utils.FgRgb(
					cursor.CC.Color,
					menu.DrawLine[cursor.CC.Store.Brush][pr][r],
				),
			},
		)

		pixel.StorePixel = [2]pixel.Pixel{pixel.StorePixel[1], p}
	}
}
