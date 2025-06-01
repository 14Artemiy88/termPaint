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

type Color struct {
	R int
	G int
	B int
}

func GetConstColor(color string) Color {
	switch color {
	case "white":
		return Color{R: maxColor, G: maxColor, B: maxColor}
	case "green":
		return Color{R: color2, G: color186, B: color31}
	case "yellow":
		return Color{R: color190, G: color175, B: minColor}
	case "gray":
		return Color{R: color150, G: color150, B: color150}
	case "cian":
		return Color{R: minColor, G: color200, B: color200}
	case "red":
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
