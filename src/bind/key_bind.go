package bind

import (
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/menu"
	"github.com/14Artemiy88/termPaint/src/message"
	"github.com/14Artemiy88/termPaint/src/pixel"
	tea "github.com/charmbracelet/bubbletea"
	"os"
	"strconv"
)

func KeyBind(msg tea.KeyMsg, s Screen) (tea.Model, tea.Cmd) {
	switch msg.Type {
	case tea.KeyCtrlC, tea.KeyEsc:
		return s, tea.Quit

	// Help
	case tea.KeyCtrlH, tea.KeyF1:
		switchMenu(menu.Help)

	// menu
	case tea.KeyTab, tea.KeyF2:
		switchMenu(menu.SymbolColor)

	// file
	case tea.KeyCtrlO, tea.KeyF3:
		switchMenu(menu.File)

	// shape
	case tea.KeyF4:
		switchMenu(menu.Shape)

	// line
	//case tea.KeyF5:
	//	switchMenu(menu.Line)

	// shape
	case tea.KeyF6, tea.KeyCtrlK:
		switchMenu(menu.Config)

	case tea.KeyCtrlF:
		if cursor.CC.Brush == cursor.Fill {
			cursor.CC.Store.Brush = cursor.Dot
		} else {
			cursor.CC.Store.Brush = cursor.Fill
		}
		//cursor.CC.SetCursor(config.Cfg.FillCursor)

	// save
	case tea.KeyCtrlS:
		s.SetSave(true)
		s.SetShowInputSave(true)
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
		s.SetShowInputSave(false)

	case tea.KeySpace:
		cursor.CC.SetCursor(msg.String())

	// set cursor or color
	case tea.KeyRunes:
		if menu.Type == menu.SymbolColor && menu.Input.Lock {
			if _, err := strconv.Atoi(string(msg.Runes)); err == nil {
				menu.Input.Value += string(msg.Runes)
				color, err := strconv.Atoi(menu.Input.Value)
				if err != nil {
					s.GetMessage().SetMessage(err.Error())
				}
				switch menu.Input.Color {
				case "r":
					cursor.CC.Color.R = pixel.SetColor(color)
				case "g":
					cursor.CC.Color.G = pixel.SetColor(color)
				case "b":
					cursor.CC.Color.B = pixel.SetColor(color)
				}
			} else {
				s.GetMessage().SetMessage(err.Error())
			}
		} else if s.IsShowInputSave() {
			menu.Input.Value += string(msg.Runes)
		} else {
			cursor.CC.SetCursor(string(msg.Runes))
		}
	}

	return s, nil
}

func switchMenu(menuType menu.MenuType) {
	if menu.Type == menuType {
		menu.Type = menu.None
	} else {
		menu.Type = menuType
	}
}
