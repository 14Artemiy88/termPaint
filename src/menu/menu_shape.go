package menu

import (
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/utils"
	"strconv"
	"strings"
)

const MenuShapeWidth = 12

var ShapeList = map[int]ShapeStruct{
	3:  {ShapeType: cursor.Dot, ShapeSymbol: "•"},
	5:  {ShapeType: cursor.GLine, ShapeSymbol: "━"},
	7:  {ShapeType: cursor.VLine, ShapeSymbol: "┃"},
	9:  {ShapeType: cursor.ESquare, ShapeSymbol: "🞎"},
	11: {ShapeType: cursor.FSquare, ShapeSymbol: "■"},
	13: {ShapeType: cursor.ECircle, ShapeSymbol: "○"},
	15: {ShapeType: cursor.FCircle, ShapeSymbol: "●"},
}

type ShapeStruct struct {
	ShapeType   cursor.Type
	ShapeSymbol string
}

func drawShapeMenu(screen [][]string) [][]string {
	ClearMenu(screen, MenuShapeWidth)
	str := "Shape " + strings.Repeat("─", MenuShapeWidth-len("Shape")-2) + "┐"
	utils.DrawString(1, 1, str, screen)

	for y, sh := range ShapeList {
		utils.DrawString(3, y, sh.ShapeSymbol, screen)
	}
	switch cursor.CC.Store.Brush {
	case cursor.GLine, cursor.VLine:
		utils.DrawString(1, 17, "Length: "+strconv.Itoa(cursor.CC.Width), screen)
	case cursor.ESquare, cursor.FSquare:
		utils.DrawString(1, 17, "Width: "+strconv.Itoa(cursor.CC.Width), screen)
		utils.DrawString(1, 18, "Height: "+strconv.Itoa(cursor.CC.Height), screen)
	case cursor.ECircle, cursor.FCircle:
		utils.DrawString(1, 17, "Radius: "+strconv.Itoa(cursor.CC.Width), screen)
	}

	return screen
}
