package main

import "github.com/mt-st1/gin-sample-app/router"

func main() {
	router := router.InitializeRouter()

	// Start and run server
	router.Run(":8080")
}
