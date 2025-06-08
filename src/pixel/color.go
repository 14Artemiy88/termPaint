package pixel

const (
	minColor = 0
	maxColor = 255

	color2   = 2
	color31  = 31
	color150 = 150
	color175 = 175
	color186 = 186
	color190 = 190
	color200 = 200
	color250 = 250
)

const (
	White  = "white"
	Yellow = "yellow"
	Gray   = "gray"
	Green  = "green"
	Cyan   = "cyan"
	Red    = "red"
)

type Color struct {
	R int
	G int
	B int
}

func GetConstColor(color string) Color {
	switch color {
	case White:
		return Color{R: maxColor, G: maxColor, B: maxColor}
	case Green:
		return Color{R: color2, G: color186, B: color31}
	case Yellow:
		return Color{R: color190, G: color175, B: minColor}
	case Gray:
		return Color{R: color150, G: color150, B: color150}
	case Cyan:
		return Color{R: minColor, G: color200, B: color200}
	case Red:
		return Color{R: color250, G: minColor, B: minColor}
	default:
		return Color{R: maxColor, G: maxColor, B: maxColor}
	}
}

func SetColor(color int) int {
	if color < maxColor {
		return color
	}

	return maxColor
}

func Decrease(color int) int {
	if color > minColor {
		color--
	}

	return color
}

func Increase(color int) int {
	if color < maxColor {
		color++
	}

	return color
}

func MinMaxColor(color int) int {
	if color > minColor {
		return minColor
	}

	return maxColor
}
