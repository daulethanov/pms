package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func GoEnvParse() error {
    if err := godotenv.Load("./config/.env"); err != nil {
        return err
    }
    return nil
}

var Config struct {
	Mongo struct {
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		Username string `mapstructure:"username"`
		Password string `mapstructure:"password"`
	} `mapstructure:"mongo"`

	Minio struct {
		Endpoint   string `mapstructure:"endpoint"`
		SecretKey  string `mapstructure:"secret-key"`
		AccessKey  string `mapstructure:"access-key"`
	} `mapstructure:"minio"`

	SMTP struct {
		SMTP     string `mapstructure:"smtp"`
		Port     string `mapstructure:"port"`
		Email    string `mapstructure:"email"`
		Password string `mapstructure:"password"`
	} `mapstructure:"smtp"`
}

func GoYamlParse() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	for _, k := range viper.AllKeys() {
		v := viper.GetString(k)
		viper.Set(k, os.ExpandEnv(v))
	}

	if err := viper.Unmarshal(&Config); err != nil {
		fmt.Printf("Unable to decode into struct %s\n", err)
		return err
	}

	return nil
}
