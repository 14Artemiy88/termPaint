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
	Store  string
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
		s.Cursor.Color["r"],
		s.Cursor.Color["g"],
		s.Cursor.Color["b"],
		s.Cursor.Symbol,
	)

	return screen
}
