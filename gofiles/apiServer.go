// NEW VERSION

// apiServer.go
package gofiles

import (
	"encoding/json"
	"html/template"
	"net/http"
	"strconv"
	//"log"
	//"os/exec"
	//"time"
	"path/filepath"
	"mime"
	"errors"
	"fmt"
)
/*
// start api server
func ApiServer() {

	// Wrap the APIHandler with the corsMiddleware
	http.Handle("/api", corsMiddleware(http.HandlerFunc(APIHandler)))

	// other API endpoints
	http.HandleFunc("/api/models", func(w http.ResponseWriter, r *http.Request) {
		JSONHandler(w, r, GetCarData)
	})
	http.HandleFunc("/api/categories", func(w http.ResponseWriter, r *http.Request) {
		JSONHandler(w, r, GetCarData)
	})
	http.HandleFunc("/api/manufacturers", func(w http.ResponseWriter, r *http.Request) {
		JSONHandler(w, r, GetCarData)
	})

	http.HandleFunc("/car-details", carDetailsHandler)
	log.Println("Starting API server at port 3000")

	// Command to start the Node.js server
	cmd := exec.Command("node", "main.js")	
	// Start the Node.js server asynchronously
	if err := cmd.Start(); err != nil {
		log.Fatalf("Failed to start Node.js server: %v", err)
	}	
	// Wait for a short duration for the Node.js server to start
	time.Sleep(2 * time.Second)	
	// Check if the Node.js server is still running
	if cmd.ProcessState != nil && cmd.ProcessState.Exited() {
		log.Fatalf("Node.js server exited with error: %v", cmd.ProcessState)
	}	
	// Start the API server
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatalf("API server failed: %v", err)
	}
}*/

// Servereiden pitäisi tehdä yhteistyötä toistensa kanssa
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Allow requests from localhost:8080
		w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// If it's a preflight request, return early with status code 200
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}

// serve images
func serveImage(w http.ResponseWriter, r *http.Request) {
	// Join the directory with the requested file path
	filePath := filepath.Join("api", "img", filepath.Base(r.URL.Path))
	ext := filepath.Ext(filePath)
	mimeType := mime.TypeByExtension(ext)

	if mimeType != "" {
		w.Header().Set("Content-Type", mimeType)
	}
	http.ServeFile(w, r, filePath)
}

/*
// handles requests to the /api endpoint
func APIHandler(w http.ResponseWriter, r *http.Request) {
	// Fetch car data
	carsData, err := GetCarData()
	if err != nil {
		http.Error(w, "Failed to fetch car data", http.StatusInternalServerError)
		return
	}
	// Set response header
	w.Header().Set("Content-Type", "application/json")
	// Encode and send JSON response
	json.NewEncoder(w).Encode(carsData)
}


// Fetch car data
func JSONHandler(w http.ResponseWriter, r *http.Request, getData func() ([]Car, error)) {
	carsData, err := getData()
	if err != nil {
		http.Error(w, "Failed to fetch car data", http.StatusInternalServerError)
		return
	}
	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	// Encode and write data to the response writer
	if err := json.NewEncoder(w).Encode(carsData); err != nil {
		http.Error(w, "Failed to encode car data", http.StatusInternalServerError)
		return
	}
}*/

// Tämä sitä varten kun avaa details napin ja carDetails.html sivun
func carDetailsHandler(w http.ResponseWriter, r *http.Request) {
    // Extract the car ID from the query parameters
    id := r.URL.Query().Get("id") // tämä id tulee html linkistä

    // Perform any necessary validation on the ID

    // Fetch car details based on the ID
    car, err := GetCarDetailsByID(id) // tämä kutsuu GetCarDeatailsByID(), mukana id
    if err != nil {
        http.Error(w, "Car not found", http.StatusNotFound)
        return
    }

    // Render the car details page using a template
    tmpl, err := template.ParseFiles("static/carDetails.html") // oikean auton data tällä sivulla
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }

    // Execute the template with the car data
    if err := tmpl.Execute(w, car); err != nil { // oikean auton data tällä sivulla
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
}

func GetCarDetailsByID(idStr string) (Car, error) { 
    // Convert the string ID to an integer
    id, err := strconv.Atoi(idStr) // tämän pitäis muuttaa string to int
    if err != nil {
        return Car{}, errors.New("invalid car ID")
    }

    // Make a request to the Go server endpoint to fetch car details
    resp, err := http.Get(fmt.Sprintf("http://localhost:3000/api/models/%d", id))
    if err != nil {
        return Car{}, err
    }
    defer resp.Body.Close()

    // Decode the response body into a Car struct
    var car Car
    if err := json.NewDecoder(resp.Body).Decode(&car); err != nil {
        return Car{}, err
    }

    return car, nil // ...lähettää sen auton data carDetails.html sivulle
}
