// getCarData.go
package gofiles

import (
	"encoding/json"
	"net/http"
)

// Sends an HTTP GET request to http://localhost:3001/cars
// to fetch the car data served by the ApiServer
func GetCarData() (CarsData, error) {
	resp, err := http.Get("http://localhost:3001/cars")
	if err != nil {
		return CarsData{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK { // 200
		return CarsData{}, err
	}

	var carsData CarsData
	// Decodes the JSON response into a CarsData struct and returns it
	err = json.NewDecoder(resp.Body).Decode(&carsData)
	if err != nil {
		return CarsData{}, err
	}
	return carsData, nil
}
