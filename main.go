package main

import (
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/message"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/screen"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
)

func main() {

	s := &screen.Screen{
		UnsavedPixels: map[string]pixel.Pixel{},
	}
	config.InitConfig(s)
	cursor.CC = cursor.NewCursor(s)

	p := tea.NewProgram(
		s,
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	)

	if _, err := os.Stat(s.Config.GetImageSaveDirectory()); os.IsNotExist(err) {
		errDir := os.MkdirAll(s.Config.GetImageSaveDirectory(), 0755)
		if errDir != nil {
			message.SetMessage(err.Error())
		}
		message.SetMessage("Directory " + s.Config.GetImageSaveDirectory() + " successfully created.")
	}

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
