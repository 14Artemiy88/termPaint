package draw

import (
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/utils"
)

func continuousLineNew(screen Screen, xCoord int, yCoord int) {
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
	if screen.GetPixel(yCoord-1, xCoord) != " " {
		line += "u"
	}
	// снизу
	if screen.GetPixel(yCoord+1, xCoord) != " " {
		line += "d"
	}
	// слева
	if screen.GetPixel(yCoord, xCoord-1) != " " {
		line += "l"
	}
	// справа
	if screen.GetPixel(yCoord, xCoord+1) != " " {
		line += "r"
	}

	screen.AddPixels(
		pixel.Pixel{
			Coord: pixel.Coord{
				X: xCoord,
				Y: yCoord,
			},
			Color: cursor.CC.Color,
			Symbol: utils.FgRgb(
				cursor.CC.Color,
				lineMap[line],
			),
		},
	)
}
