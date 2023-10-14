package screen

import (
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/coord"
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/menu"
	"github.com/14Artemiy88/termPaint/src/pixel"
	tea "github.com/charmbracelet/bubbletea"
)

func MouseBind(msg tea.MouseMsg, s *Screen) {
	switch msg.Type {
	case tea.MouseMotion:
		mouseMotion(msg)

	case tea.MouseLeft:
		mouseLeft(msg.X, msg.Y, s)

	case tea.MouseRight:
		pixel.AddPixels(pixel.Pixel{Coord: coord.Coord{X: msg.X, Y: msg.Y}, Symbol: " "})

	case tea.MouseMiddle:
		pixel.Pixels = map[string]pixel.Pixel{}

	case tea.MouseWheelDown:
		if c, ok := menu.Colors[msg.Y]; ok && cursor.CC.Brush == cursor.Pointer {
			switch c {
			case "r":
				cursor.CC.Color.R = color.Decrease(cursor.CC.Color.R)
			case "g":
				cursor.CC.Color.G = color.Decrease(cursor.CC.Color.G)
			case "b":
				cursor.CC.Color.B = color.Decrease(cursor.CC.Color.B)
			}
		}
		if cursor.CC.Brush > cursor.Dot && cursor.CC.Symbol != config.Cfg.Pointer {
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

	case tea.MouseWheelUp:
		if c, ok := menu.Colors[msg.Y]; ok && cursor.CC.Brush == cursor.Pointer {
			switch c {
			case "r":
				cursor.CC.Color.R = color.Increase(cursor.CC.Color.R)
			case "g":
				cursor.CC.Color.G = color.Increase(cursor.CC.Color.G)
			case "b":
				cursor.CC.Color.B = color.Increase(cursor.CC.Color.B)
			}
		}
		if cursor.CC.Brush > cursor.Dot && cursor.CC.Symbol != config.Cfg.Pointer {
			if msg.Ctrl {
				if cursor.CC.Store.Brush == cursor.ESquare || cursor.CC.Store.Brush == cursor.FSquare {
					cursor.CC.Height++
				}
			} else {
				cursor.CC.Width++
			}
		}
	case tea.MouseUnknown:
	case tea.MouseRelease:
	}
}
