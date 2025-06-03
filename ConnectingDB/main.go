package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}
	dbString := os.Getenv("POSTGRES_URL")
	if dbString == "" {
		fmt.Println("DB_CONNECTION_STRING is not set in .env file")
		return
	}
	db, err := sql.Open("postgres", dbString)
	if err != nil {
		fmt.Println("Error connecting to the database:", err)
		return
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Unable to connect to DB:", err)
	}

	fmt.Println("✅ Successfully connected to database!")
	defer db.Close()
}
