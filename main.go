package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"strconv"
	"strings"
	"time"
)

type screen struct {
	X             int
	Y             int
	columns       int
	rows          int
	cursor        string
	pixels        []pixel
	color         map[string]int
	showMenu      bool
	showHelp      bool
	showFile      bool
	fileList      map[int]string
	fileListWidth int
	save          bool
	inputLock     bool
	input         string
	inputColor    string
	cursorStore   string
	file          string
	messages      []message
	messageWidth  int
}

type pixel struct {
	X      int
	Y      int
	symbol string
}

const reset = "\u001B[0m"

func main() {
	p := tea.NewProgram(&screen{
		cursor:      "#",
		cursorStore: "#",
		color:       map[string]int{"R": 255, "G": 255, "B": 255},
	}, tea.WithAltScreen(), tea.WithMouseAllMotion())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func (s *screen) Init() tea.Cmd {
	return tick
}

func (s *screen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tickMsg:
		var delCount int
		for k, m := range s.messages {
			if m.liveTime > 0 {
				s.messages[k].liveTime--
			} else {
				delCount++
			}
		}
		s.messages = s.messages[delCount:]

		return s, tick

	case tea.KeyMsg:
		return keyBind(msg, s)

	case tea.MouseMsg:
		mouseBind(msg, s)

	case tea.WindowSizeMsg:
		s.X = msg.Width / 2
		s.Y = msg.Height / 2
		s.columns = msg.Width
		s.rows = msg.Height
	}

	return s, nil
}

func (s *screen) View() string {
	if s.rows == 0 {
		return ""
	}

	screen := make([][]string, s.rows)

	// draw empty screen
	for i := 0; i < s.rows; i++ {
		screen[i] = strings.Split(strings.Repeat(" ", s.columns), "")
	}

	for _, p := range s.pixels {
		screen[p.Y][p.X] = p.symbol
	}
	if s.showMenu {
		drawMenu(s, screen)
	}
	if s.showHelp {
		drawHelpMenu(s, screen)
	}
	if s.showFile {
		fileList(s, screen)
	}
	if len(s.messages) > 0 {
		drawMsg(s.messages, s.messageWidth, screen)
	}

	if !s.save {
		screen[s.Y][s.X] = fgRgb(s.color["R"], s.color["G"], s.color["B"], s.cursor)
	}

	var screenString string
	for i, line := range screen {
		screenString += strings.Join(line, "")
		if i < len(screen)-1 {
			screenString += "\n"
		}
	}

	if s.save {
		s.save = false
		saveImage(screenString, s)
	}

	return screenString
}

type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Millisecond * 10)
	return tickMsg{}
}

func fgRgb(r int, g int, b int, symbol string) string {
	return "\033[38;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m" + symbol + reset
}
