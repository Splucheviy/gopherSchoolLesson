package apiserver

import (
	"log"

	"github.com/spf13/viper"
)

// Env...
type Config struct {
	AppEnv      string `mapstructure:"app_env"`
	ServerAddr  string `mapstructure:"server_address"`
	LogLevel    string `mapstructure:"log_level"`
	DatabaseURL string `mapstructure:"db_url"`
	SessionKey  string `mapstructure:"session_key"`
}

// NewEnv...
func NewConfig() *Config {
	config := Config{}
	viper.SetConfigFile("configs/apiserver.toml")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Can't find the file .env : ", err)
	}

	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatal("Environment can't be loaded: ", err)
	}

	if config.AppEnv == "development" {
		log.Println("The App is running in development env")
	}

	return &config
}
