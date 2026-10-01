package configs

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var config *Config

func Init() error {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return err
	}

	config = &Config{
		Service: Service{
			Port:      os.Getenv("PORT"),
			SecretJWT: os.Getenv("SECRET_JWT"),
		},
		Database: Database{
			DatabaseSourceName: os.Getenv("DATABASE_URL"),
		},
	}
	if config.Service.Port == "" {
		config.Service.Port = ":8080"
	}
	if config.Service.SecretJWT == "" || config.Database.DatabaseSourceName == "" {
		return fmt.Errorf("PORT, SECRET_JWT, and DATABASE_URL must be set")
	}
	return nil
}

func Get() *Config {
	if config == nil {
		config = &Config{}
	}
	return config

}
