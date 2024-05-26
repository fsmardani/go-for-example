package config

import (
	"log"

	"github.com/spf13/viper"
	"github.com/joho/godotenv"
)

// The secret key used to sign the JWT, this must be a secure key and should not be stored in the code
const Secret = "secret"

var ConfigValues *viper.Viper

func GetViperConfig() {
	v := viper.New()

	err := godotenv.Load()
    if err != nil {
        log.Println("Error loading .env file")
    }

	// v.SetConfigType("yaml")
	v.SetConfigName("config")
	v.AddConfigPath(".")
	if err := v.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)

	}
	v.AutomaticEnv()

	ConfigValues = v
}
