package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type screen struct {
	X           int
	Y           int
	columns     int
	rows        int
	cursor      string
	pixels      []pixel
	color       map[string]int
	showMenu    bool
	showHelp    bool
	showFile    bool
	fileList    map[int]string
	save        bool
	inputLock   bool
	input       string
	inputColor  string
	cursorStore string
}

type pixel struct {
	X      int
	Y      int
	symbol string
}

const bold = "\u001b[1m"
const underline = "\u001b[4m"
const reverse = "\u001b[7m"

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

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return s, tea.Quit

		case tea.KeyCtrlO:
			s.showHelp = false
			s.showMenu = false
			s.showFile = !s.showFile

		case tea.KeyTab:
			s.showHelp = false
			s.showFile = false
			s.showMenu = !s.showMenu

		case tea.KeyEnter:
			s.showMenu = false
			s.showFile = false
			s.showHelp = !s.showHelp

		case tea.KeyCtrlS:
			s.save = true
			s.showMenu = false
			s.showHelp = false
			s.showFile = false

		case tea.KeyRunes:
			if s.showMenu && s.inputLock {
				if _, err := strconv.Atoi(string(msg.Runes)); err == nil {
					s.input += string(msg.Runes)
				}
			} else {
				s.cursorStore = string(msg.Runes)
				s.cursor = string(msg.Runes)
			}
		}

	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseMotion:
			xMin := 0
			if s.showMenu {
				xMin = menuWidth
			}
			if s.showHelp {
				xMin = helpWidth
			}
			if s.showFile {
				xMin = fileWidth
			}
			if s.showMenu && msg.X <= xMin {
				s.cursor = " "
				_, okS := symbols[msg.Y][msg.X]
				color, okC := colors[msg.Y][msg.X]
				if okS {
					s.X = msg.X - 1
					s.cursor = fgRgb(170, 170, 170, pointer)
				}
				if okC {
					s.inputLock = true
					s.inputColor = color
					s.X = msg.X - 2
					s.cursor = fgRgb(170, 170, 170, pointer)
				} else {
					s.inputLock = false
					if len(s.input) > 0 {
						s.color[s.inputColor] = setColor(s.input)
					}
					s.input = ""
				}
				if !okS && !okC {
					s.X = xMin + 1
					s.cursor = " "
				}
			} else if s.showFile && msg.X <= xMin {
				_, ok := s.fileList[msg.Y]
				if ok {
					s.X = 0
					s.cursor = fgRgb(170, 170, 170, pointer)
				} else {
					s.cursor = " "
				}
			} else if msg.X <= xMin {
				s.X = xMin + 1
				s.cursor = " "
			} else {
				s.cursor = s.cursorStore
			}

			if msg.X > xMin && msg.X < s.columns {
				s.X = msg.X
			}
			if msg.Y > 0 && msg.Y < s.rows {
				s.Y = msg.Y
			}

		case tea.MouseLeft:
			if s.showMenu && msg.X < menuWidth {
				symbol, ok := symbols[msg.Y][msg.X]
				if ok {
					s.cursorStore = symbol
					s.cursor = symbol
				}
			} else if s.showFile && msg.X < fileWidth {
				s.showFile = false
				file, ok := s.fileList[msg.Y]
				if ok {
					content, err := os.ReadFile(file)
					if err != nil {
						log.Fatal(err)
					}
					s.load(string(content))
				}
			} else {
				s.pixels = append(s.pixels, pixel{X: msg.X, Y: msg.Y, symbol: fgRgb(s.color["R"], s.color["G"], s.color["B"], s.cursor)})
			}

		case tea.MouseRight:
			s.pixels = append(s.pixels, pixel{X: msg.X, Y: msg.Y, symbol: " "})

		case tea.MouseMiddle:
			s.pixels = []pixel{}

		case tea.MouseWheelDown:
			color := colors[msg.Y][msg.X]
			s.color[color] = decrease(s.color[color])

		case tea.MouseWheelUp:
			color := colors[msg.Y][msg.X]
			s.color[color] = increase(s.color[color])
		}

	case tea.WindowSizeMsg:
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
		saveImage(screenString)
	}

	return screenString
}

type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Millisecond * 10)
	return tickMsg{}
}

func fgRgb(r int, g int, b int, symbol string) string {
	return "\033[38;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m" + symbol + "\033[0m"
}
