package menu

import (
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/message"
)

type Screen interface {
	GetPixels() [][]string
	GetDirectory() string
	SetDirectory(directory string)
	GetHeight() int
	GetWidth() int
	GetConfig() *config.Config
	GetMessage() message.Message
}
