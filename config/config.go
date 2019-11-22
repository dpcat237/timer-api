package config

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// ModeProd defines production mode
const ModeProd = "prod"

// Config defines configuration parameters
type Config struct {
	DbName   string `envconfig:"DB_NAME" default:"timer.db"`
	HTTPport string `envconfig:"HTTP_PORT" default:"8080"`
	Mode     string `envconfig:"MODE" default:"dev"`
}

// LoadConfigData loads environment parameters
func LoadConfigData() Config {
	var cnf Config
	if err := envconfig.Process("timer", &cnf); err != nil {
		panic(fmt.Sprintf("Failed reading environment variables: %s", err))
	}
	return cnf
}
