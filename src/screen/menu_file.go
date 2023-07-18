package screen

import (
	"fmt"
	"github.com/14Artemiy88/termPaint/src/config"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const fileX = 2

func FileList(s *Screen, screen [][]string, path string) [][]string {
	files, err := os.ReadDir(path)
	if err != nil {
		s.SetMessage(err.Error())
		s.Dir = config.Cfg.ImageSaveDirectory
	}

	var width int
	var FileList []string
	var dirList []string
	for _, file := range files {
		fileName := file.Name()
		if len(fileName) > width {
			width = len(fileName)
		}
		if file.IsDir() && (config.Cfg.ShowHiddenFolder || string(fileName[0]) != ".") {
			dirList = append(dirList, fileName)
			continue
		}
		if filepath.Ext(fileName) == ".txt" {
			FileList = append(FileList, fileName)
		}
	}
	s.FileListWidth = width + 6
	ClearMenu(s, screen, s.FileListWidth)
	str := "File " + strings.Repeat("─", s.FileListWidth-len("File ")) + "┐"
	DrawString(0, 0, str, screen)

	Y := 2
	if config.Cfg.ShowFolder {
		s.FileList = make(map[int]string, len(FileList)+len(dirList)+1)
		s.FileList[1] = "../"
		DrawString(fileX, 1, "..", screen)
		for _, dirName := range dirList {
			DrawString(fileX, Y, fmt.Sprintf("🗀 %v", dirName), screen)
			s.FileList[Y] = dirName + "/"
			Y++
		}
	} else {
		s.FileList = make(map[int]string, len(FileList)+1)
	}
	for y, fileName := range FileList {
		DrawString(fileX, Y+y, fmt.Sprintf("🖹 %v", fileName), screen)
		s.FileList[Y+y] = fileName
	}

	return screen
}

func SaveImage(image string, s *Screen) {
	f, err := os.Create(config.Cfg.ImageSaveDirectory + time.Now().Format(config.Cfg.ImageSaveNameFormat))
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
