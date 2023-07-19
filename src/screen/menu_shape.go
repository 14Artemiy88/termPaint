package screen

import (
	"strconv"
	"strings"
)

const MenuShapeWidth = 12

var shapeList = map[int]Shape{
	3:  {shapeType: GLine, shapeSymbol: "━"},
	5:  {shapeType: VLine, shapeSymbol: "┃"},
	7:  {shapeType: ESquare, shapeSymbol: "🞎"},
	9:  {shapeType: FSquare, shapeSymbol: "■"},
	11: {shapeType: ECircle, shapeSymbol: "○"},
	13: {shapeType: FCircle, shapeSymbol: "●"},
}

type Shape struct {
	shapeType   cursorType
	shapeSymbol string
}

func drawShapeMenu(s *Screen, screen [][]string) [][]string {
	ClearMenu(s, screen, MenuShapeWidth)
	str := "Shape " + strings.Repeat("─", MenuShapeWidth-len("Symbol")-1) + "┐"
	DrawString(1, 1, str, screen)

	for y, sh := range shapeList {
		DrawString(3, y, sh.shapeSymbol, screen)
	}
	DrawString(1, 16, "Width: "+strconv.Itoa(s.Cursor.Width), screen)

	return screen
}
