package screen

import (
	"fmt"
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/coord"
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
	"strconv"
	"strings"
)

func (s *Screen) loadFromImage(path string) {
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

	pixel.Pixels = map[string]pixel.Pixel{}
	var y int
	for i := bounds.Min.Y; i < bounds.Max.Y; i += int(float64(ratio) / pixel.Ratio) {
		var x int
		for j := bounds.Min.X; j < bounds.Max.X; j += ratio {
			clr := img.At(j, i)
			r, g, b, _ := clr.RGBA()
			symbol := utils.FgRgb(color.Color{R: int(r / 257), G: int(g / 257), B: int(b / 257)}, cursor.CC.Symbol)
			pixel.AddPixels(pixel.Pixel{Coord: coord.Coord{X: x, Y: y}, Symbol: symbol})
			x++
		}
		fmt.Print("\n")
		y++
	}
}

func (s *Screen) LoadImage(screenString string) {
	pixel.Pixels = map[string]pixel.Pixel{}
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
			pixel.AddPixels(pixel.Pixel{Coord: coord.Coord{X: x, Y: y}, Color: color.White, Symbol: symbol})
		}
	}

	return errors
}

func loadColored(lines []string, rows int, errors map[string]string) map[string]string {
	clr := color.Color{}
	var symbol string
	var err error
	for y := 0; y < rows; y++ {
		line := strings.Replace(lines[y], utils.Reset, "", -1)
		symbolWithColorCode := strings.Split(line, "[38;2;")
		x := 1
		for _, part := range symbolWithColorCode {
			if len(strings.TrimSpace(part)) == 0 {
				for ; x < len(part); x++ {
					pixel.AddPixels(pixel.Pixel{Coord: coord.Coord{X: x, Y: y}, Color: clr, Symbol: " "})
				}
				continue
			}
			colors := strings.Split(part, ";")
			if len(colors) < 3 {
				continue
			}
			clr.R, err = strconv.Atoi(colors[0])
			if err != nil {
				message.SetMessage(err.Error())
			}
			clr.G, err = strconv.Atoi(colors[1])
			if err != nil {
				message.SetMessage(err.Error())
			}
			colorNsymbol := strings.Split(colors[2], "m")
			clr.B, err = strconv.Atoi(colorNsymbol[0])
			if err != nil {
				message.SetMessage(err.Error())
			}
			symbol = colorNsymbol[1]
			lenSymbol := len(symbol)
			trimSymbol := strings.TrimSpace(symbol)
			if symbol != trimSymbol {
				leTrimSymbol := len(trimSymbol)
				pixel.AddPixels(pixel.Pixel{Coord: coord.Coord{X: x, Y: y}, Color: clr, Symbol: trimSymbol})
				for j := 0; j < lenSymbol-leTrimSymbol; j++ {
					x++
					pixel.AddPixels(pixel.Pixel{Coord: coord.Coord{X: x, Y: y}, Color: clr, Symbol: " "})
				}
			} else {
				pixel.AddPixels(pixel.Pixel{Coord: coord.Coord{X: x, Y: y}, Color: clr, Symbol: symbol})
			}
			x++
		}
	}

	return errors
}
