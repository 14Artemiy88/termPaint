package bind

import tea "github.com/charmbracelet/bubbletea"

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
}
