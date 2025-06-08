package menu

import (
	"strconv"
	"strings"
	"sync"

	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/utils"
)

const SymbolColorWidth = 15

type InputStruct struct {
	Lock  bool
	Value string
	Color string
}

var Input InputStruct

func GetColor(key int) string {
	switch key {
	case 17:
		return "r"
	case 19:
		return "g"
	case 21:
		return "b"
	default:
		return ""
	}
}

const colorX = 3

func drawSymbolColorMenu(screen Screen) {
	pixels := screen.GetPixels()
	clearMenu(screen, pixels, SymbolColorWidth)
	drawSymbolMenu(screen, pixels)
	drawColorMenu(pixels)

	title := "Help"
	str := strings.Repeat("─", SymbolColorWidth-len(title)-2) + "┤"
	height := screen.GetHeight()
	utils.DrawString(1, height-3, title, pixel.GetConstColor(pixel.Yellow), pixels)
	utils.DrawString(len(title)+2, height-3, str, pixel.GetConstColor(pixel.Gray), pixels)
	utils.DrawString(2, height-1, "Press", pixel.GetConstColor(pixel.White), pixels)
	utils.DrawString(len("Press")+3, height-1, "Ctrl+H", pixel.GetConstColor(pixel.Green), pixels)
}

func drawSymbolMenu(screen Screen, pixels [][]string) [][]string {
	// Статические переменные для кеширования
	var (
		white pixel.Color
		once  sync.Once
	)

	// Инициализация цветов один раз за вызов
	once.Do(func() {
		white = pixel.GetConstColor(pixel.White)
	})

	// Заголовок
	drawTitle(SymbolColorWidth, "Symbol", 1, "┐", pixels)

	// Оптимизируем работу с символами
	symbols := screen.GetConfig().Symbols
	for y := range symbols {
		line := symbols[y]
		for x := range line {
			// Прямой доступ к символу вместо повторного обращения через индекс
			utils.SetByKeys(x, y, line[x], white, pixels)
		}
	}

	return pixels
}

func drawColorMenu(pixels [][]string) [][]string {
	// Заголовок
	drawTitle(SymbolColorWidth, "Color", 15, "┤", pixels)

	// Статические переменные для кеширования
	var (
		white pixel.Color
		once  sync.Once
	)

	// Инициализация цветов один раз за вызов
	once.Do(func() {
		white = pixel.GetConstColor(pixel.White)
	})

	// Предварительно вычисляем значения компонент цвета
	rVal := cursor.CC.Color.R
	gVal := cursor.CC.Color.G
	bVal := cursor.CC.Color.B

	// Готовим цветные блоки
	rBlock := utils.FgRgb(pixel.Color{R: rVal}, "█")
	gBlock := utils.FgRgb(pixel.Color{G: gVal}, "█")
	bBlock := utils.FgRgb(pixel.Color{B: bVal}, "█")

	drawColorComponent(17, rVal, rBlock, white, pixels)
	drawColorComponent(19, gVal, gBlock, white, pixels)
	drawColorComponent(21, bVal, bBlock, white, pixels)

	utils.SetByKeys(3, 23, "█", cursor.CC.Color, pixels)
	utils.SetByKeys(4, 23, "█", cursor.CC.Color, pixels)
	utils.SetByKeys(5, 23, "█", cursor.CC.Color, pixels)
	utils.SetByKeys(3, 24, "█", cursor.CC.Color, pixels)
	utils.SetByKeys(4, 24, "█", cursor.CC.Color, pixels)
	utils.SetByKeys(5, 24, "█", cursor.CC.Color, pixels)

	return pixels
}

func drawColorComponent(y int, value int, block string, white pixel.Color, screen [][]string) {
	utils.DrawString(colorX+2, y, strconv.Itoa(value), white, screen)
	utils.SetByKeys(colorX, y, block, white, screen)
}
