// main.go
package main

import (
	"avitoTest/data/context"
	"avitoTest/shared"
	"log"
)

func main() {
	// Loading the configuration
	conf := shared.LoadConfig()

	// Initialize the logger
	shared.InitLogger(conf)

	// Connecting to the database
	context.ConnectDB()

	log.Printf("Server started on %s", conf.ServerAddress)
}
