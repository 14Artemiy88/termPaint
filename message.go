package main

import "strings"

const messageHeight = 3
const liveTime = 200

func (s *screen) setMessage(text string) {
	s.message = text
	s.showMessageTime = liveTime
}

func drawMsg(text string, screen [][]string) [][]string {
	clearMessage(screen, len(text)+5)
	drawString(1, 1, text, screen)
	return screen
}

func clearMessage(screen [][]string, width int) [][]string {
	for i := 0; i < messageHeight; i++ {
		for j := 0; j < width; j++ {
			screen[i][j] = " "
		}
		screen[i][width] = "│"
	}
	drawString(0, messageHeight, strings.Repeat("─", width)+"┘", screen)

	return screen
}
