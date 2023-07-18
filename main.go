package main

import (
	"github.com/14Artemiy88/termPaint/src"
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
)

func main() {
	src.InitConfig()
	c := src.Cursor{
		Symbol: src.Cfg.DefaultCursor,
		Color:  src.Cfg.DefaultColor,
		Brush:  src.Dot,
	}
	s := &src.Screen{
		Cursor:      src.Cfg.DefaultCursor,
		NewCursor:   c,
		CursorStore: src.Cfg.DefaultCursor,
		Color:       src.Cfg.DefaultColor,
		Dir:         src.Cfg.ImageSaveDirectory,
	}
	p := tea.NewProgram(
		s,
		tea.WithAltScreen(),
		tea.WithMouseAllMotion(),
	)

	if _, err := os.Stat(src.Cfg.ImageSaveDirectory); os.IsNotExist(err) {
		errDir := os.MkdirAll(src.Cfg.ImageSaveDirectory, 0755)
		if errDir != nil {
			s.SetMessage(err.Error())
		}
		s.SetMessage("Directory " + src.Cfg.ImageSaveDirectory + " successfully created.")
	}

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}
