package screen

const pixelRatio = .4583333333333333

type Pixel struct {
	X      int
	Y      int
	Symbol string
}
type pixels []Pixel

var StorePixel [2]Pixel

func (p *pixels) add(pixel ...Pixel) {
	*p = append(*p, pixel...)
}
