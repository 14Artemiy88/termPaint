package menu

import (
	"strings"

	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/utils"
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

func drawLineMenu(s Screen) {
	white := pixel.GetConstColor(pixel.White)
	pixels := s.GetPixels()
	clearMenu(s, pixels, ShapeWidth)

	str := "Line " + strings.Repeat("─", LineWidth-len("Line")) + "┐"
	utils.DrawString(1, 1, str, white, pixels)

	for y, line := range LineList {
		utils.DrawString(3, y, line.LineMenu, white, pixels)
	}
}
