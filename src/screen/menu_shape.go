package screen

import (
	"github.com/14Artemiy88/termPaint/src/cursor"
	"strconv"
	"strings"
)

const MenuShapeWidth = 12

var shapeList = map[int]Shape{
	3:  {shapeType: cursor.Dot, shapeSymbol: "•"},
	5:  {shapeType: cursor.GLine, shapeSymbol: "━"},
	7:  {shapeType: cursor.VLine, shapeSymbol: "┃"},
	9:  {shapeType: cursor.ESquare, shapeSymbol: "🞎"},
	11: {shapeType: cursor.FSquare, shapeSymbol: "■"},
	13: {shapeType: cursor.ECircle, shapeSymbol: "○"},
	15: {shapeType: cursor.FCircle, shapeSymbol: "●"},
}

type Shape struct {
	shapeType   cursor.Type
	shapeSymbol string
}

func drawShapeMenu(screen [][]string) [][]string {
	ClearMenu(screen, MenuShapeWidth)
	str := "Shape " + strings.Repeat("─", MenuShapeWidth-len("Shape")-2) + "┐"
	DrawString(1, 1, str, screen)

	for y, sh := range shapeList {
		DrawString(3, y, sh.shapeSymbol, screen)
	}
	switch cursor.CC.Store.Brush {
	case cursor.GLine, cursor.VLine:
		DrawString(1, 17, "Length: "+strconv.Itoa(cursor.CC.Width), screen)
	case cursor.ESquare, cursor.FSquare:
		DrawString(1, 17, "Width: "+strconv.Itoa(cursor.CC.Width), screen)
		DrawString(1, 18, "Height: "+strconv.Itoa(cursor.CC.Height), screen)
	case cursor.ECircle, cursor.FCircle:
		DrawString(1, 17, "Radius: "+strconv.Itoa(cursor.CC.Width), screen)
	}

	return screen
}
