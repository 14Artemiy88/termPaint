package main

import (
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/menu"
	"github.com/14Artemiy88/termPaint/src/message"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/screen"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
)

func main() {
	config.InitConfig()
	pixel.Pixels = map[string]pixel.Pixel{}
	s := &screen.Screen{}
	cursor.CC = cursor.Cursor{
		Symbol: config.Cfg.DefaultCursor,
		Color:  config.Cfg.DefaultColor,
		Brush:  cursor.Dot,
		Width:  4,
		Height: 4,
		Store: cursor.Store{
			Symbol: config.Cfg.DefaultCursor,
			Brush:  cursor.Dot,
		},
	}
	menu.Dir = config.Cfg.ImageSaveDirectory
	p := tea.NewProgram(
		s,
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	)

	if _, err := os.Stat(config.Cfg.ImageSaveDirectory); os.IsNotExist(err) {
		errDir := os.MkdirAll(config.Cfg.ImageSaveDirectory, 0755)
		if errDir != nil {
			message.SetMessage(err.Error())
		}
		message.SetMessage("Directory " + config.Cfg.ImageSaveDirectory + " successfully created.")
	}

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
