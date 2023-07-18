package screen

import (
	"github.com/14Artemiy88/termPaint/src/utils"
)

type Cursor struct {
	X      int
	Y      int
	Brush  cursorType
	Symbol string
	Color  map[string]int
}

type cursorType int

const (
	Dot cursorType = iota
	GLine
	VLine
	ESquare
	FSquare
	ECircle
	FCircle
)

func DrawCursor(s *Screen, screen [][]string) [][]string {
	screen[s.Y][s.X] = utils.FgRgb(
		s.NewCursor.Color["r"],
		s.NewCursor.Color["g"],
		s.NewCursor.Color["b"],
		s.NewCursor.Symbol,
	)

	return screen
}
