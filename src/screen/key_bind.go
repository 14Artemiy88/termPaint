package screen

import (
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/menu"
	"github.com/14Artemiy88/termPaint/src/message"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"strconv"
)

func KeyBind(msg tea.KeyMsg, s *Screen) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyCtrlC, tea.KeyEsc:
		return s, tea.Quit

	// Help
	case tea.KeyCtrlH, tea.KeyF1:
		if menu.Type == menu.Help {
			menu.Type = menu.None
		} else {
			menu.Type = menu.Help
		}

	// menu
	case tea.KeyTab, tea.KeyF2:
		if menu.Type == menu.SymbolColor {
			menu.Type = menu.None
		} else {
			menu.Type = menu.SymbolColor
		}

	// file
	case tea.KeyCtrlO, tea.KeyF3:
		if menu.Type == menu.File {
			menu.Type = menu.None
		} else {
			menu.Type = menu.File
		}

	// shape
	case tea.KeyF4:
		if menu.Type == menu.Shape {
			menu.Type = menu.None
		} else {
			menu.Type = menu.Shape
		}

	// line
	//case tea.KeyF5:
	//	if menu.Type == menu.Line {
	//		menu.Type = menu.None
	//	} else {
	//		menu.Type = menu.Line
	//	}

	case tea.KeyCtrlF:
		if cursor.CC.Brush == cursor.Fill {
			cursor.CC.Store.Brush = cursor.Dot
		} else {
			cursor.CC.Store.Brush = cursor.Fill
		}
		//cursor.CC.SetCursor(config.Cfg.FillCursor)

	// save
	case tea.KeyCtrlS:
		s.Save = true
		s.ShowInputSave = true
		menu.Type = menu.None
		message.Msg = []message.Message{}

	// del file
	case tea.KeyDelete:
		if len(menu.FilePath) > 0 {
			_ = os.Remove(menu.FilePath)
		}

	case tea.KeyBackspace:
		menu.Input.Value = menu.Input.Value[:len(menu.Input.Value)-1]

	case tea.KeyEnter:
		s.ShowInputSave = false

	case tea.KeySpace:
		cursor.CC.SetCursor(msg.String())

	// set cursor or color
	case tea.KeyRunes:
		if menu.Type == menu.SymbolColor && menu.Input.Lock {
			if _, err := strconv.Atoi(string(msg.Runes)); err == nil {
				menu.Input.Value += string(msg.Runes)
				cursor.CC.Color[menu.Input.Color] = color.SetColor(menu.Input.Value)
			} else {
				message.SetMessage(err.Error())
			}
		} else if s.ShowInputSave {
			menu.Input.Value += string(msg.Runes)
		} else {
			cursor.CC.SetCursor(string(msg.Runes))
		}
	}

	return s, nil
}
