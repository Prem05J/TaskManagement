package main

import (
	"log"

	application "github.com/taskManagement/Application"
	db "github.com/taskManagement/DB"
	"github.com/taskManagement/Utils"
)

func main() {
	// DB Operation
	db, err := db.InitializeDatabase()
	if err != nil {
		log.Fatal("Database Connection Not Established")
	}
	port := Utils.GetEnv("API_PORT", "8080")
	// Initialising Application
	app := application.NewAPIServer(":"+port, db)

	if err := app.Run(); err != nil {
		log.Fatal(err)
	}
}
