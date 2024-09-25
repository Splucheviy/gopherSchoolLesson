package store

// Config...
type Config struct {
	DatabaseURL string `mapstructure:"db_url"`
}

// New...
func NewConfig() *Config {
	return &Config{}
}