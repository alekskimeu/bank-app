package main

import (
	"bankapp/app"
	"bankapp/logger"
)

func main() {
	// log.Println("Starting server...")
	logger.LogInfo("Starting application...")
	app.Boot()
}
