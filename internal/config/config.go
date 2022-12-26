package config

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/ProovGroup/claim/internal/logger"
	"github.com/spf13/viper"
)

type Config struct {
	AppPort string
}

var ConfigValues *Config

func InitConfig() {
	var err error
	ConfigValues, err = initConfig()
	if err != nil {
		logger.Fatal(err, "Can't init config")
	}
}

func initConfig() (*Config, error) {
	viper.SetDefault("APP_PORT", 8901)

	if os.Getenv("ENVIRONMENT") == "DEV" {
		_, dirname, _, _ := runtime.Caller(0)
		viper.SetConfigName("config")
		viper.SetConfigType("toml")
		viper.AddConfigPath(filepath.Dir(dirname))
		err := viper.ReadInConfig()
		if err != nil {
			return nil, err
		}
	} else {
		viper.AutomaticEnv()
	}

	return &Config{
		AppPort: viper.GetString("APP_PORT"),
	}, nil
}
