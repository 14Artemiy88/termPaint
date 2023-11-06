package bind

import (
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/message"
	"github.com/14Artemiy88/termPaint/src/pixel"
	tea "github.com/charmbracelet/bubbletea"
)

type Screen interface {
	SetSave(bool)
	SetShowInputSave(bool)
	IsShowInputSave() bool
	Init() tea.Cmd
	Update(msg tea.Msg) (tea.Model, tea.Cmd)
	View() string
	LoadImage(screenString string)
	LoadFromImage(path string)
	GetPixel(y int, x int) string
	GetPixels() [][]string
	GetDirectory() string
	SetDirectory(string)
	GetWidth() int
	GetHeight() int
	ClearUnsavedPixels()
	AddPixels(pixels ...pixel.Pixel)
	GetConfig() *config.Config
	GetMessage() message.Message
}
