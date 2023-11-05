package draw

import "github.com/14Artemiy88/termPaint/src/pixel"

type Screen interface {
	GetWidth() int
	GetPixels() [][]string
	GetPixel(y int, x int) string
	AddPixels(pixels ...pixel.Pixel)
}
