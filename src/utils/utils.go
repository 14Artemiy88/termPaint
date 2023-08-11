package utils

import (
	"fmt"
	"github.com/14Artemiy88/termPaint/src/color"
	"strings"
)

const Reset = "\u001B[0m"

func FgRgb(r int, g int, b int, symbol string) string {
	return fmt.Sprintf("\033[38;2;%d;%d;%dm%s", r, g, b, symbol)
}

func DrawString(X int, Y int, val string, color color.Color, screen [][]string) [][]string {
	str := strings.Split(val, "")
	for k, symbol := range str {
		SetByKeys(X+k, Y, symbol, color, screen)
	}

	return screen
}

func Isset(arr [][]string, first int, second int) bool {
	return first > 0 && second > 0 && first < len(arr) && second < len(arr[first])
}

func SetByKeys(X int, Y int, val string, c color.Color, screen [][]string) [][]string {
	if Isset(screen, Y, X) {
		screen[Y][X] = FgRgb(c.R, c.G, c.B, val)
	}

	return screen
}
