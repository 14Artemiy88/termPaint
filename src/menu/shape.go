package menu

import (
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/utils"
	"strconv"
)

const ShapeWidth = 15

var ShapeList = map[int]ShapeStruct{
	3:  {ShapeType: cursor.Dot, ShapeSymbol: "\uF444"},
	5:  {ShapeType: cursor.GLine, ShapeSymbol: "━━"},
	7:  {ShapeType: cursor.VLine, ShapeSymbol: "┃"},
	9:  {ShapeType: cursor.ESquare, ShapeSymbol: "\uEA72"},
	11: {ShapeType: cursor.FSquare, ShapeSymbol: "\U000F0764"},
	13: {ShapeType: cursor.ECircle, ShapeSymbol: "\uEABC"},
	15: {ShapeType: cursor.FCircle, ShapeSymbol: "\uF111"},
	17: {ShapeType: cursor.Fill, ShapeSymbol: "\U000F0266"},
}

type ShapeStruct struct {
	ShapeType   cursor.Type
	ShapeSymbol string
}

func drawShapeMenu(s Screen) {
	white := pixel.GetConstColor(pixel.White)
	green := pixel.GetConstColor(pixel.Green)

	pixels := s.GetPixels()
	ClearMenu(s, pixels, ShapeWidth)

	// Заголовок
	drawTitle(ShapeWidth, "Shape", 1, "┐", pixels)

	for y, sh := range ShapeList {
		utils.DrawString(3, y, sh.ShapeSymbol, white, pixels)
	}

	switch cursor.CC.Store.Brush {
	case cursor.GLine, cursor.VLine:
		utils.DrawString(1, 19, "Length: "+strconv.Itoa(cursor.CC.Width), green, pixels)
		utils.DrawString(len("Length:")+2, 19, strconv.Itoa(cursor.CC.Width), white, pixels)
	case cursor.ESquare, cursor.FSquare:
		utils.DrawString(1, 19, "Width: "+strconv.Itoa(cursor.CC.Width), green, pixels)
		utils.DrawString(1, 20, "Height: "+strconv.Itoa(cursor.CC.Height), green, pixels)
		utils.DrawString(len("Width:")+2, 19, strconv.Itoa(cursor.CC.Width), white, pixels)
		utils.DrawString(len("Height:")+2, 20, strconv.Itoa(cursor.CC.Height), white, pixels)
	case cursor.ECircle, cursor.FCircle:
		utils.DrawString(1, 19, "Radius: ", green, pixels)
		utils.DrawString(len("Radius:")+2, 19, strconv.Itoa(cursor.CC.Width), white, pixels)
	case cursor.Empty:
	case cursor.Pointer:
	case cursor.Dot:
	case cursor.ContinuousLine:
	case cursor.SmoothContinuousLine:
	case cursor.FatContinuousLine:
	case cursor.DoubleContinuousLine:
	case cursor.Fill:
	}
}
