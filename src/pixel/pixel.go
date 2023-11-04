package pixel

import (
	"fmt"
)

const Ratio = .4583333333333333

type Pixel struct {
	Coord  Coord
	Color  Color
	Symbol string
}

var Pixels map[string]Pixel
var StorePixel [2]Pixel

func AddPixels(pixels ...Pixel) {
	for _, pixel := range pixels {
		key := fmt.Sprintf("%d-%d", pixel.Coord.Y, pixel.Coord.X)
		Pixels[key] = pixel
	}
}
