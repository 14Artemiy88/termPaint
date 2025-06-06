package screen

import (
	"fmt"
	"strings"
	"time"

	"github.com/14Artemiy88/termPaint/src/bind"
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/menu"
	"github.com/14Artemiy88/termPaint/src/message"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/utils"
	tea "github.com/charmbracelet/bubbletea"
)

type Screen struct {
	Width         int
	Height        int
	ShowInputSave bool
	Save          bool
	SavedPixels   [][]string
	UnsavedPixels map[string]pixel.Pixel

	Message
	Config
}

type Config interface {
	GetNotificationTime() int
	GetImageSaveDirectory() string
	SetImageSaveDirectory(directory string)
	WithBackground() bool
	GetBackgroundColor() map[string]int
}

type Message interface {
	SetMessage(string)
}

func (s *Screen) Init() tea.Cmd {
	return tick
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

		menu.Blink()

		return s, tick

	case tea.KeyMsg:
		return bind.KeyBind(msg, s)

	case tea.MouseMsg:
		bind.MouseBind(msg, s)

	case tea.WindowSizeMsg:
		cursor.CC.X = msg.Width / 2
		cursor.CC.Y = msg.Height / 2
		s.Width = msg.Width
		s.Height = msg.Height
	}

	return s, nil
}

func (s *Screen) View() string {
	if s.Height == 0 {
		return ""
	}

	s.drawClearScreen()
	s.drawScreen()
	menu.DrawMenu(s)
	s.showMsg()
	s.showCursor()
	s.showSaveInput()
	screenString := s.setScreenString()
	s.save(screenString)

	return s.screen(screenString)
}

func (s *Screen) SetSave(save bool) {
	s.Save = save
}

func (s *Screen) GetPixels() [][]string {
	return s.SavedPixels
}

func (s *Screen) GetPixel(y int, x int) string {
	if !utils.Isset(s.SavedPixels, y, x) {
		return ""
	}

	return s.SavedPixels[y][x]
}

func (s *Screen) SetShowInputSave(showInputSave bool) {
	s.ShowInputSave = showInputSave
}

func (s *Screen) IsShowInputSave() bool {
	return s.ShowInputSave
}

func (s *Screen) GetDirectory() string {
	return s.Config.GetImageSaveDirectory()
}

func (s *Screen) SetDirectory(directory string) {
	s.Config.SetImageSaveDirectory(directory)
}

func (s *Screen) GetWidth() int {
	return s.Width
}

func (s *Screen) GetHeight() int {
	return s.Height
}

func (s *Screen) AddPixels(pixels ...pixel.Pixel) {
	for _, p := range pixels {
		key := fmt.Sprintf("%d-%d", p.Coord.Y, p.Coord.X)
		s.UnsavedPixels[key] = p
	}
}

func (s *Screen) ClearUnsavedPixels() {
	s.UnsavedPixels = map[string]pixel.Pixel{}
}

func (s *Screen) SetConfig(c config.Config) {
	s.Config = &c
}

func (s *Screen) GetConfig() *config.Config {
	return s.Config.(*config.Config)
}

func (s *Screen) GetMessage() message.Message {
	return s.Message.(message.Message)
}

func (s *Screen) showCursor() {
	if !s.Save {
		cursor.CC.DrawCursor(s)
	}
}

func (s *Screen) save(screenString string) {
	if s.Save && !s.ShowInputSave {
		s.Save = false
		menu.SaveImage(s.Message, s.Config.GetImageSaveDirectory(), screenString)
	}
}

func (s *Screen) showSaveInput() {
	if s.ShowInputSave {
		menu.DrawSaveInput(s.SavedPixels)
	}
}

func (s *Screen) showMsg() {
	if len(message.Msg) > 0 {
		message.DrawMsg(message.Msg, message.MsgWidth, s.SavedPixels)
	}
}

func (s *Screen) setScreenString() string {
	var screenString string
	for i, line := range s.SavedPixels {
		screenString += strings.Join(line, "")
		if i < len(s.SavedPixels)-1 {
			screenString += "\n"
		}
	}

	return screenString
}

func (s *Screen) screen(screenString string) string {
	if s.Config.WithBackground() {
		color := s.Config.GetBackgroundColor()
		screenString = fmt.Sprintf(
			"\033[48;2;%d;%d;%dm%s",
			color["r"],
			color["g"],
			color["b"],
			screenString,
		)
	}

	return screenString
}

func (s *Screen) drawScreen() {
	for _, p := range s.UnsavedPixels {
		utils.SetByKeys(p.Coord.X, p.Coord.Y, p.Symbol, p.Color, s.SavedPixels)
	}
}

func (s *Screen) drawClearScreen() {
	s.SavedPixels = make([][]string, s.Height)
	for i := 0; i < s.Height; i++ {
		if len(s.SavedPixels) > i {
			s.SavedPixels[i] = strings.Split(strings.Repeat(" ", s.Width), "")
		}
	}
}

type tickMsg time.Time

func tick() tea.Msg {
	time.Sleep(time.Millisecond * 10)

	return tickMsg{}
}
