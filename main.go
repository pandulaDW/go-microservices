package main

import (
	"microservices.com/app"
	"microservices.com/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}
