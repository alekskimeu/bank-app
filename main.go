package main

import (
	"log"
	"os"

	"bankapp/app"
	"bankapp/logger"
)

// Check envs are defined correctly
func sanityCheck() {
	if os.Getenv("SERVER_HOST") == "" || os.Getenv("SERVER_PORT") == "" {
		log.Fatal("One or more configs missing")
	}

}

func main() {

	sanityCheck()

	// log.Println("Starting server...")
	logger.LogInfo("Starting application...")
	app.Boot()
}
