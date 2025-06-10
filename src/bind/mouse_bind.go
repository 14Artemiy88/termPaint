package bind

import (
	"github.com/14Artemiy88/termPaint/src/cursor"
	"github.com/14Artemiy88/termPaint/src/menu"
	"github.com/14Artemiy88/termPaint/src/pixel"
	tea "github.com/charmbracelet/bubbletea"
)

func MouseBind(msg tea.MouseMsg, s Screen) {
	// Выносим общие проверки
	isPointer := cursor.CC.Brush == cursor.Pointer
	canChangeSize := cursor.CC.Brush > cursor.Dot && cursor.CC.Symbol != s.GetConfig().Pointer
	minSize := 1

	switch msg.Type {
	case tea.MouseMotion:
		mouseMotion(msg, s)

	case tea.MouseLeft:
		mouseLeft(msg.X, msg.Y, s)

	case tea.MouseRight:
		s.AddPixels(pixel.Pixel{Coord: pixel.Coord{X: msg.X, Y: msg.Y}, Symbol: " "})

	case tea.MouseMiddle:
		s.ClearUnsavedPixels()

	case tea.MouseWheelDown:
		// Обработка изменения цвета
		if isPointer {
			switch menu.GetColor(msg.Y) {
			case "r":
				cursor.CC.Color.R = pixel.Decrease(cursor.CC.Color.R)
			case "g":
				cursor.CC.Color.G = pixel.Decrease(cursor.CC.Color.G)
			case "b":
				cursor.CC.Color.B = pixel.Decrease(cursor.CC.Color.B)
			}
		}

		// Обработка изменения размера
		if !canChangeSize {
			break
		}

		if msg.Ctrl {
			if isSquareBrush(cursor.CC.Store.Brush) && cursor.CC.Height > minSize {
				cursor.CC.Height--
			}
		} else if cursor.CC.Width > minSize {
			cursor.CC.Width--
		}

	case tea.MouseWheelUp:
		// Обработка изменения цвета
		if isPointer {
			switch menu.GetColor(msg.Y) {
			case "r":
				cursor.CC.Color.R = pixel.Increase(cursor.CC.Color.R)
			case "g":
				cursor.CC.Color.G = pixel.Increase(cursor.CC.Color.G)
			case "b":
				cursor.CC.Color.B = pixel.Increase(cursor.CC.Color.B)
			}
		}

		// Обработка изменения размера
		if !canChangeSize {
			break
		}

		if msg.Ctrl {
			if isSquareBrush(cursor.CC.Store.Brush) {
				cursor.CC.Height++
			}
		} else {
			cursor.CC.Width++
		}

	case tea.MouseUnknown, tea.MouseRelease:
	}
}

// Вспомогательная функция для проверки квадратных кистей
func isSquareBrush(brush cursor.Type) bool {
	return brush == cursor.ESquare || brush == cursor.FSquare
}
