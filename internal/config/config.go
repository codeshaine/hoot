package config

import (
	"github.com/codeshaine/hoot/internal/logging"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("$HOME/.config/hoot")
	if err := viper.ReadInConfig(); err != nil {
		logging.Log("config error: " + err.Error())
	}
}
