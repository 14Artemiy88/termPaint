package color

import (
	"strconv"
)

var White = Color{R: 255, G: 255, B: 255}
var Green = Color{R: 2, G: 186, B: 31}
var Yellow = Color{R: 190, G: 175, B: 0}
var Gray = Color{R: 150, G: 150, B: 150}
var Cian = Color{R: 0, G: 200, B: 200}
var Red = Color{R: 250, G: 0, B: 0}

type Color struct {
	R int
	G int
	B int
}

func SetColor(color string) int {
	c, _ := strconv.Atoi(color)
	if c < 255 {
		return c
	}

	return 255
}

func Decrease(color int) int {
	if color > 0 {
		color--
	}

	return color
}

func Increase(color int) int {
	if color < 255 {
		color++
	}

	return color
}

func MinMaxColor(color int) int {
	if color > 0 {
		return 0
	}

	return 255
}
