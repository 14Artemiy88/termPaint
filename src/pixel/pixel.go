package pixel

const PixelRatio = .4583333333333333

type Pixel struct {
	X      int
	Y      int
	Symbol string
}
type pixels []Pixel

var Pixels pixels

var StorePixel [2]Pixel

func (p *pixels) Add(pixel ...Pixel) {
	*p = append(*p, pixel...)
}