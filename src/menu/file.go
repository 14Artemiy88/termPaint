package menu

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/utils"
)

const fileX = 3

var (
	FilePath      string
	FileList      map[int]string
	FileListWidth int
)

type Message interface {
	SetMessage(string)
}

func fileMenu(s Screen) {
	white := pixel.GetConstColor("white")

	screen := s.GetPixels()

	files, err := os.ReadDir(s.GetDirectory())
	if err != nil {
		s.GetMessage().SetMessage(err.Error())
		s.SetDirectory(s.GetConfig().ImageSaveDirectory)
	}

	var width int

	var fileList []string

	var dirList []string

	for _, file := range files {
		fileName := file.Name()
		if len(fileName) > width {
			width = len(fileName)
		}

		if file.IsDir() && (s.GetConfig().ShowHiddenFolder || string(fileName[0]) != ".") {
			dirList = append(dirList, fileName)

			continue
		}

		ext := filepath.Ext(fileName)
		if ext == ".txt" || ext == ".jpg" || ext == ".png" {
			fileList = append(fileList, fileName)
		}
	}

	FileListWidth = width + 10
	ClearMenu(s, screen, FileListWidth)

	title := "FilePath"
	str := strings.Repeat("─", FileListWidth-len(title)-2) + "┐"
	utils.DrawString(1, 1, title, pixel.GetConstColor("yellow"), screen)
	utils.DrawString(len(title)+2, 1, str, pixel.GetConstColor("gray"), screen)

	Y := 3

	if s.GetConfig().ShowFolder {
		cian := pixel.GetConstColor("cian")
		FileList = make(map[int]string, len(fileList)+len(dirList)+1)
		FileList[2] = "../"

		utils.DrawString(fileX, 2, "..", white, screen)

		for _, dirName := range dirList {
			utils.DrawString(
				fileX,
				Y,
				fmt.Sprintf("\uE5FF  %v", dirName),
				cian,
				screen,
			)

			FileList[Y] = dirName + "/"
			Y++
		}
	} else {
		FileList = make(map[int]string, len(fileList)+1)
	}

	extIcon := map[string]string{
		".txt": "\uF15C",
		".png": "",
		".jpg": "",
	}

	for y, fileName := range fileList {
		ext := filepath.Ext(fileName)
		icon := extIcon[ext]
		utils.DrawString(fileX, Y+y, fmt.Sprintf("%s  %s", icon, fileName), white, screen)
		FileList[Y+y] = fileName
	}
}

func SaveImage(m Message, imageSaveDirectory string, image string) {
	fileName := imageSaveDirectory + time.Now().Format(imageSaveDirectory)
	if len(Input.Value) > 0 {
		fileName = Input.Value + ".txt"
	}

	f, err := os.Create(fileName)
	if err != nil {
		m.SetMessage(err.Error())
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			m.SetMessage(err.Error())
		}
	}(f)

	lines := strings.Split(image, "\n")

	var newImage string

	for _, line := range lines {
		newImage += line[1:len(line)-1] + "\n"
	}

	_, err = f.WriteString(newImage)
	if err != nil {
		m.SetMessage(err.Error())
	}

	m.SetMessage("Saved as " + f.Name())
}

func DrawSaveInput(screen [][]string) [][]string {
	width := 20
	fileNameLen := len(Input.Value + BlinkCursor + ".txt")

	if fileNameLen >= width {
		width = fileNameLen + 2
	}

	clearSaveInput(screen, width, 3)
	utils.DrawString(1, 1, Input.Value+BlinkCursor+".txt", pixel.GetConstColor("white"), screen)

	return screen
}

func clearSaveInput(screen [][]string, width int, height int) [][]string {
	white := pixel.GetConstColor("white")

	for y := -1; y < height; y++ {
		for x := -1; x < width; x++ {
			utils.SetByKeys(x, y, " ", white, screen)
		}

		utils.SetByKeys(width, y, "│", white, screen)
	}

	utils.DrawString(0, height, strings.Repeat("─", width)+"┘", white, screen)

	return screen
}
