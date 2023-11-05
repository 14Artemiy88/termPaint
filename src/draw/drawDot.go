package draw

import "github.com/14Artemiy88/termPaint/src/pixel"

func dot(s Screen, p pixel.Pixel) {
	s.AddPixels(p)
}
