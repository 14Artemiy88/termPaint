package main

import "strings"

const liveTime = 3000

type message struct {
	liveTime int
	text     string
}

func (s *screen) setMessage(text string) {
	s.messages = append(s.messages, message{text: text, liveTime: liveTime})
	textLen := len(text)
	if textLen > s.messageWidth {
		s.messageWidth = textLen
	}
}

func drawMsg(messages []message, width int, screen [][]string) [][]string {
	clearMessage(screen, width+5, len(messages)+2)
	for k, m := range messages {
		drawString(1, 1+k, m.text, screen)
	}
	return screen
}

func clearMessage(screen [][]string, width int, height int) [][]string {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			screen[i][j] = " "
		}
		screen[i][width] = "│"
	}
	drawString(0, height, strings.Repeat("─", width)+"┘", screen)

	return screen
}
