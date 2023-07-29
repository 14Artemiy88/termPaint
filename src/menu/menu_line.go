package menu

import (
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/utils"
	"strings"
)

const LineWidth = 10

type LineStruct struct {
	LineType cursor.Type
	LineMenu string
	Cursor   string
}

var LineList = map[int]LineStruct{
	3: {
		LineType: cursor.Dot,
		LineMenu: "•",
	},
	5: {
		LineType: cursor.SmoothContinuousLine,
		LineMenu: "╭─╯",
		Cursor:   "─",
	},
	7: {
		LineType: cursor.ContinuousLine,
		LineMenu: "┌─┘",
		Cursor:   "─",
	},
	9: {
		LineType: cursor.FatContinuousLine,
		LineMenu: "┏━┛",
		Cursor:   "━",
	},
	11: {
		LineType: cursor.DoubleContinuousLine,
		LineMenu: "╔═╝",
		Cursor:   "═",
	},
}

var lineList = map[Route]map[Route]string{
	upLeft:    {up: "└", left: "┐"},
	upRight:   {up: "┘", right: "┌"},
	downLeft:  {down: "┌", left: "┘"},
	downRight: {down: "┐", right: "└"},
	stay:      {stay: "O", right: "─", left: "─", up: "│", down: "│"},
}
var smoothLineList = map[Route]map[Route]string{
	upLeft:    {up: "╰", left: "╮"},
	upRight:   {up: "╯", right: "╭"},
	downLeft:  {down: "╭", left: "╯"},
	downRight: {down: "╮", right: "╰"},
	stay:      {stay: "O", right: "─", left: "─", up: "│", down: "│"},
}

var fatLineList = map[Route]map[Route]string{
	upLeft:    {up: "┗", left: "┓"},
	upRight:   {up: "┛", right: "┏"},
	downLeft:  {down: "┏", left: "┛"},
	downRight: {down: "┓", right: "┗"},
	stay:      {stay: "O", right: "━", left: "━", up: "┃", down: "┃"},
}

var doubleLineList = map[Route]map[Route]string{
	upLeft:    {up: "╚", left: "╗"},
	upRight:   {up: "╝", right: "╔"},
	downLeft:  {down: "╔", left: "╝"},
	downRight: {down: "╗", right: "╚"},
	stay:      {stay: "O", right: "═", left: "═", up: "║", down: "║"},
}

var DrawLine = map[cursor.Type]map[Route]map[Route]string{
	cursor.ContinuousLine:       lineList,
	cursor.SmoothContinuousLine: smoothLineList,
	cursor.FatContinuousLine:    fatLineList,
	cursor.DoubleContinuousLine: doubleLineList,
}

var GVLine = map[string]map[string]string{
	"─": {"v": "│", "g": "─"},
	"━": {"v": "┃", "g": "━"},
	"═": {"v": "║", "g": "═"},
}

type Route int

const (
	stay Route = iota
	up
	down
	left
	right
	upLeft
	upRight
	downLeft
	downRight
)

var GetRoute = map[int]map[int]Route{
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

func drawLineMenu(screen [][]string) [][]string {
	ClearMenu(screen, ShapeWidth)
	str := "Line " + strings.Repeat("─", LineWidth-len("Line")) + "┐"
	utils.DrawString(1, 1, str, screen)

	for y, line := range LineList {
		utils.DrawString(3, y, line.LineMenu, screen)
	}

	return screen
}
