package bind

import "C"

import (
	"os"
	"path/filepath"

	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/draw"
	"github.com/14Artemiy88/termPaint/src/menu"
	"github.com/14Artemiy88/termPaint/src/pixel"
)

func mouseLeft(X int, Y int, s Screen) {
	if menu.Type == menu.SymbolColor && X < menu.SymbolColorWidth {
		selectColor(Y)
		selectSymbol(s, X, Y)
	} else if menu.Type == menu.File && X < menu.FileListWidth {
		selectFile(Y, s)
	} else if menu.Type == menu.Shape && X < menu.ShapeWidth {
		selectShape(Y)
	} else if menu.Type == menu.Line && X < menu.LineWidth {
		selectLine(s, Y)
	} else {
		draw.Draw(s, X, Y)
	}
}

func selectLine(s Screen, Y int) {
	if line, ok := menu.LineList[Y]; ok {
		cursor.CC.Store.Brush = line.LineType
		if line.LineType == cursor.Dot {
			cursor.CC.SetCursor(s.GetConfig().DefaultCursor)
		} else {
			cursor.CC.SetCursor(line.Cursor)
		}
	}
}

func selectShape(Y int) {
	if sh, ok := menu.ShapeList[Y]; ok {
		cursor.CC.Store.Brush = sh.ShapeType
	}
}

func selectSymbol(s Screen, X int, Y int) {
	if symbol, ok := s.GetConfig().Symbols[Y][X]; ok {
		cursor.CC.SetCursor(symbol)
		if s.GetConfig().Notifications.SetSymbol {
			s.GetMessage().SetMessage("Set " + symbol)
		}
	}
}

func selectColor(Y int) {
	if c, ok := menu.Colors[Y]; ok {
		switch c {
		case "r":
			cursor.CC.Color.R = pixel.MinMaxColor(cursor.CC.Color.R)
		case "g":
			cursor.CC.Color.G = pixel.MinMaxColor(cursor.CC.Color.G)
		case "b":
			cursor.CC.Color.B = pixel.MinMaxColor(cursor.CC.Color.B)
		}
	}
}

func selectFile(Y int, s Screen) {
	if filePath, ok := menu.FileList[Y]; ok {
		dir := s.GetDirectory()
		info, err := os.Stat(dir + filePath)
		if err != nil {
			s.GetMessage().SetMessage(err.Error())
		}
		if info.IsDir() {
			s.SetDirectory(dir + filePath)
		} else {
			menu.Type = menu.None
			ext := filepath.Ext(dir + filePath)
			if ext == ".txt" {
				content, err := os.ReadFile(dir + filePath)
				if err != nil {
					s.GetMessage().SetMessage(err.Error())
				}
				s.LoadImage(string(content))
			}
			if ext == ".jpg" || ext == ".png" {
				s.LoadFromImage(dir + filePath)
			}
		}
	}
}
