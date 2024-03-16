package config

import (
	"github.com/spf13/viper"
)

func GoYamlParse() error{
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

