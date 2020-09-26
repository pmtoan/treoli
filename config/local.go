package config

import (
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"os"
	"path"
)

func LoadConfig() bool {
	home, err := homedir.Dir()
	if err != nil{
		panic(err)
	}

	viper.AddConfigPath(home)
	viper.SetConfigName(".treoli")
	viper.SetConfigType("yml")
	err = viper.ReadInConfig() // Find and read the config file
	if err != nil { // Handle errors reading the config file
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			_, err := os.Stat(path.Join(home, ".treoli.yml"))
			if !os.IsExist(err) {
				if _, err := os.Create(path.Join(home, ".treoli.yml")); err != nil {
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
