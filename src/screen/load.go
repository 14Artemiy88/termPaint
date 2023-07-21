package screen

import (
	"fmt"
	"github.com/14Artemiy88/termPaint/src/utils"
	"strings"
)

func (s *Screen) LoadImage(screenStrong string) {
	s.Pixels = []Pixel{}
	lines := strings.Split(screenStrong, "\n")
	rows := len(lines)
	errors := make(map[string]string, 2)
	if rows > s.Rows {
		errors["rows"] = fmt.Sprintf("Image rows more then terminal rows (%d > %d)", rows, s.Rows)
	}
	if strings.Contains(screenStrong, "\u001B") {
		loadColored(lines, rows, s, errors)
	} else {
		loadWhite(lines, rows, s, errors)
	}
	if len(errors) > 0 {
		for _, i := range errors {
			s.SetMessage(i)
		}
	}
}

func loadWhite(lines []string, rows int, s *Screen, errors map[string]string) map[string]string {
	for y := 0; y < rows; y++ {
		line := strings.Split(lines[y], "")
		var maxX int
		for x, symbol := range line {
			if x >= s.Columns-1 {
				if maxX == 0 {
					maxX = x
					errors["columns"] = fmt.Sprintf("Image columns more then terminal columns (%d > %d)", maxX+1, s.Columns)
				}
				maxX++
			}
			s.Pixels.add(Pixel{X: x, Y: y, Symbol: symbol})
		}
	}

	return errors
}

func loadColored(lines []string, rows int, s *Screen, errors map[string]string) map[string]string {
	for y := 0; y < rows; y++ {
		line := strings.Split(lines[y], "")
		var str string
		var x int
		var skip int
		var maxX int
		for _, symbol := range line {
			if x >= s.Columns-1 {
				if maxX == 0 {
					maxX = x
					errors["columns"] = fmt.Sprintf("Image columns more then terminal columns (%d > %d)", maxX+1, s.Columns)
				}
				maxX++
			}
			if skip > 0 {
				skip--
				continue
			}
			if symbol == " " {
				x++
				continue
			}
			if symbol == "\u001B" {
				if len(str) > 0 {
					pixel := Pixel{X: x, Y: y, Symbol: str + utils.Reset}
					s.Pixels.add(pixel)
					skip = len(utils.Reset) - 1
					str = ""
					continue
				}
				str = "\u001B"
				x++
				continue
			}
			if len(str) == 0 {
				x++
				pixel := Pixel{X: x, Y: y, Symbol: symbol}
				s.Pixels.add(pixel)
				continue
			}
			str += symbol
		}
		x++
		s.Pixels.add(Pixel{X: x, Y: y, Symbol: str})
	}

	return errors
}
