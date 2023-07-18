package src

import (
	tea "github.com/charmbracelet/bubbletea"
	"strings"
	"time"
)

type Screen struct {
	X             int
	Y             int
	Columns       int
	Rows          int
	Cursor        string
	NewCursor     Cursor
	Pixels        []pixel
	Color         map[string]int
	ShowMenu      bool
	ShowHelp      bool
	ShowFile      bool
	FileList      map[int]string
	FileListWidth int
	Save          bool
	InputLock     bool
	Input         string
	InputColor    string
	CursorStore   string
	File          string
	Messages      []message
	MessageWidth  int
	Dir           string
}

type pixel struct {
	X      int
	Y      int
	symbol string
}

func (s *Screen) Init() tea.Cmd {
	return tick
}

func (s *Screen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		var delCount int
		for k, m := range s.Messages {
			if m.liveTime > 0 {
				s.Messages[k].liveTime--
			} else {
				delCount++
			}
		}
		s.Messages = s.Messages[delCount:]

		return s, tick

	case tea.KeyMsg:
		return keyBind(msg, s)

	case tea.MouseMsg:
		mouseBind(msg, s)

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

	// draw empty Screen
	for i := 0; i < s.Rows; i++ {
		screen[i] = strings.Split(strings.Repeat(" ", s.Columns), "")
	}

	for _, p := range s.Pixels {
		SetByKeys(p.X, p.Y, p.symbol, screen)
	}
	if s.ShowMenu {
		drawMenu(s, screen)
	}
	if s.ShowHelp {
		drawHelpMenu(s, screen)
	}
	if s.ShowFile {
		fileList(s, screen, s.Dir)
	}
	if len(s.Messages) > 0 {
		DrawMsg(s.Messages, s.MessageWidth, screen)
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
		saveImage(screenString, s)
	}

	return screenString
}

type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Millisecond * 10)
	return tickMsg{}
}
