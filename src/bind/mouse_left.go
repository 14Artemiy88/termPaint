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

func mouseLeft(xCoord int, yCoord int, screen Screen) {
	if menu.Type != menu.None && xCoord < getMenuWidth(menu.Type) {
		handleMenuAction(menu.Type, screen, xCoord, yCoord)

		return
	}

	draw.Draw(screen, xCoord, yCoord)
}

func getMenuWidth(t menu.MenuType) int {
	switch t {
	case menu.SymbolColor:
		return menu.SymbolColorWidth
	case menu.File:
		return menu.FileListWidth
	case menu.Shape:
		return menu.ShapeWidth
	case menu.Line:
		return menu.LineWidth
	default:
		return 0
	}
}

func handleMenuAction(t menu.MenuType, screen Screen, xCoord, yCoord int) {
	switch t {
	case menu.SymbolColor:
		selectColor(yCoord)
		selectSymbol(screen, xCoord, yCoord)
	case menu.File:
		selectFile(yCoord, screen)
	case menu.Shape:
		selectShape(yCoord)
	case menu.Line:
		selectLine(screen, yCoord)
	}
}

func selectLine(s Screen, yCoord int) {
	if line, ok := menu.LineList[yCoord]; ok {
		cursor.CC.Store.Brush = line.LineType
		if line.LineType == cursor.Dot {
			cursor.CC.SetCursor(s.GetConfig().DefaultCursor)
		} else {
			cursor.CC.SetCursor(line.Cursor)
		}
	}
}

func selectShape(yCoord int) {
	if sh, ok := menu.ShapeList[yCoord]; ok {
		cursor.CC.Store.Brush = sh.ShapeType
	}
}

func selectSymbol(s Screen, xCoord int, yCoord int) {
	if symbol, ok := s.GetConfig().Symbols[yCoord][xCoord]; ok {
		cursor.CC.SetCursor(symbol)

		if s.GetConfig().Notifications.SetSymbol {
			s.GetMessage().SetMessage("Set " + symbol)
		}
	}
}

func selectColor(yCoord int) {
	if c, ok := menu.Colors[yCoord]; ok {
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

func selectFile(yCoord int, screen Screen) {
	filePath, ok := menu.FileList[yCoord]
	if !ok {
		return
	}

	dir := screen.GetDirectory()
	fullPath := filepath.Join(dir, filePath) // Используем Join для корректного пути

	info, err := os.Stat(fullPath)
	if err != nil {
		screen.GetMessage().SetMessage(err.Error())

		return
	}

	if info.IsDir() {
		screen.SetDirectory(fullPath)

		return
	}

	menu.Type = menu.None
	ext := strings.ToLower(filepath.Ext(fullPath))

	switch ext {
	case ".txt":
		content, err := os.ReadFile(fullPath)
		if err != nil {
			screen.GetMessage().SetMessage(err.Error())

			return
		}

		screen.LoadImage(string(content))

	case ".jpg", ".jpeg", ".png":
		screen.LoadFromImage(fullPath)
	}
}
