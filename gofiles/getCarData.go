// getCarData.go
package gofiles

import (
	"encoding/json"
	"net/http"
)

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
}
