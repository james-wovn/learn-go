package main

import (
	"go-learn/config"
	"go-learn/router"
)

func main() {
	// Open database connection
	config.OpenDatabase()

	// Start server
	router.StartServer()
}
