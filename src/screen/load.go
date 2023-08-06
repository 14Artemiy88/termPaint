package screen

import (
	"fmt"
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/message"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/size"
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
		message.SetMessage(err.Error())
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			message.SetMessage(err.Error())
		}
	}(file)

	img, _, err := image.Decode(file)
	if err != nil {
		message.SetMessage(err.Error())
	}

	bounds := img.Bounds()
	ratio := 1
	if size.Size.Height > size.Size.Width {
		if bounds.Max.X > size.Size.Width {
			ratio = int(math.Ceil(float64(bounds.Max.X) / float64(size.Size.Width)))
		}
	} else {
		if bounds.Max.Y > size.Size.Height {
			ratio = int(math.Ceil(float64(bounds.Max.Y)/float64(size.Size.Height)) / 2)
		}
	}

	pixel.Pixels = []pixel.Pixel{}
	var y int
	for i := bounds.Min.Y; i < bounds.Max.Y; i += int(float64(ratio) / pixel.Ratio) {
		var x int
		for j := bounds.Min.X; j < bounds.Max.X; j += ratio {
			color := img.At(j, i)
			r, g, b, _ := color.RGBA()
			symbol := utils.FgRgb(int(r/257), int(g/257), int(b/257), cursor.CC.Symbol)
			pixel.Pixels.Add(pixel.Pixel{X: x, Y: y, Symbol: symbol})
			x++
		}
		fmt.Print("\n")
		y++
	}
}

func (s *Screen) LoadImage(screenString string) {
	pixel.Pixels = []pixel.Pixel{}
	lines := strings.Split(screenString, "\n")
	rows := len(lines)
	errors := make(map[string]string, 2)
	if rows > size.Size.Height {
		errors["rows"] = fmt.Sprintf("Image rows more then terminal rows (%d > %d)", rows, size.Size.Height)
	}
	if strings.Contains(screenString, "\u001B") {
		loadColored(lines, rows, errors)
	} else {
		loadWhite(lines, rows, errors)
	}
	if len(errors) > 0 {
		for _, i := range errors {
			message.SetMessage(i)
		}
	}
}

func loadWhite(lines []string, rows int, errors map[string]string) map[string]string {
	for y := 0; y < rows; y++ {
		line := strings.Split(lines[y], "")
		var maxX int
		for x, symbol := range line {
			if x >= size.Size.Width-1 {
				if maxX == 0 {
					maxX = x
					errors["columns"] = fmt.Sprintf("Image columns more then terminal columns (%d > %d)", maxX+1, size.Size.Width)
				}
				maxX++
			}
			pixel.Pixels.Add(pixel.Pixel{X: x, Y: y, Symbol: symbol})
		}
	}

	return errors
}

func loadColored(lines []string, rows int, errors map[string]string) map[string]string {
	for y := 0; y < rows; y++ {
		line := strings.Split(lines[y], "")
		var str string
		var x int
		var skip int
		var maxX int
		for _, symbol := range line {
			if x >= size.Size.Width-1 {
				if maxX == 0 {
					maxX = x
					errors["columns"] = fmt.Sprintf("Image columns more then terminal columns (%d > %d)", maxX+1, size.Size.Width)
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
					pixel.Pixels.Add(pixel.Pixel{X: x, Y: y, Symbol: str + utils.Reset})
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
				pixel.Pixels.Add(pixel.Pixel{X: x, Y: y, Symbol: symbol})
				continue
			}
			str += symbol
		}
		x++
		pixel.Pixels.Add(pixel.Pixel{X: x, Y: y, Symbol: str})
	}

	return errors
}
