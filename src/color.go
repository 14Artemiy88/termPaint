package src

import "strconv"

func setColor(color string) int {
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

func minMsxColor(color int) int {
	if color > 0 {
		return 0
	}

	return 255
}