package src

import (
	"strings"
)

const liveTime = 200

type message struct {
	liveTime int
	text     string
}

func (s *Screen) SetMessage(text string) {
	s.Messages = append(s.Messages, message{text: text, liveTime: liveTime})
	textLen := len(text)
	if textLen > s.MessageWidth {
		s.MessageWidth = textLen
	}
}

func DrawMsg(messages []message, width int, screen [][]string) [][]string {
	ClearMessage(screen, width+5, len(messages)+2)
	for k, m := range messages {
		DrawString(1, 1+k, m.text, screen)
	}
	return screen
}

func ClearMessage(screen [][]string, width int, height int) [][]string {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			SetByKeys(x, y, " ", screen)
		}
		SetByKeys(width, y, "│", screen)
	}
	DrawString(0, height, strings.Repeat("─", width)+"┘", screen)

	return screen
}
