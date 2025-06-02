package main

import (
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

	comma := utils.FgRgb(white, ",") + utils.FgRgb(green, " ")

	fmt.Println("Drawing in the terminal")
	fmt.Println()
	fmt.Println(utils.FgRgb(yellow, "KEYS"))
	fmt.Println(utils.FgRgb(green, "      ESC"+comma+"Ctrl+C         ") + utils.FgRgb(white, "Exit"))
	fmt.Println(utils.FgRgb(green, "      Tab"+comma+"F2             ") + utils.FgRgb(white, "Menu"))
	fmt.Println(utils.FgRgb(green, "      Ctrl+S              ") + utils.FgRgb(white, "Save in txt file"))
	fmt.Println(utils.FgRgb(green, "      Ctrl+O"+comma+"F3          ") + utils.FgRgb(white, "Load Image"))
	fmt.Println(utils.FgRgb(green, "      Ctrl-H"+comma+"F1          ") + utils.FgRgb(white, "Help menu"))
	fmt.Println(utils.FgRgb(green, "      Any char            ") + utils.FgRgb(white, "Set as a Symbol"))
	fmt.Println(utils.FgRgb(green, "      F3                  ") + utils.FgRgb(white, "SHape menu"))

	fmt.Println()

	fmt.Println(utils.FgRgb(yellow, "MOUSE"))
	fmt.Println(utils.FgRgb(green, "      Left                ") + utils.FgRgb(white, "Draw"))
	fmt.Println(utils.FgRgb(green, "      Right               ") + utils.FgRgb(white, "Erase"))
	fmt.Println(utils.FgRgb(green, "      Middle              ") + utils.FgRgb(white, "Clear Screen"))
}
