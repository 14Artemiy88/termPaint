package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func fileList(s *screen, screen [][]string) [][]string {
	files, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	y := 1
	width := 0
	filelist := make(map[int]string)
	for _, file := range files {
		fileName := file.Name()
		if len(fileName) > width {
			width = len(fileName)
		}
		if filepath.Ext(fileName) == ".txt" {
			filelist[y] = fileName
			y++
		}
	}
	s.fileList = filelist
	s.fileListWidth = width + 4
	clearMenu(s, screen, s.fileListWidth)
	str := "File " + strings.Repeat("─", s.fileListWidth-2-len("File ")) + "┐"
	drawString(0, 0, str, screen)
	for y, fileName := range filelist {
		drawString(1, y, fileName, screen)
	}

	return screen
}

func saveImage(image string) {
	f, err := os.Create(time.Now().Format("termPaint_01-02-2006_15:04:05.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	lines := strings.Split(image, "\n")
	var newImage string
	for _, line := range lines {
		newImage += line[1:len(line)-1] + "\n"
	}
	_, err = f.WriteString(newImage)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *screen) load(screenStrong string) {
	s.pixels = []pixel{}
	lines := strings.Split(screenStrong, "\n")
	rows := len(lines)
	if rows > s.rows {
		rows = s.rows
	}
	for y := 0; y < rows; y++ {
		line := strings.Split(lines[y], "")
		var str string
		x := 0
		skip := 0
		for _, symbol := range line {
			if x >= s.columns-1 {
				break
			}
			if skip > 0 {
				skip--
				continue
			}
			if symbol == " " {
				x++
				s.pixels = append(s.pixels, pixel{X: x, Y: y, symbol: " "})
				continue
			}
			if symbol == "\u001B" {
				if len(str) > 0 {
					pixel := pixel{X: x, Y: y, symbol: str + "\u001B[0m"}
					s.pixels = append(s.pixels, pixel)
					fmt.Println(x, y, pixel)
					skip = len("\u001B[0m") - 1
					str = ""
					continue
				}
				str = "\u001B"
				x++
				continue
			}
			str += symbol
		}
		x++
		if x < s.columns {
			s.pixels = append(s.pixels, pixel{X: x, Y: y, symbol: str})
		}
	}
}
