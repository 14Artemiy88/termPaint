package screen

import (
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/utils"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
	"time"
)

type Screen struct {
	Save bool
}

type size struct {
	Columns int
	Rows    int
}

var Size size

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
		cursor.CC.X = msg.Width / 2
		cursor.CC.Y = msg.Height / 2
		Size.Columns = msg.Width
		Size.Rows = msg.Height
	}

	return s, nil
}

func (s *Screen) View() string {
	if Size.Rows == 0 {
		return ""
	}

	screen := make([][]string, Size.Rows)

	// draw Empty Screen
	for i := 0; i < Size.Rows; i++ {
		screen[i] = strings.Split(strings.Repeat(" ", Size.Columns), "")
	}

	for _, p := range pixel.Pixels {
		utils.SetByKeys(p.X, p.Y, p.Symbol, screen)
	}

	drawMenu(s, screen)

	if len(Msg) > 0 {
		DrawMsg(Msg, MsgWidth, screen)
	}

	if !s.Save {
		cursor.CC.DrawCursor(screen)
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
		SaveImage(screenString)
	}

	return screenString
}

type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Millisecond * 10)
	return tickMsg{}
}
