package screen

import (
	"fmt"
	"github.com/14Artemiy88/termPaint/src/config"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const fileX = 3

func FileList(s *Screen, screen [][]string, path string) [][]string {
	files, err := os.ReadDir(path)
	if err != nil {
		s.SetMessage(err.Error())
		Dir = config.Cfg.ImageSaveDirectory
	}

	var width int
	var fileList []string
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
		ext := filepath.Ext(fileName)
		if ext == ".txt" || ext == ".jpg" || ext == ".png" {
			fileList = append(fileList, fileName)
		}
	}
	s.FileListWidth = width + 6
	ClearMenu(s, screen, s.FileListWidth)
	str := "File " + strings.Repeat("─", s.FileListWidth-len("File")-2) + "┐"
	DrawString(1, 1, str, screen)

	Y := 3
	if config.Cfg.ShowFolder {
		s.FileList = make(map[int]string, len(fileList)+len(dirList)+1)
		s.FileList[2] = "../"
		DrawString(fileX, 2, "..", screen)
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
