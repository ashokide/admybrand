package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoDb() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot Load Environment Variables")
	}
	return os.Getenv("MONGOURI")
}

func GetPort() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Cannot Load Environment Variables")
	}
	return os.Getenv("PORT")
}
