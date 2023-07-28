package screen

import (
	"github.com/14Artemiy88/termPaint/src/utils"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
	"time"
)

type Screen struct {
	X             int
	Y             int
	Columns       int
	Rows          int
	Pixels        pixels
	StorePixel    storePixels
	MenuType      menuType
	FileList      map[int]string
	FileListWidth int
	Save          bool
	File          string
}

func (s *Screen) Init() tea.Cmd {
	return tick
}

func (s *Screen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		var delCount int
		for k, m := range Msg {
			if m.liveTime > 0 {
				Msg[k].liveTime--
			} else {
				delCount++
			}
		}
		Msg = Msg[delCount:]

		return s, tick

	case tea.KeyMsg:
		return KeyBind(msg, s)

	case tea.MouseMsg:
		MouseBind(msg, s)

	case tea.WindowSizeMsg:
		s.X = msg.Width / 2
		s.Y = msg.Height / 2
		s.Columns = msg.Width
		s.Rows = msg.Height
	}

	return s, nil
}

func (s *Screen) View() string {
	if s.Rows == 0 {
		return ""
	}

	screen := make([][]string, s.Rows)

	// draw Empty Screen
	for i := 0; i < s.Rows; i++ {
		screen[i] = strings.Split(strings.Repeat(" ", s.Columns), "")
	}

	for _, p := range s.Pixels {
		utils.SetByKeys(p.X, p.Y, p.Symbol, screen)
	}

	drawMenu(s, screen)

	if len(Msg) > 0 {
		DrawMsg(Msg, MsgWidth, screen)
	}

	if !s.Save {
		DrawCursor(s, screen)
	}

	var screenString string
	for i, line := range screen {
		screenString += strings.Join(line, "")
		if i < len(screen)-1 {
			screenString += "\n"
		}
	}

	if s.Save {
		s.Save = false
		SaveImage(screenString, s)
	}

	return screenString
}

type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Millisecond * 10)
	return tickMsg{}
}
