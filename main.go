package main

import (
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/screen"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
)

func main() {
	config.InitConfig()
	s := &screen.Screen{}
	screen.CC = screen.Cursor{
		Symbol: config.Cfg.DefaultCursor,
		Color:  config.Cfg.DefaultColor,
		Brush:  screen.Dot,
		Width:  4,
		Height: 4,
		Store: screen.Store{
			Symbol: config.Cfg.DefaultCursor,
			Brush:  screen.Dot,
		},
	}
	screen.Dir = config.Cfg.ImageSaveDirectory
	p := tea.NewProgram(
		s,
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	)

	if _, err := os.Stat(config.Cfg.ImageSaveDirectory); os.IsNotExist(err) {
		errDir := os.MkdirAll(config.Cfg.ImageSaveDirectory, 0755)
		if errDir != nil {
			screen.SetMessage(err.Error())
		}
		screen.SetMessage("Directory " + config.Cfg.ImageSaveDirectory + " successfully created.")
	}

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
