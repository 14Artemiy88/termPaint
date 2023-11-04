package screen

import (
	"fmt"
	"github.com/14Artemiy88/termPaint/src/bind"
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/menu"
	"github.com/14Artemiy88/termPaint/src/message"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/utils"
	tea "github.com/charmbracelet/bubbletea"
	"strings"
	"time"
)

type Screen struct {
	size          size
	ShowInputSave bool
	Save          bool
	Directory     string
}

func (s *Screen) Init() tea.Cmd {
	return tick
}

var blink = map[bool]string{
	true:  "|",
	false: " ",
}

var Pixels [][]string

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
		return bind.KeyBind(msg, s)

	case tea.MouseMsg:
		bind.MouseBind(msg, s)

	case tea.WindowSizeMsg:
		cursor.CC.X = msg.Width / 2
		cursor.CC.Y = msg.Height / 2
		Size.Width = msg.Width
		Size.Height = msg.Height
	}

	return s, nil
}

func (s *Screen) View() string {
	if Size.Height == 0 {
		return ""
	}

	drawClearScreen()
	drawScreen()
	menu.DrawMenu(s)
	showMsg()
	s.showCursor()
	s.showSaveInput()
	screenString := setScreenString()
	s.save(screenString)

	return screen(screenString)
}

func (s *Screen) SetSave(save bool) {
	s.Save = save
}

func (s *Screen) GetPixels() [][]string {
	return Pixels
}

func (s *Screen) GetPixel(y int, x int) string {
	return Pixels[y][x]
}

func (s *Screen) SetShowInputSave(showInputSave bool) {
	s.ShowInputSave = showInputSave
}

func (s *Screen) IsShowInputSave() bool {
	return s.ShowInputSave
}

func (s *Screen) GetDirectory() string {
	return s.Directory
}

func (s *Screen) SetDirectory(directory string) {
	s.Directory = directory
}

func (s *Screen) GetWidth() int {
	return Size.Width
}

func (s *Screen) GetHeight() int {
	return Size.Height
}

func (s *Screen) showCursor() {
	if !s.Save {
		cursor.CC.DrawCursor(s)
	}
}

func (s *Screen) save(screenString string) {
	if s.Save && !s.ShowInputSave {
		s.Save = false
		menu.SaveImage(screenString)
	}
}

func (s *Screen) showSaveInput() {
	if s.ShowInputSave {
		menu.DrawSaveInput(Pixels)
	}
}

func showMsg() {
	if len(message.Msg) > 0 {
		message.DrawMsg(message.Msg, message.MsgWidth, Pixels)
	}
}

func setScreenString() string {
	var screenString string
	for i, line := range Pixels {
		screenString += strings.Join(line, "")
		if i < len(Pixels)-1 {
			screenString += "\n"
		}
	}

	return screenString
}

func screen(screenString string) string {
	if config.Cfg.Background {
		screenString = fmt.Sprintf(
			"\033[48;2;%d;%d;%dm%s",
			config.Cfg.BackgroundColor["r"],
			config.Cfg.BackgroundColor["g"],
			config.Cfg.BackgroundColor["b"],
			screenString,
		)
	}

	return screenString
}

func drawScreen() {
	for _, p := range pixel.Pixels {
		utils.SetByKeys(p.Coord.X, p.Coord.Y, p.Symbol, p.Color, Pixels)
	}
}

func drawClearScreen() {
	Pixels = make([][]string, Size.Height)
	for i := 0; i < Size.Height; i++ {
		if len(Pixels) > i {
			Pixels[i] = strings.Split(strings.Repeat(" ", Size.Width), "")
		}
	}
}

type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Millisecond * 10)
	return tickMsg{}
}
