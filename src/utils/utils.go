package utils

import (
	"fmt"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"strings"
)

const Reset = "\u001B[0m"

type Screen interface {
	GetPixels() [][]string
}

func FgRgb(c pixel.Color, symbol string) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s", c.R, c.G, c.B, symbol)
}

func DrawString(X int, Y int, val string, color pixel.Color, screen [][]string) [][]string {
	str := strings.Split(val, "")
	for k, symbol := range str {
		SetByKeys(X+k, Y, symbol, color, screen)
	}

	return screen
}

func Isset(arr [][]string, first int, second int) bool {
	return first > 0 && second > 0 && first < len(arr) && second < len(arr[first])
}

func SetByKeys(X int, Y int, val string, c pixel.Color, screen [][]string) [][]string {
	if Isset(screen, Y, X) {
		screen[Y][X] = FgRgb(c, val)
	}

	return screen
}

func InArray[T comparable](searchValue T, ar []T) bool {
	for _, v := range ar {
		if searchValue == v {
			return true
		}
	}

	return false
}
