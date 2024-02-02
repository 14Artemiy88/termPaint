package bind

import (
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/menu"
	"github.com/14Artemiy88/termPaint/src/pixel"
	tea "github.com/charmbracelet/bubbletea"
)

func MouseBind(msg tea.MouseMsg, s Screen) {
	switch msg.Action {
	case tea.MouseActionMotion:
		mouseMotion(msg, s)
	case tea.MouseActionPress:
	case tea.MouseActionRelease:
	}

	switch msg.Button {
	case tea.MouseButtonLeft:
		mouseLeft(msg.X, msg.Y, s)

	case tea.MouseButtonRight:
		s.AddPixels(pixel.Pixel{Coord: pixel.Coord{X: msg.X, Y: msg.Y}, Symbol: " "})

	case tea.MouseButtonMiddle:
		s.ClearUnsavedPixels()

	case tea.MouseButtonWheelDown:
		if c, ok := menu.Colors[msg.Y]; ok && cursor.CC.Brush == cursor.Pointer {
			switch c {
			case "r":
				cursor.CC.Color.R = pixel.Decrease(cursor.CC.Color.R)
			case "g":
				cursor.CC.Color.G = pixel.Decrease(cursor.CC.Color.G)
			case "b":
				cursor.CC.Color.B = pixel.Decrease(cursor.CC.Color.B)
			}
		}
		if cursor.CC.Brush > cursor.Dot && cursor.CC.Symbol != s.GetConfig().Pointer {
			if msg.Ctrl {
				if cursor.CC.Store.Brush == cursor.ESquare || cursor.CC.Store.Brush == cursor.FSquare {
					if cursor.CC.Height > 1 {
						cursor.CC.Height--
					}
				}
			} else {
				if cursor.CC.Width > 1 {
					cursor.CC.Width--
				}
			}
		}

	case tea.MouseButtonWheelUp:
		if c, ok := menu.Colors[msg.Y]; ok && cursor.CC.Brush == cursor.Pointer {
			switch c {
			case "r":
				cursor.CC.Color.R = pixel.Increase(cursor.CC.Color.R)
			case "g":
				cursor.CC.Color.G = pixel.Increase(cursor.CC.Color.G)
			case "b":
				cursor.CC.Color.B = pixel.Increase(cursor.CC.Color.B)
			}
		}
		if cursor.CC.Brush > cursor.Dot && cursor.CC.Symbol != s.GetConfig().Pointer {
			if msg.Ctrl {
				if cursor.CC.Store.Brush == cursor.ESquare || cursor.CC.Store.Brush == cursor.FSquare {
					cursor.CC.Height++
				}
			} else {
				cursor.CC.Width++
			}
		}
	case tea.MouseButtonNone:
	case tea.MouseButtonWheelLeft:
	case tea.MouseButtonWheelRight:
	case tea.MouseButtonBackward:
	case tea.MouseButtonForward:
	case tea.MouseButton10:
	case tea.MouseButton11:
	}
}
