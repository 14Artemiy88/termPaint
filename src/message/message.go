package message

import (
	"strings"

	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/utils"
)

type Message struct {
	LiveTime int
	text     string
}

var Msg []Message
var MsgWidth int

type Config interface {
	GetNotificationTime() int
}

func (m Message) SetMessage(text string) {
	Msg = append(Msg, Message{text: text, LiveTime: m.LiveTime})

	textLen := len(text)
	if textLen > MsgWidth {
		MsgWidth = textLen
	}
}

func DrawMsg(messages []Message, width int, screen [][]string) [][]string {
	clearMessage(screen, width+5, len(messages)+2)

	white := pixel.GetConstColor("white")
	for k, m := range messages {
		utils.DrawString(1, 1+k, m.text, white, screen)
	}

	return screen
}

func clearMessage(screen [][]string, width int, height int) [][]string {
	white := pixel.GetConstColor("white")

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			utils.SetByKeys(x, y, " ", white, screen)
		}

		utils.SetByKeys(width, y, "│", white, screen)
	}

	utils.DrawString(0, height, strings.Repeat("─", width)+"┘", white, screen)

	return screen
}
