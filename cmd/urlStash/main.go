package main

import (
	"log"
	"os"

	"github.com/ary82/urlStash/internal/app"
	"github.com/ary82/urlStash/internal/database"
	"github.com/joho/godotenv"
)

func main() {

	// Load env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbConnStr := os.Getenv("DB_CONN")
	database, err := database.NewPostgresDB(dbConnStr)
	if err != nil {
		log.Fatal(err)
	}
	server := app.NewApiServer(os.Getenv("PORT"), database)
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
