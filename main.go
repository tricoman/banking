package main

import (
	"github.com/tricoman/banking/app"
	"github.com/tricoman/banking/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}
