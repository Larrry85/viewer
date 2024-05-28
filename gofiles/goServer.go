// NEW VERSION

// goServer.go
package gofiles

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"fmt"
)


// data.json path
var directory = "api"
var fileName = "data.json"
var filePath = filepath.Join(directory, fileName)


// start Go server
func GoServer() error {
	http.HandleFunc("/", homePage)

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.Handle("/img/", corsMiddleware(http.StripPrefix("/img/", http.HandlerFunc(serveImage))))

	log.Println("Starting server at port 8080")
	return http.ListenAndServe(":8080", nil)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// read json file
	data, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read file %s: %v", filePath, err), http.StatusInternalServerError)
		return
	}

	// Parse json data into variables
	var carData CarsData
	if err := json.Unmarshal(data, &carData); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse JSON: %v", err), http.StatusInternalServerError)
		return
	}
	
	// Render the HTML template with the car data
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := tmpl.Execute(w, carData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

// ParseCarModels parses car models data from the provided JSON file.
func ParseCarModels() []Car {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", filePath, err)
	}
	var carModels []Car
	if err := json.Unmarshal(data, &carModels); err != nil {
		log.Fatalf("Failed to parse car models data: %v", err)
	}
	return carModels
}

// ParseManufacturers parses manufacturers data from the provided JSON file.
func ParseManufacturers() []Manufacturer {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", filePath, err)
	}
	var manufacturers []Manufacturer
	if err := json.Unmarshal(data, &manufacturers); err != nil {
		log.Fatalf("Failed to parse manufacturers data: %v", err)
	}
	return manufacturers
}

// ParseCategories parses categories data from the provided JSON file.
func ParseCategories() []Category {
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Failed to read file %s: %v", filePath, err)
	}
	var categories []Category
	if err := json.Unmarshal(data, &categories); err != nil {
		log.Fatalf("Failed to parse categories data: %v", err)
	}
	return categories
}
