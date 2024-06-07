// main.go
package main

import (
	"cars/gofiles" // local package for handling go files
	"log"          // logging
	"net/http"     // HTTP server functionality
	"os"           // executing external commands and...
	"os/exec"      // ...managing the operating system environment
)

// The code sets up a Go HTTP server with several routes and handlers for serving a homepage,
// filtering cars, and displaying car details. It also serves static files and images with
// CORS support. Additionally, it starts a Node.js application concurrently using a goroutine.
// This setup allows the Go server to handle web requests while the Node.js application runs
// separately, potentially handling other tasks like API requests

// starts the application
func main() {

	// Starts Node.js application concurrently with the Go server using a goroutine
	go startNode()

	// start Go server
	GoServer()

}

// start Go server
func GoServer() error {
	http.HandleFunc("/", gofiles.HomePage)                     // routes the root URL ("/") to the HomePage
	http.HandleFunc("/filter", gofiles.FilterPage)             // routes "/filter" to the FilterPage
	http.HandleFunc("/car-details", gofiles.CarDetailsHandler) // routes "/car-details" to the CarDetailsHandler

	// serves static files from the "static" directory
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// serves images from the "img" directory with CORS support
	http.Handle("/img/", gofiles.CorsMiddleware(http.StripPrefix("/img/", http.HandlerFunc(gofiles.ServeImage))))

	log.Println("Starting server at port 8080")
	return http.ListenAndServe(":8080", nil) // starts the HTTP server on port 8080
}

func startNode() {
	// creates a command to start the Node.js application with main.js
	cmd := exec.Command("node", "main.js")

	// sets the working directory to the "api" directory where main.js is located
	cmd.Dir = "api"

	// sets the environment variables for the command
	cmd.Env = os.Environ()

	// Start the Node.js application
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Failed to start main.js application: %v", err)
	} // logs the PID of the started Node.js process
	log.Printf("starting node server")
}
