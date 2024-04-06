package main

import (
	"log"
	"os"

	"github.com/ary82/urlStash/api"
	"github.com/ary82/urlStash/database"
	"github.com/joho/godotenv"
)

func main() {

	// Load env vars
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	dbConnStr := "postgres://ary:123@localhost:5431/urlStash"
	database, err := database.Connect(dbConnStr)
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewApiServer(os.Getenv("PORT"), database)
	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
