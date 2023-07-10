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
	X          int
	Y          int
	columns    int
	rows       int
	pixel      string
	pixels     []pixel
	color      map[string]int
	showMenu   bool
	showHelp   bool
	save       bool
	inputLock  bool
	input      string
	inputColor string
}

type pixel struct {
	X      int
	Y      int
	symbol string
}

func main() {
	p := tea.NewProgram(&screen{
		pixel: "#",
		color: map[string]int{"R": 255, "G": 255, "B": 255},
	}, tea.WithAltScreen(), tea.WithMouseAllMotion())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func (s screen) Init() tea.Cmd {
	return tick
}

func (s *screen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return s, tea.Quit

		case tea.KeyTab:
			s.showHelp = false
			s.showMenu = !s.showMenu

		case tea.KeyCtrlS:
			s.save = true
			s.showMenu = false
			s.showHelp = false

		case tea.KeyEnter:
			s.showMenu = false
			s.showHelp = !s.showHelp

		case tea.KeyRunes:
			if s.showMenu && s.inputLock {
				if _, err := strconv.Atoi(string(msg.Runes)); err == nil {
					s.input += string(msg.Runes)
				}
			} else {
				s.pixel = string(msg.Runes)
			}
		}

	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseMotion:
			color, ok := colors[msg.Y][msg.X]
			if ok {
				s.inputLock = true
				s.inputColor = color
			} else {
				s.inputLock = false
				if len(s.input) > 0 {
					s.color[s.inputColor] = setColor(s.input)
				}
				s.input = ""
			}
			xMin := 0
			if s.showMenu {
				xMin = menuWidth
			}
			if s.showHelp {
				xMin = helpWidth
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
					s.pixel = symbol
				}
			} else {
				s.pixels = append(s.pixels, pixel{X: msg.X, Y: msg.Y, symbol: fgRgb(s.color["R"], s.color["G"], s.color["B"], s.pixel)})
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

func setColor(color string) int {
	c, _ := strconv.Atoi(color)
	if c < 255 {
		return c
	}

	return 255
}

func decrease(color int) int {
	if color > 0 {
		color--
	}

	return color
}

func increase(color int) int {
	if color < 255 {
		color++
	}

	return color
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

	if !s.save {
		screen[s.Y][s.X] = fgRgb(s.color["R"], s.color["G"], s.color["B"], s.pixel)
	}

	if s.showMenu {
		drawMenu(s, screen)
	}
	if s.showHelp {
		drawHelpMenu(s, screen)
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

func saveImage(image string) {
	f, err := os.Create(time.Now().Format("termPaint_01-02-2006_15:04:05.txt"))
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)
	_, err = f.WriteString(image)
	if err != nil {
		log.Fatal(err)
	}
}

func fgRgb(r int, g int, b int, symbol string) string {
	return "\033[38;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m" + symbol + "\033[0m"
}
