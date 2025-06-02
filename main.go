package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"

	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/message"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/screen"
	"github.com/14Artemiy88/termPaint/src/utils"
	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "-h" || os.Args[1] == "--help" {
			help()

			return
		}

		if os.Args[1] == "-v" || os.Args[1] == "--version" {
			version()

			return
		}
	}

	s := &screen.Screen{
		UnsavedPixels: map[string]pixel.Pixel{},
	}

	config.InitConfig(s)
	s.Message = message.Message{
		LiveTime: s.Config.GetNotificationTime(),
	}
	cursor.CC = cursor.NewCursor(s.GetConfig())

	p := tea.NewProgram(
		s,
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	)

	if _, err := os.Stat(s.Config.GetImageSaveDirectory()); os.IsNotExist(err) {
		errDir := os.MkdirAll(s.Config.GetImageSaveDirectory(), 0755)
		if errDir != nil {
			s.Message.SetMessage(err.Error())
		}

		s.Message.SetMessage(
			"Directory " + s.Config.GetImageSaveDirectory() + " successfully created.",
		)
	}

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func version() {
	cmd := exec.Command("git", "describe", "--tags")
	res, _ := cmd.CombinedOutput()
	fmt.Println(string(res))
}

func help() {
	green := pixel.GetConstColor("green")
	yellow := pixel.GetConstColor("yellow")
	white := pixel.GetConstColor("white")

	// Функции для цветного текста
	colorText := func(c pixel.Color, s string) string {
		return utils.FgRgb(c, s)
	}
	greenText := func(s string) string { return colorText(green, s) }
	yellowText := func(s string) string { return colorText(yellow, s) }
	whiteText := func(s string) string { return colorText(white, s) }

	// Предварительно рассчитываем повторяющиеся элементы
	comma := whiteText(",") + greenText(" ")
	keyPrefix := greenText("      ")
	descPrefix := whiteText("")

	// Используем strings.Builder для построения вывода
	var b strings.Builder
	b.WriteString("\n")
	b.WriteString("Drawing in the terminal\n\n")

	// Секция KEYS
	b.WriteString(yellowText("KEYS\n"))
	b.WriteString(keyPrefix + "ESC" + comma + "Ctrl+C         " + descPrefix + "Exit\n")
	b.WriteString(keyPrefix + "Tab" + comma + "F2             " + descPrefix + "Menu\n")
	b.WriteString(keyPrefix + "Ctrl+S              " + descPrefix + "Save in txt file\n")
	b.WriteString(keyPrefix + "Ctrl+O" + comma + "F3          " + descPrefix + "Load Image\n")
	b.WriteString(keyPrefix + "Ctrl-H" + comma + "F1          " + descPrefix + "Help menu\n")
	b.WriteString(keyPrefix + "Any char            " + descPrefix + "Set as a Symbol\n")
	b.WriteString(keyPrefix + "F3                  " + descPrefix + "Shape menu\n")
	b.WriteString("\n")

	// Секция MOUSE
	b.WriteString(yellowText("MOUSE\n"))
	b.WriteString(keyPrefix + "Left                " + descPrefix + "Draw\n")
	b.WriteString(keyPrefix + "Right               " + descPrefix + "Erase\n")
	b.WriteString(keyPrefix + "Middle              " + descPrefix + "Clear Screen")
	b.WriteString("\n")

	// Выводим все одной операцией
	fmt.Println(b.String())
}
