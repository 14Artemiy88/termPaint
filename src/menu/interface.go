package menu

type Screen interface {
	GetPixels() [][]string
	GetDirectory() string
	SetDirectory(directory string)
}
