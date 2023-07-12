package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const fileX = 2

func fileList(s *screen, screen [][]string, path string) [][]string {
	files, err := os.ReadDir(path)
	if err != nil {
		s.setMessage(err.Error())
		s.dir = cfg.DefaultDirectory
	}

	var width int
	var fileList []string
	var dirList []string
	for _, file := range files {
		fileName := file.Name()
		if len(fileName) > width {
			width = len(fileName)
		}
		if file.IsDir() && (cfg.ShowHiddenFolder || string(fileName[0]) != ".") {
			dirList = append(dirList, fileName)
			continue
		}
		if filepath.Ext(fileName) == ".txt" {
			fileList = append(fileList, fileName)
		}
	}
	s.fileListWidth = width + 6
	clearMenu(s, screen, s.fileListWidth)
	str := "File " + strings.Repeat("─", s.fileListWidth-len("File ")) + "┐"
	drawString(0, 0, str, screen)

	Y := 2
	if cfg.ShowFolder {
		s.fileList = make(map[int]string, len(fileList)+len(dirList)+1)
		s.fileList[1] = "../"
		drawString(fileX, 1, "..", screen)
		for _, dirName := range dirList {
			drawString(fileX, Y, fmt.Sprintf("🗀 %v", dirName), screen)
			s.fileList[Y] = dirName + "/"
			Y++
		}
	} else {
		s.fileList = make(map[int]string, len(fileList)+1)
	}
	for y, fileName := range fileList {
		drawString(fileX, Y+y, fmt.Sprintf("🖹 %v", fileName), screen)
		s.fileList[Y+y] = fileName
	}

	return screen
}

func saveImage(image string, s *screen) {
	f, err := os.Create(time.Now().Format("termPaint_01-02-2006_15:04:05.txt"))
	if err != nil {
		s.setMessage(err.Error())
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			s.setMessage(err.Error())
		}
	}(f)
	lines := strings.Split(image, "\n")
	var newImage string
	for _, line := range lines {
		newImage += line[1:len(line)-1] + "\n"
	}
	_, err = f.WriteString(newImage)
	if err != nil {
		s.setMessage(err.Error())
	}
	s.setMessage("Saved as " + f.Name())
}

func (s *screen) loadImage(screenStrong string) {
	s.pixels = []pixel{}
	lines := strings.Split(screenStrong, "\n")
	rows := len(lines)
	errors := make(map[string]string, 2)
	if rows > s.rows {
		errors["rows"] = fmt.Sprintf("Image rows more then terminal rows (%d > %d)", rows, s.rows)
	}
	if strings.Contains(screenStrong, "\u001B") {
		loadColored(lines, rows, s, errors)
	} else {
		loadWhite(lines, rows, s, errors)
	}
	if len(errors) > 0 {
		for _, i := range errors {
			s.setMessage(i)
		}
	}
}

func loadWhite(lines []string, rows int, s *screen, errors map[string]string) map[string]string {
	for y := 0; y < rows; y++ {
		line := strings.Split(lines[y], "")
		var maxX int
		for x, symbol := range line {
			if x >= s.columns-1 {
				if maxX == 0 {
					maxX = x
					errors["columns"] = fmt.Sprintf("Image columns more then terminal columns (%d > %d)", maxX+1, s.columns)
				}
				maxX++
			}
			s.pixels = append(s.pixels, pixel{X: x, Y: y, symbol: symbol})
		}
	}

	return errors
}

func loadColored(lines []string, rows int, s *screen, errors map[string]string) map[string]string {
	for y := 0; y < rows; y++ {
		line := strings.Split(lines[y], "")
		var str string
		var x int
		var skip int
		var maxX int
		for _, symbol := range line {
			if x >= s.columns-1 {
				if maxX == 0 {
					maxX = x
					errors["columns"] = fmt.Sprintf("Image columns more then terminal columns (%d > %d)", maxX+1, s.columns)
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
					pixel := pixel{X: x, Y: y, symbol: str + reset}
					s.pixels = append(s.pixels, pixel)
					skip = len(reset) - 1
					str = ""
					continue
				}
				str = "\u001B"
				x++
				continue
			}
			if len(str) == 0 {
				x++
				pixel := pixel{X: x, Y: y, symbol: symbol}
				s.pixels = append(s.pixels, pixel)
				continue
			}
			str += symbol
		}
		x++
		s.pixels = append(s.pixels, pixel{X: x, Y: y, symbol: str})
	}

	return errors
}
