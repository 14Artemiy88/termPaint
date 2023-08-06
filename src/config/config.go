package config

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var Cfg Config

type Config struct {
	DefaultCursor       string                 `mapstructure:"default_cursor"`
	DefaultColor        map[string]int         `mapstructure:"default_color"`
	Pointer             string                 `mapstructure:"pointer"`
	PointerColor        map[string]int         `mapstructure:"pointer_color"`
	FillCursor          string                 `mapstructure:"fill_cursor"`
	Symbols             map[int]map[int]string `mapstructure:"symbols"`
	ShowFolder          bool                   `mapstructure:"show_folder"`
	ShowHiddenFolder    bool                   `mapstructure:"show_hidden_folder"`
	ImageSaveDirectory  string                 `mapstructure:"image_save_directory"`
	ImageSaveNameFormat string                 `mapstructure:"image_save_name_format"`
	NotificationTime    int                    `mapstructure:"notification_time"`
	Notifications       struct {
		SetSymbol           bool `mapstructure:"set_symbol"`
		Error               bool `mapstructure:"error"`
		SaveImage           bool `mapstructure:"save_image"`
		LoadImageSizeErrors bool `mapstructure:"load_image_size_errors"`
	} `mapstructure:"notifications"`
}

func InitConfig() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("Cannot determine the user's home dir:", err)
	}
	if os.Getenv("ENV") == "dev" {
		viper.SetConfigFile("config.yaml")
	} else {
		viper.SetConfigFile(homeDir + "/.config/termPaint/config.yaml")
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading Cfg file, %s", err)
	}
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
