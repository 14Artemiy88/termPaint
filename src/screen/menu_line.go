package screen

import (
	"strings"
)

const MenuLineWidth = 10

var menuLineList = map[int]Line{
	3: {
		LineType: Dot,
		LineMenu: "•",
	},
	5: {
		LineType: SmoothContinuousLine,
		LineMenu: "╭─╯",
		Cursor:   "─",
	},
	7: {
		LineType: ContinuousLine,
		LineMenu: "┌─┘",
		Cursor:   "─",
	},
	9: {
		LineType: FatContinuousLine,
		LineMenu: "┏━┛",
		Cursor:   "━",
	},
	11: {
		LineType: DoubleContinuousLine,
		LineMenu: "╔═╝",
		Cursor:   "═",
	},
}

var lineList = map[route]map[route]string{
	upLeft:    {up: "└", left: "┐"},
	upRight:   {up: "┘", right: "┌"},
	downLeft:  {down: "┌", left: "┘"},
	downRight: {down: "┐", right: "└"},
	stay:      {stay: "O", right: "─", left: "─", up: "│", down: "│"},
}
var smoothLineList = map[route]map[route]string{
	upLeft:    {up: "╰", left: "╮"},
	upRight:   {up: "╯", right: "╭"},
	downLeft:  {down: "╭", left: "╯"},
	downRight: {down: "╮", right: "╰"},
	stay:      {stay: "O", right: "─", left: "─", up: "│", down: "│"},
}

var fatLineList = map[route]map[route]string{
	upLeft:    {up: "┗", left: "┓"},
	upRight:   {up: "┛", right: "┏"},
	downLeft:  {down: "┏", left: "┛"},
	downRight: {down: "┓", right: "┗"},
	stay:      {stay: "O", right: "━", left: "━", up: "┃", down: "┃"},
}

var doubleLineList = map[route]map[route]string{
	upLeft:    {up: "╚", left: "╗"},
	upRight:   {up: "╝", right: "╔"},
	downLeft:  {down: "╔", left: "╝"},
	downRight: {down: "╗", right: "╚"},
	stay:      {stay: "O", right: "═", left: "═", up: "║", down: "║"},
}

var drawLine = map[cursorType]map[route]map[route]string{
	ContinuousLine:       lineList,
	SmoothContinuousLine: smoothLineList,
	FatContinuousLine:    fatLineList,
	DoubleContinuousLine: doubleLineList,
}

var gvLine = map[string]map[string]string{
	"─": {"v": "│", "g": "─"},
	"━": {"v": "┃", "g": "━"},
	"═": {"v": "║", "g": "═"},
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
	// Y: X
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

	for y, line := range menuLineList {
		DrawString(3, y, line.LineMenu, screen)
	}

	return screen
}
