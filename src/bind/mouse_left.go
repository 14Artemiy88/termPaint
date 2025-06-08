package bind

import "C"

import (
	"os"
	"path/filepath"
	"strings"

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
	filePath, ok := menu.FileList[Y]
	if !ok {
		return
	}

	dir := s.GetDirectory()
	fullPath := filepath.Join(dir, filePath) // Используем Join для корректного пути

	info, err := os.Stat(fullPath)
	if err != nil {
		s.GetMessage().SetMessage(err.Error())

		return
	}

	if info.IsDir() {
		s.SetDirectory(fullPath)

		return
	}

	menu.Type = menu.None
	ext := strings.ToLower(filepath.Ext(fullPath))

	switch ext {
	case ".txt":
		content, err := os.ReadFile(fullPath)
		if err != nil {
			s.GetMessage().SetMessage(err.Error())

			return
		}

		s.LoadImage(string(content))

	case ".jpg", ".jpeg", ".png":
		s.LoadFromImage(fullPath)
	}
}
