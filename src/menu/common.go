package menu

import (
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/utils"
	"strings"
	"sync"
)

func drawTitle(width int, title string, yCoord int, end string, pixels [][]string) {
	var (
		yellow, gray pixel.Color
		once         sync.Once
	)

	once.Do(func() {
		yellow = pixel.GetConstColor(pixel.Yellow)
		gray = pixel.GetConstColor(pixel.Gray)
	})

	titleLen := len(title)
	sepLen := width - titleLen - 2
	utils.DrawString(1, yCoord, title, yellow, pixels)
	utils.DrawString(titleLen+2, yCoord, strings.Repeat("─", sepLen)+end, gray, pixels)
}
