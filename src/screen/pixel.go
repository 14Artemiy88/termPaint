package screen

type Pixel struct {
	X      int
	Y      int
	Symbol string
}
type pixels []Pixel

type storePixels [2]Pixel

func (p *pixels) add(pixel ...Pixel) {
	*p = append(*p, pixel...)
}
