// getCarData.go
package gofiles

import (
	"encoding/json"
	"net/http"
	"errors"
	"strconv"
	"fmt"
	"html/template"
)

// Handles detailed car data fetching and rendering.

// Fetches car details from an API endpoint.
// GetCarDetails fetches car details from another API endpoint.
func GetCarDetails(idStr string) (Car, error) { 
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


// CarDetailsHandler extracts the car ID from query parameters, 
// fetches car details, and renders them using carDetails.html.
// Handles requests for car details and renders an HTML template with the fetched data.
// Tämä sitä varten kun avaa details napin ja carDetails.html sivun
func CarDetailsHandler(w http.ResponseWriter, r *http.Request) {
    // Extract the car ID from the query parameters
    id := r.URL.Query().Get("id") // tämä id tulee html linkistä

    // Perform any necessary validation on the ID

    // Fetch car details based on the ID
    car, err := GetCarDetails(id) // tämä kutsuu GetCarDeatailsByID(), mukana id
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

/*
// Sends an HTTP GET request to http://localhost:3000/api
// to fetch the car data served by the ApiServer
func GetCarData() ([]Car, error) {
	resp, err := http.Get("http://localhost:3000/api")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, err
	}

	var carsData []Car
	err = json.NewDecoder(resp.Body).Decode(&carsData)
	if err != nil {
		return nil, err
	}
	return carsData, nil
}*/