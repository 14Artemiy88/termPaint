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
	SetMessage(message string)
}

func fileMenu(s Screen) {
	white := pixel.GetConstColor(pixel.White)

	pixels := s.GetPixels()

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
	clearMenu(s, pixels, FileListWidth)

	// Заголовок
	drawTitle(FileListWidth, "FilePath", 1, "┐", pixels)

	YCoord := 3

	if s.GetConfig().ShowFolder {
		cyan := pixel.GetConstColor(pixel.Cyan)
		FileList = make(map[int]string, len(fileList)+len(dirList)+1)
		FileList[2] = "../"

		utils.DrawString(fileX, 2, "..", white, pixels)

		for _, dirName := range dirList {
			utils.DrawString(
				fileX,
				YCoord,
				fmt.Sprintf("\uE5FF  %v", dirName),
				cyan,
				pixels,
			)

			FileList[YCoord] = dirName + "/"
			YCoord++
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
		utils.DrawString(fileX, YCoord+y, fmt.Sprintf("%s  %s", icon, fileName), white, pixels)
		FileList[YCoord+y] = fileName
	}
}

func SaveImage(message Message, imageSaveDirectory string, image string) {
	fileName := imageSaveDirectory + time.Now().Format(imageSaveDirectory)
	if len(Input.Value) > 0 {
		fileName = Input.Value + ".txt"
	}

	file, err := os.Create(fileName)
	if err != nil {
		message.SetMessage(err.Error())
	}

	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			message.SetMessage(err.Error())
		}
	}(file)

	lines := strings.Split(image, "\n")

	var newImage string

	for _, line := range lines {
		newImage += line[1:len(line)-1] + "\n"
	}

	_, err = file.WriteString(newImage)
	if err != nil {
		message.SetMessage(err.Error())
	}

	message.SetMessage("Saved as " + file.Name())
}

func DrawSaveInput(screen [][]string) [][]string {
	width := 20
	fileNameLen := len(Input.Value + BlinkCursor + ".txt")

	if fileNameLen >= width {
		width = fileNameLen + 2
	}

	clearSaveInput(screen, width, 3)
	utils.DrawString(1, 1, Input.Value+BlinkCursor+".txt", pixel.GetConstColor(pixel.White), screen)

	return screen
}

func clearSaveInput(screen [][]string, width int, height int) [][]string {
	white := pixel.GetConstColor(pixel.White)

	for y := -1; y < height; y++ {
		for x := -1; x < width; x++ {
			utils.SetByKeys(x, y, " ", white, screen)
		}

		utils.SetByKeys(width, y, "│", white, screen)
	}

	utils.DrawString(0, height, strings.Repeat("─", width)+"┘", white, screen)

	return screen
}
