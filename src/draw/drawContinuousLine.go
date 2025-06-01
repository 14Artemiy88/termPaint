package draw

import (
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/utils"
)

func continuousLineNew(s Screen, X int, Y int) {
	lineMap := map[string]string{
		"": "─",

		"u": "│",
		"d": "│",
		"l": "─",
		"r": "─",

		"ud": "│",
		"lr": "─",

		"ul": "┘",
		"ur": "└",
		"dl": "┐",
		"dr": "┌",

		"udl": "┤",
		"udr": "├",
		"ulr": "┴",
		"dlr": "┬",

		"udlr": "┼",
	}

	var line string
	// сверху
	if s.GetPixel(Y-1, X) != " " {
		line += "u"
	}
	// снизу
	if s.GetPixel(Y+1, X) != " " {
		line += "d"
	}
	// слева
	if s.GetPixel(Y, X-1) != " " {
		line += "l"
	}
	// справа
	if s.GetPixel(Y, X+1) != " " {
		line += "r"
	}

	s.AddPixels(
		pixel.Pixel{
			Coord: pixel.Coord{
				X: X,
				Y: Y,
			},
			Color: cursor.CC.Color,
			Symbol: utils.FgRgb(
				cursor.CC.Color,
				lineMap[line],
			),
		},
	)
}
