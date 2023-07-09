package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
	"strconv"
	"strings"
	"time"
)

type screen struct {
	X        int
	Y        int
	columns  int
	rows     int
	pixel    string
	pixels   []pixel
	color    clr
	showMenu bool
}

type clr struct {
	R int
	G int
	B int
}
type pixel struct {
	X      int
	Y      int
	symbol string
}

const menuWidth = 12

var symbols = map[string]pixel{
	"2,1": {X: 2, Y: 1, symbol: "."},
	"4,1": {X: 4, Y: 1, symbol: ":"},
	"6,1": {X: 6, Y: 1, symbol: "!"},
	"8,1": {X: 8, Y: 1, symbol: "/"},
	"2,3": {X: 2, Y: 3, symbol: "r"},
	"4,3": {X: 4, Y: 3, symbol: "("},
	"6,3": {X: 6, Y: 3, symbol: "l"},
	"8,3": {X: 8, Y: 3, symbol: "1"},
	"2,5": {X: 2, Y: 5, symbol: "Z"},
	"4,5": {X: 4, Y: 5, symbol: "4"},
	"6,5": {X: 6, Y: 5, symbol: "H"},
	"8,5": {X: 8, Y: 5, symbol: "9"},
	"2,7": {X: 2, Y: 7, symbol: "W"},
	"4,7": {X: 4, Y: 7, symbol: "8"},
	"6,7": {X: 6, Y: 7, symbol: "$"},
	"8,7": {X: 8, Y: 7, symbol: "@"},
	"2,9": {X: 2, Y: 9, symbol: "░"},
	"4,9": {X: 4, Y: 9, symbol: "▒"},
	"6,9": {X: 6, Y: 9, symbol: "▓"},
	"8,9": {X: 8, Y: 9, symbol: "█"},
}

type menuClr struct {
	X      int
	Y      int
	symbol string
}

var colors = map[string]menuClr{
	"2,16": {X: 2, Y: 16, symbol: "R"},
	"2,18": {X: 2, Y: 18, symbol: "G"},
	"2,20": {X: 2, Y: 20, symbol: "B"},
}

func main() {
	p := tea.NewProgram(screen{
		pixel: "#",
		color: clr{R: 255, G: 255, B: 255},
	}, tea.WithAltScreen(), tea.WithMouseAllMotion())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

func (s screen) Init() tea.Cmd {
	return tick
}

func (s screen) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC, tea.KeyEsc:
			return s, tea.Quit

		case tea.KeyTab:
			s.showMenu = !s.showMenu

		case tea.KeyRunes:
			switch string(msg.Runes) {
			case "c":
				s.pixels = []pixel{}
			}
		}

	case tea.MouseMsg:
		switch msg.Type {
		case tea.MouseMotion:
			if msg.X > 0 && msg.X < s.columns {
				s.X = msg.X
			}
			if msg.Y > 0 && msg.Y < s.rows {
				s.Y = msg.Y
			}

		case tea.MouseLeft:
			if s.showMenu && msg.X < menuWidth {
				s.pixel = symbols[strconv.Itoa(msg.X)+","+strconv.Itoa(msg.Y)].symbol
			} else {
				s.pixels = append(s.pixels, pixel{X: msg.X, Y: msg.Y, symbol: fgRgb(s.color.R, s.color.G, s.color.B, s.pixel)})
			}

		case tea.MouseRight:
			s.pixels = append(s.pixels, pixel{X: msg.X, Y: msg.Y, symbol: " "})

		case tea.MouseWheelDown:
			color := colors[strconv.Itoa(msg.X)+","+strconv.Itoa(msg.Y)]
			switch color.symbol {
			case "R":
				s.color.R = decrease(s.color.R)
			case "G":
				s.color.G = decrease(s.color.G)
			case "B":
				s.color.B = decrease(s.color.B)
			}

		case tea.MouseWheelUp:
			color := colors[strconv.Itoa(msg.X)+","+strconv.Itoa(msg.Y)]
			switch color.symbol {
			case "R":
				s.color.R = increase(s.color.R)
			case "G":
				s.color.G = increase(s.color.G)
			case "B":
				s.color.B = increase(s.color.B)
			}
		}

	case tea.WindowSizeMsg:
		s.columns = msg.Width
		s.rows = msg.Height
	}

	return s, nil
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
func (s screen) View() string {
	if s.rows == 0 {
		return ""
	}

	board := make([][]string, s.rows)

	// draw empty board
	for i := 0; i < s.rows; i++ {
		line := strings.Split(strings.Repeat(" ", s.columns), "")
		board[i] = line
	}

	for _, p := range s.pixels {
		board[p.Y][p.X] = p.symbol
	}

	board[s.Y][s.X] = fgRgb(s.color.R, s.color.G, s.color.B, s.pixel)

	if s.showMenu {
		drawMenu(s, board)
	}

	var boardString string
	for i, line := range board {
		boardString += strings.Join(line, "")
		if i < len(board)-1 {
			boardString += "\n"
		}
	}

	return boardString
}

type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Millisecond * 10)
	return tickMsg{}
}

func drawMenu(s screen, board [][]string) [][]string {
	for i := 0; i < s.rows; i++ {
		for j := 0; j < menuWidth; j++ {
			board[i][j] = " "
		}
		board[i][menuWidth] = "│"
	}
	drawSymbolMenu(s, board)
	drawColorMenu(s, board)

	return board
}
func drawSymbolMenu(s screen, board [][]string) [][]string {
	for _, symbol := range symbols {
		board[symbol.Y][symbol.X] = fgRgb(s.color.R, s.color.G, s.color.B, symbol.symbol)
	}

	return board
}

func drawColorMenu(s screen, board [][]string) [][]string {
	for _, c := range colors {
		switch c.symbol {
		case "R":
			drawColorValue(c.X, c.Y, s.color.R, board)
			board[c.Y][c.X] = fgRgb(s.color.R, 0, 0, "█")
		case "G":
			drawColorValue(c.X, c.Y, s.color.G, board)
			board[c.Y][c.X] = fgRgb(0, s.color.G, 0, "█")
		case "B":
			drawColorValue(c.X, c.Y, s.color.B, board)
			board[c.Y][c.X] = fgRgb(0, 0, s.color.B, "█")
		}

	}
	board[22][2] = fgRgb(s.color.R, s.color.G, s.color.B, "█")
	board[22][3] = fgRgb(s.color.R, s.color.G, s.color.B, "█")
	board[22][4] = fgRgb(s.color.R, s.color.G, s.color.B, "█")
	board[23][2] = fgRgb(s.color.R, s.color.G, s.color.B, "█")
	board[23][3] = fgRgb(s.color.R, s.color.G, s.color.B, "█")
	board[23][4] = fgRgb(s.color.R, s.color.G, s.color.B, "█")

	return board
}

func drawColorValue(X int, Y int, color int, board [][]string) [][]string {
	r := strings.Split(strconv.Itoa(color), "")
	for k, i := range r {
		board[Y][X+2+k] = i
	}

	return board
}
func fgRgb(r int, g int, b int, symbol string) string {
	return "\033[38;2;" + strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m" + symbol + "\033[0m"
}
