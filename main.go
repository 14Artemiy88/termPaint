package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"

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

	scr := &screen.Screen{
		UnsavedPixels: map[string]pixel.Pixel{},
	}

	config.InitConfig(scr)
	scr.Message = message.Message{
		LiveTime: scr.Config.GetNotificationTime(),
	}
	cursor.CC = cursor.NewCursor(scr.GetConfig())

	program := tea.NewProgram(
		scr,
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	)

	if _, err := os.Stat(scr.Config.GetImageSaveDirectory()); os.IsNotExist(err) {
		errDir := os.MkdirAll(scr.Config.GetImageSaveDirectory(), 0755)
		if errDir != nil {
			scr.Message.SetMessage(err.Error())
		}

		scr.Message.SetMessage(
			"Directory " + scr.Config.GetImageSaveDirectory() + " successfully created.",
		)
	}

	if _, err := program.Run(); err != nil {
		log.Fatal(err)
	}
}

func version() {
	cmd := exec.Command("git", "describe", "--tags")

	res, err := cmd.Output()
	if err != nil {
		// Выводим ошибку в stderr и выходим с кодом 1
		fmt.Fprintln(os.Stderr, "Error getting version:", err)
		os.Exit(1)
	}

	os.Stdout.Write(res)
}

func help() {
	green := pixel.GetConstColor(pixel.Green)
	yellow := pixel.GetConstColor(pixel.Yellow)
	white := pixel.GetConstColor(pixel.White)

	// Создаем буфер для вывода
	var buf bytes.Buffer

	// Функция для безопасной записи с обработкой ошибок
	write := func(s string) {
		if _, err := buf.WriteString(s); err != nil {
			log.Printf("Error writing help: %v", err)

			return
		}
	}

	// Функция для записи с цветом
	writeColor := func(c pixel.Color, text string) {
		write(utils.FgRgb(c, text))
	}

	// Предварительно рассчитываем повторяющиеся элементы
	comma := utils.FgRgb(white, ",") + utils.FgRgb(green, " ")
	keyPrefix := utils.FgRgb(green, "      ")
	descPrefix := utils.FgRgb(white, "")

	// Формируем вывод
	write("\nDrawing in the terminal\n\n")

	// Секция KEYS
	writeColor(yellow, "KEYS")
	write("\n")
	write(keyPrefix + "ESC" + comma + "Ctrl+C         " + descPrefix + "Exit\n")
	write(keyPrefix + "Tab" + comma + "F2             " + descPrefix + "Menu\n")
	write(keyPrefix + "Ctrl+S              " + descPrefix + "Save in txt file\n")
	write(keyPrefix + "Ctrl+O" + comma + "F3          " + descPrefix + "Load Image\n")
	write(keyPrefix + "Ctrl-H" + comma + "F1          " + descPrefix + "Help menu\n")
	write(keyPrefix + "Any char            " + descPrefix + "Set as a Symbol\n")
	write(keyPrefix + "F3                  " + descPrefix + "Shape menu\n\n")

	// Секция MOUSE
	writeColor(yellow, "MOUSE")
	write("\n")
	write(keyPrefix + "Left                " + descPrefix + "Draw\n")
	write(keyPrefix + "Right               " + descPrefix + "Erase\n")
	write(keyPrefix + "Middle              " + descPrefix + "Clear Screen\n\n")

	// Выводим результат напрямую в stdout с проверкой ошибки
	if _, err := os.Stdout.Write(buf.Bytes()); err != nil {
		log.Printf("Error writing help: %v", err)
	}
}
