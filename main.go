package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Starting Calculator...")
	go manager.start() // start goroutine

	http.HandleFunc("/ws", webSocket)
	http.ListenAndServe(GetPort(), nil) // port localhost:12345
}

// For Heroku deployment, which didnt work sadly
func GetPort() string {
	var port = os.Getenv("PORT")
	// Set a default port if there is nothing in the environment
	if port == "" {
		port = "4747"
		fmt.Println("INFO: No PORT environment variable detected, defaulting to " + port)
	}
	return ":" + port
}
