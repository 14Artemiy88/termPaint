package pixel

import "github.com/14Artemiy88/termPaint/src/color"

const Ratio = .4583333333333333

type Pixel struct {
	X      int
	Y      int
	Color  color.Color
	Symbol string
}
type pixels []Pixel

var Pixels pixels

var StorePixel [2]Pixel

func (p *pixels) Add(pixel ...Pixel) {
	*p = append(*p, pixel...)
}
