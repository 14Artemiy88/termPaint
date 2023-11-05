package pixel

const Ratio = .4583333333333333

type Pixel struct {
	Coord  Coord
	Color  Color
	Symbol string
}

var StorePixel [2]Pixel
