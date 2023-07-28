package screen

import (
	"github.com/14Artemiy88/termPaint/src/color"
	"github.com/14Artemiy88/termPaint/src/config"
	tea "github.com/charmbracelet/bubbletea"
)

func MouseBind(msg tea.MouseMsg, s *Screen) {
	switch msg.Type {
	case tea.MouseMotion:
		mouseMotion(msg, s)

	case tea.MouseLeft:
		mouseLeft(msg, s)

	case tea.MouseRight:
		Pixels.add(Pixel{X: msg.X, Y: msg.Y, Symbol: " "})

	case tea.MouseMiddle:
		Pixels = []Pixel{}

	case tea.MouseWheelDown:
		if c, ok := Colors[msg.Y]; ok && CC.Brush == Pointer {
			CC.Color[c] = color.Decrease(CC.Color[c])
		}
		if CC.Brush > Dot && CC.Symbol != config.Cfg.Pointer {
			if msg.Ctrl {
				if CC.Store.Brush == ESquare || CC.Store.Brush == FSquare {
					if CC.Height > 1 {
						CC.Height--
					}
				}
			} else {
				if CC.Width > 1 {
					CC.Width--
				}
			}
		}

	case tea.MouseWheelUp:
		if c, ok := Colors[msg.Y]; ok && CC.Brush == Pointer {
			CC.Color[c] = color.Increase(CC.Color[c])
		}
		if CC.Brush > Dot && CC.Symbol != config.Cfg.Pointer {
			if msg.Ctrl {
				if CC.Store.Brush == ESquare || CC.Store.Brush == FSquare {
					CC.Height++
				}
			} else {
				CC.Width++
			}
		}
	}
}
