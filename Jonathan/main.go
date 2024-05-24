package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
)

// Manufacturer represents a car manufacturer
type Manufacturer struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Country      string `json:"country"`
	FoundingYear int    `json:"foundingYear"`
}

// PageData holds data for rendering HTML templates
type PageData struct {
	Title        string
	Title2       string
	CarNames     []string
	Manufacturer *Manufacturer
}

func main() {
	// Start the API server in a goroutine
	go startAPIServer()

	// Start the main server
	startMainServer()
}

func startAPIServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/cars", carsAPI)
	mux.HandleFunc("/car/", carAPI)

	log.Println("API server listening on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", mux))
}

func carsAPI(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("api/data.json")
	if err != nil {
		log.Printf("Error opening data.json file: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var data struct {
		Manufacturers []Manufacturer `json:"manufacturers"`
	}

	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		log.Printf("Error decoding data.json: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	var carNames []string
	for _, m := range data.Manufacturers {
		carNames = append(carNames, m.Name)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(carNames)
}

func carAPI(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("api/data.json")
	if err != nil {
		log.Printf("Error opening data.json file: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var data struct {
		Manufacturers []Manufacturer `json:"manufacturers"`
	}

	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		log.Printf("Error decoding data.json: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	carName := strings.TrimPrefix(r.URL.Path, "/car/")
	for _, m := range data.Manufacturers {
		if strings.EqualFold(m.Name, carName) {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(m)
			return
		}
	}

	http.Error(w, "Car not found", http.StatusNotFound)
}

func startMainServer() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/car/", carPage)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Main server listening on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	HomeData := PageData{
		Title:  "Cars Api",
		Title2: "Cars homework",
	}

	// Fetch car names from the API server
	resp, err := http.Get("http://localhost:3000/cars")
	if err != nil {
		http.Error(w, "Error fetching car names", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Error fetching car names", http.StatusInternalServerError)
		return
	}

	err = json.NewDecoder(resp.Body).Decode(&HomeData.CarNames)
	if err != nil {
		http.Error(w, "Error decoding car names", http.StatusInternalServerError)
		return
	}

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	err = tmpl.Execute(w, HomeData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func carPage(w http.ResponseWriter, r *http.Request) {
	carName := strings.TrimPrefix(r.URL.Path, "/car/")

	// Fetch car details from the API server
	resp, err := http.Get(fmt.Sprintf("http://localhost:3000/car/%s", carName))
	if err != nil {
		http.Error(w, "Error fetching car details", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Error fetching car details", http.StatusInternalServerError)
		return
	}

	var manufacturer Manufacturer
	err = json.NewDecoder(resp.Body).Decode(&manufacturer)
	if err != nil {
		http.Error(w, "Error decoding car details", http.StatusInternalServerError)
		return
	}

	CarData := PageData{
		Title:        "Car Details",
		Title2:       manufacturer.Name,
		Manufacturer: &manufacturer,
	}

	tmpl, err := template.ParseFiles("static/car.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)

	err = tmpl.Execute(w, CarData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
