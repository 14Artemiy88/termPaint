package screen

import (
	"fmt"
	"github.com/14Artemiy88/termPaint/src/utils"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"math"
	"os"
	"strings"
)

func (s *Screen) loadFromImafe(path string) {
	file, err := os.Open(path)
	if err != nil {
		SetMessage(err.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			SetMessage(err.Error())
		}
	}(file)

	img, _, err := image.Decode(file)
	if err != nil {
		SetMessage(err.Error())
	}

	bounds := img.Bounds()
	ratio := 1
	if s.Rows > s.Columns {
		if bounds.Max.X > s.Columns {
			ratio = int(math.Ceil(float64(bounds.Max.X) / float64(s.Columns)))
		}
	} else {
		if bounds.Max.Y > s.Rows {
			ratio = int(math.Ceil(float64(bounds.Max.Y)/float64(s.Rows)) / 2)
		}
	}

	Pixels = []Pixel{}
	var y int
	for i := bounds.Min.Y; i < bounds.Max.Y; i += int(float64(ratio) / pixelRatio) {
		var x int
		for j := bounds.Min.X; j < bounds.Max.X; j += ratio {
			color := img.At(j, i)
			r, g, b, _ := color.RGBA()
			symbol := utils.FgRgb(int(r/257), int(g/257), int(b/257), CC.Symbol)
			Pixels.add(Pixel{X: x, Y: y, Symbol: symbol})
			x++
		}
		fmt.Print("\n")
		y++
	}
}

func (s *Screen) LoadImage(screenStrong string) {
	Pixels = []Pixel{}
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
			SetMessage(i)
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
			Pixels.add(Pixel{X: x, Y: y, Symbol: symbol})
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
					Pixels.add(pixel)
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
				Pixels.add(pixel)
				continue
			}
			str += symbol
		}
		x++
		Pixels.add(Pixel{X: x, Y: y, Symbol: str})
	}

	return errors
}
