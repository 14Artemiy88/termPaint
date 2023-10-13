package pixel

import (
	"fmt"
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/coord"
)

const Ratio = .4583333333333333

type Pixel struct {
	Coord  coord.Coord
	Color  color.Color
	Symbol string
}
type pixels map[string]Pixel

var Pixels pixels

var StorePixel [2]Pixel

func AddPixels(pixels ...Pixel) {
	for _, pixel := range pixels {
		key := fmt.Sprintf("%d-%d", pixel.Coord.Y, pixel.Coord.X)
		Pixels[key] = pixel
	}
}
