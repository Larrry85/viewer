// NO NEEDED IN NEW VERSION


// handlers.go
package gofiles

import (
	"encoding/json"
	"net/http"
)

func apiHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{
		"models":       "/api/models",
		"categories":   "/api/categories",
		"manufacturers": "/api/manufacturers",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func modelsHandler(w http.ResponseWriter, r *http.Request) {
	/*carsData, err := GetCarData()
	if err != nil {
		http.Error(w, "Failed to fetch car data", http.StatusInternalServerError) // 500
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(carsData.CarModels)*/
}

func categoriesHandler(w http.ResponseWriter, r *http.Request) {
	/*carsData, err := GetCarData()
	if err != nil {
		http.Error(w, "Failed to fetch car data", http.StatusInternalServerError) // 500
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(carsData.Categories)*/
}

func manufacturersHandler(w http.ResponseWriter, r *http.Request) {
	/*carsData, err := GetCarData()
	if err != nil {
		http.Error(w, "Failed to fetch car data", http.StatusInternalServerError) // 500
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(carsData.Manufacturers)*/
}