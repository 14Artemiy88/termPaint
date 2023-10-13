package config

import (
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
)

const configFileName = "config.yaml"
const configPath = "/.config/termPaint/"
const githubConfigFile = "https://raw.githubusercontent.com/14Artemiy88/termPaint/main/config.yaml"

var Cfg Config

type Config struct {
	Background          bool                   `mapstructure:"background"`
	BackgroundColor     map[string]int         `mapstructure:"background_color"`
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
		viper.SetConfigFile(configFileName)
	} else {
		viper.SetConfigFile(homeDir + configPath + configFileName)
	}
	if err := viper.ReadInConfig(); err != nil {
		err := createConfigFIle(homeDir + configPath)
		if err != nil {
			log.Fatalf("Error creating Cfg file, %s", err)
		}
		err = viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Error reading Cfg file, %s", err)
		}
	}
	err = viper.Unmarshal(&Cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}

func createConfigFIle(path string) error {
	sourceFile, err := os.Open(configFileName)
	if err != nil {
		return createConfigFileFromGithub(path)
	}
	defer sourceFile.Close()

	err = os.Mkdir(path, 0755)
	if err != nil {
		log.Fatal(err)
	}

	destinationFile, err := os.Create(path + configFileName)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func createConfigFileFromGithub(path string) error {
	resp, err := http.Get(githubConfigFile)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	err = os.Mkdir(path, 0755)
	if err != nil {
		log.Fatal(err)
	}

	destinationFile, err := os.Create(path + configFileName)
	if err != nil {
		return err
	}

	_, err = destinationFile.Write(data)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}
