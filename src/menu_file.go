package src

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const fileX = 2

func fileList(s *Screen, screen [][]string, path string) [][]string {
	files, err := os.ReadDir(path)
	if err != nil {
		s.SetMessage(err.Error())
		s.Dir = Cfg.ImageSaveDirectory
	}

	var width int
	var fileList []string
	var dirList []string
	for _, file := range files {
		fileName := file.Name()
		if len(fileName) > width {
			width = len(fileName)
		}
		if file.IsDir() && (Cfg.ShowHiddenFolder || string(fileName[0]) != ".") {
			dirList = append(dirList, fileName)
			continue
		}
		if filepath.Ext(fileName) == ".txt" {
			fileList = append(fileList, fileName)
		}
	}
	s.FileListWidth = width + 6
	ClearMenu(s, screen, s.FileListWidth)
	str := "File " + strings.Repeat("─", s.FileListWidth-len("File ")) + "┐"
	DrawString(0, 0, str, screen)

	Y := 2
	if Cfg.ShowFolder {
		s.FileList = make(map[int]string, len(fileList)+len(dirList)+1)
		s.FileList[1] = "../"
		DrawString(fileX, 1, "..", screen)
		for _, dirName := range dirList {
			DrawString(fileX, Y, fmt.Sprintf("🗀 %v", dirName), screen)
			s.FileList[Y] = dirName + "/"
			Y++
		}
	} else {
		s.FileList = make(map[int]string, len(fileList)+1)
	}
	for y, fileName := range fileList {
		DrawString(fileX, Y+y, fmt.Sprintf("🖹 %v", fileName), screen)
		s.FileList[Y+y] = fileName
	}

	return screen
}

func saveImage(image string, s *Screen) {
	f, err := os.Create(Cfg.ImageSaveDirectory + time.Now().Format(Cfg.ImageSaveNameFormat))
	if err != nil {
		s.SetMessage(err.Error())
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			s.SetMessage(err.Error())
		}
	}(f)
	lines := strings.Split(image, "\n")
	var newImage string
	for _, line := range lines {
		newImage += line[1:len(line)-1] + "\n"
	}
	_, err = f.WriteString(newImage)
	if err != nil {
		s.SetMessage(err.Error())
	}
	s.SetMessage("Saved as " + f.Name())
}

func (s *Screen) loadImage(screenStrong string) {
	s.Pixels = []pixel{}
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
			s.Pixels = append(s.Pixels, pixel{X: x, Y: y, symbol: symbol})
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
					pixel := pixel{X: x, Y: y, symbol: str + reset}
					s.Pixels = append(s.Pixels, pixel)
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
				s.Pixels = append(s.Pixels, pixel)
				continue
			}
			str += symbol
		}
		x++
		s.Pixels = append(s.Pixels, pixel{X: x, Y: y, symbol: str})
	}

	return errors
}
