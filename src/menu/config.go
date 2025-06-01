package menu

import (
	"fmt"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/utils"
	"reflect"
	"strings"
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
	red := pixel.GetConstColor("red")
	gray := pixel.GetConstColor("gray")
	green := pixel.GetConstColor("green")
	white := pixel.GetConstColor("white")
	yellow := pixel.GetConstColor("yellow")

	screen := s.GetPixels()
	ClearMenu(s, screen, ConfigWidth)
	v := reflect.ValueOf(*s.GetConfig())
	typeOfConfig := v.Type()
	title := "Config"
	lenTitle := len(title)
	str := strings.Repeat("─", ConfigWidth-lenTitle-2) + "┐"
	utils.DrawString(1, 1, title, yellow, screen)
	utils.DrawString(lenTitle+2, 1, str, gray, screen)

	h := 3
	height := s.GetHeight()
	for i := 0; i < v.NumField(); i++ {
		if utils.InArray(typeOfConfig.Field(i).Name, availableFields) {
			field := v.Field(i).Interface()
			utils.DrawString(firstLvlX, h, typeOfConfig.Field(i).Name, white, screen)
			clr := white
			switch field {
			case true:
				clr = green
			case false:
				clr = red
			}
			switch reflect.TypeOf(field).String() {
			case "string", "int", "bool":
				utils.DrawString(valueX, h, fmt.Sprintf("%v", field), clr, screen)
				h++
			case "map[string]int":
				h++
				for _, c := range []string{"r", "g", "b"} {
					for _, k := range v.Field(i).MapKeys() {
						if c == k.String() {
							utils.DrawString(secondLvlX, h, strings.ToUpper(c), clr, screen)
							utils.DrawString(valueX, h, fmt.Sprintf("%v", v.Field(i).MapIndex(k).Interface()), clr, screen)
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
					utils.DrawString(secondLvlX, h, typeOfConfig.Field(i).Name, white, screen)
					utils.DrawString(valueX, h, fmt.Sprintf("%v", v.Field(i).Interface()), clr, screen)
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

	title = "Note"
	lenTitle = len(title)
	str = strings.Repeat("─", ConfigWidth-lenTitle-2) + "┤"
	utils.DrawString(1, height-4, title, yellow, screen)
	utils.DrawString(lenTitle+2, height-4, str, gray, screen)
	utils.DrawString(firstLvlX, height-2, "All configuration parameters are", white, screen)
	utils.DrawString(firstLvlX, height-1, "stored in", white, screen)
	utils.DrawString(len("stored in")+4, height-1, "~/.config/termPaint", green, screen)
}
