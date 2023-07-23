package screen

import (
	"strings"
)

const MenuLineWidth = 10

var lineList = map[int]Line{
	//3: {
	//	LineType: SmoothContinuousLine,
	//	LineMenu: "╭─╯",
	//	Cursor:   "⁓",
	//},
	5: {
		LineType: ContinuousLine,
		LineMenu: "┌─┘",
		Cursor:   "─",
	},
	//7: {
	//	LineType: FatContinuousLine,
	//	LineMenu: "┏━┛",
	//	Cursor:   "━",
	//},
	//9: {
	//	LineType: DoubleContinuousLine,
	//	LineMenu: "╔═╝",
	//	Cursor:   "═",
	//},
}

var drawLineList = map[route]map[route]string{
	upLeft: {
		up:   "└",
		left: "┐",
	},
	upRight: {
		up:    "┘",
		right: "┌",
	},
	downLeft: {
		down: "┌",
		left: "┘",
	},
	downRight: {
		down:  "┐",
		right: "└",
	},
	stay: {
		stay:  "O", //"─",
		right: "─", //"─",
		left:  "─", //"─",
		up:    "│", //"│",
		down:  "│", //"│",
	},
}

type route int

const (
	stay route = iota
	up
	down
	left
	right
	upLeft
	upRight
	downLeft
	downRight
)

var getRoute = map[int]map[int]route{
	// Y   X
	-1: {
		-1: upLeft,
		1:  upRight,
		0:  up,
	},
	1: {
		-1: downLeft,
		1:  downRight,
		0:  down,
	},
	0: {
		-1: left,
		1:  right,
		0:  stay,
	},
}

type Line struct {
	LineType cursorType
	LineMenu string
	Cursor   string
}

func drawLineMenu(s *Screen, screen [][]string) [][]string {
	ClearMenu(s, screen, MenuShapeWidth)
	str := "Line " + strings.Repeat("─", MenuLineWidth-len("Line")) + "┐"
	DrawString(1, 1, str, screen)

	for y, line := range lineList {
		DrawString(3, y, line.LineMenu, screen)
	}

	return screen
}
