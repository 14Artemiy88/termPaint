package screen

import (
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/menu"
	"github.com/14Artemiy88/termPaint/src/message"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/size"
	"github.com/14Artemiy88/termPaint/src/utils"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
	"time"
)

type Screen struct {
	ShowInputSave bool
	Save          bool
}

func (s *Screen) Init() tea.Cmd {
	return tick
}

var blink = map[bool]string{
	true:  "|",
	false: " ",
}

func (s *Screen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		var delCount int
		for k, m := range message.Msg {
			if m.LiveTime > 0 {
				message.Msg[k].LiveTime--
			} else {
				delCount++
			}
		}
		message.Msg = message.Msg[delCount:]

		menu.BlinkCursor = blink[menu.BlinkPhase]
		menu.BlinkTime--
		if menu.BlinkTime == 0 {
			menu.BlinkPhase = !menu.BlinkPhase
			menu.BlinkTime = menu.DefBlinkTime
		}

		return s, tick

	case tea.KeyMsg:
		return KeyBind(msg, s)

	case tea.MouseMsg:
		MouseBind(msg, s)

	case tea.WindowSizeMsg:
		cursor.CC.X = msg.Width / 2
		cursor.CC.Y = msg.Height / 2
		size.Size.Columns = msg.Width
		size.Size.Rows = msg.Height
	}

	return s, nil
}

func (s *Screen) View() string {
	if size.Size.Rows == 0 {
		return ""
	}

	screen := make([][]string, size.Size.Rows)

	// draw Empty Screen
	for i := 0; i < size.Size.Rows; i++ {
		screen[i] = strings.Split(strings.Repeat(" ", size.Size.Columns), "")
	}

	for _, p := range pixel.Pixels {
		utils.SetByKeys(p.X, p.Y, p.Symbol, screen)
	}

	menu.DrawMenu(screen)

	if len(message.Msg) > 0 {
		message.DrawMsg(message.Msg, message.MsgWidth, screen)
	}

	if !s.Save {
		cursor.CC.DrawCursor(screen)
	}
	if s.ShowInputSave {
		menu.DrawSaveInput(screen)
	}

	var screenString string
	for i, line := range screen {
		screenString += strings.Join(line, "")
		if i < len(screen)-1 {
			screenString += "\n"
		}
	}

	if s.Save && !s.ShowInputSave {
		s.Save = false
		menu.SaveImage(screenString)
	}

	return screenString
}

type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Millisecond * 10)
	return tickMsg{}
}
