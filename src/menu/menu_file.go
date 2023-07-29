package menu

import (
	"fmt"
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/message"
	"github.com/14Artemiy88/termPaint/src/utils"
	"os"
	"path/filepath"
	"strings"
	"time"
)

const fileX = 3

var (
	Dir           string
	FilePath      string
	FileList      map[int]string
	FileListWidth int
)

func FileMenu(screen [][]string, path string) [][]string {
	files, err := os.ReadDir(path)
	if err != nil {
		message.SetMessage(err.Error())
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
	FileListWidth = width + 6
	ClearMenu(screen, FileListWidth)
	str := "FilePath " + strings.Repeat("─", FileListWidth-len("FilePath")-2) + "┐"
	utils.DrawString(1, 1, str, screen)

	Y := 3
	if config.Cfg.ShowFolder {
		FileList = make(map[int]string, len(fileList)+len(dirList)+1)
		FileList[2] = "../"
		utils.DrawString(fileX, 2, "..", screen)
		for _, dirName := range dirList {
			utils.DrawString(fileX, Y, fmt.Sprintf("🗀 %v", dirName), screen)
			FileList[Y] = dirName + "/"
			Y++
		}
	} else {
		FileList = make(map[int]string, len(fileList)+1)
	}
	for y, fileName := range fileList {
		utils.DrawString(fileX, Y+y, fmt.Sprintf("🖹 %v", fileName), screen)
		FileList[Y+y] = fileName
	}

	return screen
}

func SaveImage(image string) {
	f, err := os.Create(config.Cfg.ImageSaveDirectory + time.Now().Format(config.Cfg.ImageSaveNameFormat))
	if err != nil {
		message.SetMessage(err.Error())
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			message.SetMessage(err.Error())
		}
	}(f)
	lines := strings.Split(image, "\n")
	var newImage string
	for _, line := range lines {
		newImage += line[1:len(line)-1] + "\n"
	}
	_, err = f.WriteString(newImage)
	if err != nil {
		message.SetMessage(err.Error())
	}
	message.SetMessage("Saved as " + f.Name())
}
