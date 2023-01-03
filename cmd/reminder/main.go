package main

import (
	"fmt"
	"log"
	"os"
	"reminder/pkg/database"

	"github.com/joho/godotenv"
)

func main()  {
	//Loading .env file 
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	//Initiate database connection
	dbConnection := database.NewConnection()

	//Read port number 
	Port := os.Getenv("SERVERPORT")
	strPort := fmt.Sprintf(":%s",Port)
	
	
}