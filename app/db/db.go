package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	e := godotenv.Load()
	if e != nil {
		log.Fatalf("[Database] Error loading .env file")
	}

	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbType := os.Getenv("DB_TYPE")

	dbUri := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbPort, username, dbName, password)
	db, err := sql.Open(dbType, dbUri)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("[Database] No connection to server")
	}
	
	return db
}
