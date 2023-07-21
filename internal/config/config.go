package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func GetConfigValue(key string) string {

	//TODO add (?) check if all needed variables is in .env
	//idk how to implement this ^

	err := godotenv.Load(".env")

	if err != nil {
		log.Println("Error loading .env file")
	}

	return os.Getenv(key)
}
