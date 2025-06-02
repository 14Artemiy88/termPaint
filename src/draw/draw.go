package draw

import (
	"fmt"

	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/menu"
	"github.com/14Artemiy88/termPaint/src/pixel"
)

func Draw(s Screen, X int, Y int) {
	switch cursor.CC.Brush {
	case cursor.Dot:
		dot(s, pixel.Pixel{Coord: pixel.Coord{X: X, Y: Y}, Color: cursor.CC.Color, Symbol: cursor.CC.Symbol})
	case cursor.GLine:
		gLine(s, X, Y)
	case cursor.VLine:
		vLine(s, X, Y)
	case cursor.ESquare:
		eSquare(s, X, Y)
	case cursor.FSquare:
		fSquare(s, X, Y)
	case cursor.ECircle:
		eCircle(s, X, Y)
	case cursor.FCircle:
		fCircle(s, X, Y)
	case cursor.Fill:
		menu.Type = menu.None
		changedSymbols := make(map[string]pixel.Coord)
		key := fmt.Sprintf("%d-%d", Y, X)
		changedSymbols[key] = pixel.Coord{X: X, Y: Y}
		fill(s, s.GetPixel(Y, X), changedSymbols, s.GetWidth())
	case cursor.ContinuousLine, cursor.SmoothContinuousLine, cursor.FatContinuousLine, cursor.DoubleContinuousLine:
		continuousLineNew(s, X, Y)
	case cursor.Empty:
	case cursor.Pointer:
	}
}
