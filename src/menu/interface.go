package menu

import "github.com/14Artemiy88/termPaint/src/config"

type Screen interface {
	GetPixels() [][]string
	GetDirectory() string
	SetDirectory(directory string)
	GetHeight() int
	GetWidth() int
	GetConfig() *config.Config
}
