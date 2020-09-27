package config

import (
	"fmt"
	"os"
	"path"

	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

func LoadConfig() bool {
	home, err := homedir.Dir()
	if err != nil {
		panic(err)
	}
	absoluteConfigFolder := path.Join(home, FolderPath)

	viper.AddConfigPath(absoluteConfigFolder)
	_, err = os.Stat(absoluteConfigFolder)
	if err != nil {
		_ = os.Mkdir(absoluteConfigFolder, 0755)
	}

	viper.SetConfigName("treoli")
	viper.SetConfigType("yml")
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil {            // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			_, err := os.Stat(path.Join(absoluteConfigFolder, FileName))
			if !os.IsExist(err) {
				if _, err := os.Create(path.Join(absoluteConfigFolder, FileName)); err != nil {
					fmt.Print(err)
					panic("Can not create config in your HOME directory")
				}
			}
			return false
		} else {
			panic(err)
		}
	}
	return true
}

func SaveConfig(key, value string) {
	viper.Set(key, value)
	_ = viper.WriteConfig()
}
