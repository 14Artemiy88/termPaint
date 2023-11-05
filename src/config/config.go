package config

import (
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
	"os"
)

const configFileName = "config.yaml"
const configPath = "/termPaint/"
const githubConfigFile = "https://raw.githubusercontent.com/14Artemiy88/termPaint/main/config.yaml"

type Config struct {
	Background          bool                   `mapstructure:"background"`
	BackgroundColor     map[string]int         `mapstructure:"background_color"`
	DefaultCursor       string                 `mapstructure:"default_cursor"`
	DefaultColor        map[string]int         `mapstructure:"default_color"`
	Pointer             string                 `mapstructure:"pointer"`
	PointerColor        map[string]int         `mapstructure:"pointer_color"`
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

type Screen interface {
	SetConfig(Config)
}

func (c *Config) WithBackground() bool {
	return c.Background
}

func (c *Config) GetBackgroundColor() map[string]int {
	return c.BackgroundColor
}

func (c Config) GetNotificationTime() int {
	return c.NotificationTime
}

func (c *Config) GetImageSaveDirectory() string {
	return c.ImageSaveDirectory
}

func (c *Config) SetImageSaveDirectory(directory string) {
	c.ImageSaveDirectory = directory
}

func InitConfig(s Screen) {
	homeDir, err := os.UserConfigDir()
	if err != nil {
		log.Println("Cannot determine the user's home dir:", err)
	}
	if os.Getenv("ENV") == "dev" {
		viper.SetConfigFile(configFileName)
	} else {
		viper.SetConfigFile(homeDir + configPath + configFileName)
	}
	if err = viper.ReadInConfig(); err != nil {
		err = createConfigFIle(homeDir + configPath)
		if err != nil {
			log.Fatalf("Error creating Cfg file, %s", err)
		}
		err = viper.ReadInConfig()
		if err != nil {
			log.Fatalf("Error reading Cfg file, %s", err)
		}
	}
	var Cfg Config
	err = viper.Unmarshal(&Cfg)
	s.SetConfig(Cfg)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}
}

func createConfigFIle(path string) error {
	// create dir
	err := os.MkdirAll(path, 0755)
	if err != nil {
		return err
	}

	// create file
	destinationFile, err := os.Create(path + configFileName)
	if err != nil {
		return err
	}
	defer destinationFile.Close()

	// read file
	sourceFile, err := os.Open(configFileName)
	if err != nil {
		return createConfigFileFromGithub(destinationFile)
	}
	defer sourceFile.Close()

	// copy file
	_, err = io.Copy(destinationFile, sourceFile)
	if err != nil {
		return err
	}

	return nil
}

func createConfigFileFromGithub(destinationFile *os.File) error {
	// get file by link
	resp, err := http.Get(githubConfigFile)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// read file
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// write file
	_, err = destinationFile.Write(data)
	if err != nil {
		return err
	}

	return nil
}
