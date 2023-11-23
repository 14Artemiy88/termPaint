package main

import (
	"fmt"
	"github.com/14Artemiy88/termPaint/src/utils"
	"log"
	"os"
	"os/exec"

	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/message"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/screen"
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
	comma := utils.FgRgb(pixel.White, ",") + utils.FgRgb(pixel.Green, " ")
	fmt.Println("Drawing in the terminal")
	fmt.Println()
	fmt.Println(utils.FgRgb(pixel.Yellow, "KEYS"))
	fmt.Println(utils.FgRgb(pixel.Green, "      ESC"+comma+"Ctrl+C         ") + utils.FgRgb(pixel.White, "Exit"))
	fmt.Println(utils.FgRgb(pixel.Green, "      Tab"+comma+"F2             ") + utils.FgRgb(pixel.White, "Menu"))
	fmt.Println(utils.FgRgb(pixel.Green, "      Ctrl+S              ") + utils.FgRgb(pixel.White, "Save in txt file"))
	fmt.Println(utils.FgRgb(pixel.Green, "      Ctrl+O"+comma+"F3          ") + utils.FgRgb(pixel.White, "Load Image"))
	fmt.Println(utils.FgRgb(pixel.Green, "      Ctrl-H"+comma+"F1          ") + utils.FgRgb(pixel.White, "Help menu"))
	fmt.Println(utils.FgRgb(pixel.Green, "      Any char            ") + utils.FgRgb(pixel.White, "Set as a Symbol"))
	fmt.Println(utils.FgRgb(pixel.Green, "      F3                  ") + utils.FgRgb(pixel.White, "SHape menu"))

	fmt.Println()

	fmt.Println(utils.FgRgb(pixel.Yellow, "MOUSE"))
	fmt.Println(utils.FgRgb(pixel.Green, "      Left                ") + utils.FgRgb(pixel.White, "Draw"))
	fmt.Println(utils.FgRgb(pixel.Green, "      Right               ") + utils.FgRgb(pixel.White, "Erase"))
	fmt.Println(utils.FgRgb(pixel.Green, "      Middle              ") + utils.FgRgb(pixel.White, "Clear Screen"))
}
