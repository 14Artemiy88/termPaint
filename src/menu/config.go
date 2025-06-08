package menu

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/utils"
)

const ConfigWidth = 65
const firstLvlX = 3
const secondLvlX = 10
const valueX = 31

var availableFields = []string{
	"Background",
	"BackgroundColor",
	"DefaultCursor",
	"DefaultColor",
	"Pointer",
	"PointerColor",
	"FillCursor",
	"ShowFolder",
	"ShowHiddenFolder",
	"ImageSaveDirectory",
	"ImageSaveNameFormat",
	"NotificationTime",
	"Notifications",
}

func drawConfigMenu(s Screen) {
	red := pixel.GetConstColor(pixel.Red)
	green := pixel.GetConstColor(pixel.Green)
	white := pixel.GetConstColor(pixel.White)
	pixels := s.GetPixels()
	v := reflect.ValueOf(*s.GetConfig())
	typeOfConfig := v.Type()

	ClearMenu(s, pixels, ConfigWidth)

	// Заголовок
	drawTitle(ConfigWidth, "Config", 1, "┐", pixels)

	h := 3
	height := s.GetHeight()

	for index := 0; index < v.NumField(); index++ {
		if utils.InArray(typeOfConfig.Field(index).Name, availableFields) {
			field := v.Field(index).Interface()
			utils.DrawString(firstLvlX, h, typeOfConfig.Field(index).Name, white, pixels)

			clr := white

			switch field {
			case true:
				clr = green
			case false:
				clr = red
			}

			switch reflect.TypeOf(field).String() {
			case "string", "int", "bool":
				utils.DrawString(valueX, h, fmt.Sprintf("%v", field), clr, pixels)

				h++
			case "map[string]int":
				h++

				for _, c := range []string{"r", "g", "b"} {
					for _, k := range v.Field(index).MapKeys() {
						if c == k.String() {
							utils.DrawString(secondLvlX, h, strings.ToUpper(c), clr, pixels)
							utils.DrawString(valueX, h, fmt.Sprintf("%v", v.Field(index).MapIndex(k).Interface()), clr, pixels)

							h++

							break
						}
					}
				}
			default:
				v = reflect.ValueOf(field)
				typeOfConfig = v.Type()
				h++

				for i := 0; i < v.NumField(); i++ {
					field = v.Field(i).Interface()
					switch field {
					case true:
						clr = green
					case false:
						clr = red
					}

					utils.DrawString(secondLvlX, h, typeOfConfig.Field(i).Name, white, pixels)
					utils.DrawString(valueX, h, fmt.Sprintf("%v", v.Field(i).Interface()), clr, pixels)

					h++
					if h >= height-6 {
						break
					}
				}
			}

			if h >= height-6 {
				break
			}
		}
	}

	// Заголовок 2
	drawTitle(ConfigWidth, "Note", height-4, "┤", pixels)

	utils.DrawString(firstLvlX, height-2, "All configuration parameters are", white, pixels)
	utils.DrawString(firstLvlX, height-1, "stored in", white, pixels)
	utils.DrawString(len("stored in")+4, height-1, "~/.config/termPaint", green, pixels)
}
