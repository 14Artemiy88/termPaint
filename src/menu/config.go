package menu

import (
	"fmt"
	"github.com/14Artemiy88/termPaint/src/config"
	"github.com/14Artemiy88/termPaint/src/pixel"
	"github.com/14Artemiy88/termPaint/src/size"
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

func drawConfigMenu(screen [][]string) [][]string {
	ClearMenu(screen, ConfigWidth)
	v := reflect.ValueOf(config.Cfg)
	typeOfConfig := v.Type()
	title := "Config"
	lenTitle := len(title)
	str := strings.Repeat("─", ConfigWidth-lenTitle-2) + "┐"
	utils.DrawString(1, 1, title, pixel.Yellow, screen)
	utils.DrawString(lenTitle+2, 1, str, pixel.Gray, screen)

	h := 3
	for i := 0; i < v.NumField(); i++ {
		if utils.InArray(typeOfConfig.Field(i).Name, availableFields) {
			field := v.Field(i).Interface()
			utils.DrawString(firstLvlX, h, typeOfConfig.Field(i).Name, pixel.White, screen)
			clr := pixel.White
			switch field {
			case true:
				clr = pixel.Green
			case false:
				clr = pixel.Red
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
						clr = pixel.Green
					case false:
						clr = pixel.Red
					}
					utils.DrawString(secondLvlX, h, typeOfConfig.Field(i).Name, pixel.White, screen)
					utils.DrawString(valueX, h, fmt.Sprintf("%v", v.Field(i).Interface()), clr, screen)
					h++
					if h >= size.Size.Height-6 {
						break
					}
				}
			}
			if h >= size.Size.Height-6 {
				break
			}
		}
	}

	title = "Note"
	lenTitle = len(title)
	str = strings.Repeat("─", ConfigWidth-lenTitle-2) + "┤"
	utils.DrawString(1, size.Size.Height-4, title, pixel.Yellow, screen)
	utils.DrawString(lenTitle+2, size.Size.Height-4, str, pixel.Gray, screen)
	utils.DrawString(firstLvlX, size.Size.Height-2, "All configuration parameters are", pixel.White, screen)
	utils.DrawString(firstLvlX, size.Size.Height-1, "stored in", pixel.White, screen)
	utils.DrawString(len("stored in")+4, size.Size.Height-1, "~/.config/termPaint", pixel.Green, screen)

	return screen
}
