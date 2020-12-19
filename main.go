package main

import (
	"github.com/diogoqds/banking/app"
	"github.com/diogoqds/banking/logger"
)

func main() {
	logger.Info("Starting the application")
	app.Start()
}
