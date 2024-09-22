package apiserver

import (
	"log"

	"github.com/spf13/viper"
)

// Env...
type Config struct {
	AppEnv     string `mapstructure:"APP_ENV"`
	ServerAddr string `mapstructure:"SERVER_ADDRESS"`
	DBPort     int    `mapstructure:"DB_PORT"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPass     string `mapstructure:"DB_PASS"`
	LogLevel   string `mapstructure:"LOG_LEVEL"`
}

// NewEnv...
func NewConfig() *Config {
	config := Config{}
	viper.SetConfigFile("configs/.env")

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
