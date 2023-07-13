package main

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

var cfg Config

type Config struct {
	DefaultCursor       string                 `mapstructure:"default_cursor"`
	DefaultColor        map[string]int         `mapstructure:"default_color"`
	Pointer             string                 `mapstructure:"pointer"`
	PointerColor        map[string]int         `mapstructure:"pointer_color"`
	Symbols             map[int]map[int]string `mapstructure:"symbols"`
	ShowFolder          bool                   `mapstructure:"show_folder"`
	ShowHiddenFolder    bool                   `mapstructure:"show_hidden_folder"`
	ImageSaveDirectory  string                 `mapstructure:"image_save_directory"`
	ImageSaveNameFormat string                 `mapstructure:"image_save_name_format"`
}

func initConfig() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Println("Cannot determine the user's home dir:", err)
	}
	viper.SetConfigFile(homeDir + "/.config/termPaint/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading cfg file, %s", err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}
