package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func NewConnection() *sql.DB {
	
	//Load .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file")
	}

	//Load data from .env file 
	Host := os.Getenv("POSTGRESHOST")
	Port := os.Getenv("POSTGRESPORT")
	User := os.Getenv("POSTGRESUSER")
	Password := os.Getenv("POSTGRESPASSWORD")
	Database := os.Getenv("POSTGRESDATABASE")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	Host,Port,User,Password,Database)

	//Open connection to database server
	conn,err := sql.Open("postgres",dsn)

	if err != nil {
		log.Fatalln("Database error : Can't Open Connection")
	}

	if err := conn.Ping(); err != nil {
		log.Fatalln("Database error : Database Unreachable")
	}

	log.Println("Database connection is successfully")
	return conn
}