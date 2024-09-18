package main

import (
	"fintrack/internal/infra/db"
	"log"
)

func main() {
	database, err := db.InitializeDB()
	if err != nil {
		log.Fatal("Failed to initialize the database: ", err)
	}
	defer database.Close()

	if err := db.RunMigrations(database); err != nil {
		log.Fatal("Failed to run migrations: ", err)
	}
}
