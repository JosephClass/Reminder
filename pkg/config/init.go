package config

import (
	"log"

	"github.com/joho/godotenv"
)

func init() {
	//Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}


}
