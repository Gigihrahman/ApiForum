package configs

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var config *Config

func Init() error {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return err
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	portNumber, err := strconv.Atoi(port)
	if err != nil || portNumber < 1 || portNumber > 65535 {
		return fmt.Errorf("PORT must be a number between 1 and 65535")
	}

	config = &Config{
		Service: Service{
			Port:      fmt.Sprintf(":%d", portNumber),
			SecretJWT: os.Getenv("SECRET_JWT"),
		},
		Database: Database{
			DatabaseSourceName: os.Getenv("DATABASE_URL"),
		},
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
