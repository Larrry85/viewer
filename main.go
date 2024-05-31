// main.go
package main

import (
	"cars/gofiles" // cars/gofiles: A local package for handling various server functions.
	"net/http"
	"log"
	"os"
	"os/exec"
)

// The code defines an HTTP server in Go that serves a homepage and car details. 
// It reads car data from a JSON file, processes it, and renders it in HTML templates. 
// The server also handles image serving with CORS support and fetches car details 
// from an API based on a car ID

// main(): Starts the application
func main() {

	// Here how to start node mani.js
	go startNode()

	// start Go server
	GoServer()

}


// start Go server
func GoServer() error {
	http.HandleFunc("/", gofiles.HomePage) // "/" handled by HomePage

	
	// FILTER CARS ///////////////////
	http.HandleFunc("/filter", gofiles.FilterPage)
	

	http.HandleFunc("/car-details", gofiles.CarDetailsHandler) // /car-details handled by CarDetailsHandler

	// /static/ serves static files
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	// /img/ serves images with CORS middleware
	http.Handle("/img/", gofiles.CorsMiddleware(http.StripPrefix("/img/", http.HandlerFunc(gofiles.ServeImage))))

	log.Println("Starting server at port 8080")
	return http.ListenAndServe(":8080", nil)
}

func startNode() {
	// Define the command to start the main.js application
	cmd := exec.Command("node", "main.js")

	// Set the working directory to the Node.js app directory
	cmd.Dir = "api"

	// Set the environment variables if needed
	cmd.Env = os.Environ()

	// Start the Node.js application
	err := cmd.Start()
	if err != nil {
		log.Fatalf("Failed to start main.js application: %v", err)
	}
	log.Printf("main.js application started with PID %d", cmd.Process.Pid)

}
