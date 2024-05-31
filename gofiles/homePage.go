// NEW VERSION

// goServer.go
package gofiles

import (
	"encoding/json"
	"html/template"
	"net/http"
	"os"
	"path/filepath"
	"fmt"
)

// data.json path
var directory = "api"
var fileName = "data.json"
var filePath = filepath.Join(directory, fileName)

// HomePage reads data from data.json, parses it into CarsData, 
// and renders it using index.html
func HomePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	// Reads and parses JSON data from data.json
	// read json file
	data, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to read file %s: %v", filePath, err), http.StatusInternalServerError)
		return
	}

	// The JSON data is then parsed into a CarsData struct.
	// Parse json data into variables
	// contains slices of car models, manufacturers, and categories.
	var carData CarsData
	if err := json.Unmarshal(data, &carData); err != nil {
		http.Error(w, fmt.Sprintf("Failed to parse JSON: %v", err), http.StatusInternalServerError)
		return
	}

	// Render the HTML template with the car data (parsed data)
	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// The final HTML is sent to the client's browser.
	// When tmpl.Execute(w, carData) is called, the template 
	// engine replaces {{.ID}} with the actual ID of each car
	// in the CarModels slice
	if err := tmpl.Execute(w, carData); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}