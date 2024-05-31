package utils

import (
	"fmt"

	"github.com/spf13/viper"
)

func InitConfig(stage string) error {
	viper.AddConfigPath(fmt.Sprintf(`./config/%v/`, stage))
	viper.AutomaticEnv()
	return viper.ReadInConfig()
}

func LoadConfig[T any](config T) (T, error) {
	err := viper.Unmarshal(&config)
	return config, err
}
