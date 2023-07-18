package utils

import "strconv"

const Reset = "\u001B[0m"

func FgRgb(r int, g int, b int, symbol string) string {
	if r == 255 && g == 255 && b == 255 {
		return symbol
	}
	return "\033[38;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m" + symbol + Reset
}

func isset(arr [][]string, first int, second int) bool {
	return first < len(arr) && second < len(arr[first])
}

func SetByKeys(X int, Y int, val string, screen [][]string) [][]string {
	if isset(screen, Y, X) {
		screen[Y][X] = val
	}

	return screen
}
