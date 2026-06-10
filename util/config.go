package util

import (
	"errors"

	"github.com/spf13/viper"
)

type Config struct {
	DBSource      string `mapstructure:"DB_SOURCE"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
}

func LoadConfig(path string) (config Config, err error) {
	v := viper.New()
	v.AddConfigPath(path)
	v.SetConfigName("app")
	v.SetConfigType("env")
	v.AutomaticEnv()
	if err = v.BindEnv("DB_SOURCE"); err != nil {
		return
	}
	if err = v.BindEnv("SERVER_ADDRESS"); err != nil {
		return
	}

	err = v.ReadInConfig()
	if err != nil {
		var notFound viper.ConfigFileNotFoundError
		if !errors.As(err, &notFound) {
			return
		}
	}

	err = v.Unmarshal(&config)
	return
}
