package screen

import (
	"strconv"
	"strings"
)

const MenuShapeWidth = 12

var shapeList = map[int]Shape{
	3:  {shapeType: Dot, shapeSymbol: "•"},
	5:  {shapeType: GLine, shapeSymbol: "━"},
	7:  {shapeType: VLine, shapeSymbol: "┃"},
	9:  {shapeType: ESquare, shapeSymbol: "🞎"},
	11: {shapeType: FSquare, shapeSymbol: "■"},
	13: {shapeType: ECircle, shapeSymbol: "○"},
	15: {shapeType: FCircle, shapeSymbol: "●"},
}

type Shape struct {
	shapeType   cursorType
	shapeSymbol string
}

func drawShapeMenu(s *Screen, screen [][]string) [][]string {
	ClearMenu(s, screen, MenuShapeWidth)
	str := "Shape " + strings.Repeat("─", MenuShapeWidth-len("Shape")-2) + "┐"
	DrawString(1, 1, str, screen)

	for y, sh := range shapeList {
		DrawString(3, y, sh.shapeSymbol, screen)
	}
	switch CC.Store.Brush {
	case GLine, VLine:
		DrawString(1, 17, "Length: "+strconv.Itoa(CC.Width), screen)
	case ESquare, FSquare:
		DrawString(1, 17, "Width: "+strconv.Itoa(CC.Width), screen)
		DrawString(1, 18, "Height: "+strconv.Itoa(CC.Height), screen)
	case ECircle, FCircle:
		DrawString(1, 17, "Radius: "+strconv.Itoa(CC.Width), screen)
	}

	return screen
}
