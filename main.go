package main

import (
	"microservices.com/app"
	"microservices.com/helpers"
	"microservices.com/logger"
)

func main() {
	logger.Info("Starting the application")

	// set the environment variables
	helpers.DotEnv()

	app.Start()
}
