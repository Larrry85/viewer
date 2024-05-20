//apiServer.go
package gofiles

import (
    "net/http"
    "encoding/json"
    "os"
    "path/filepath"
)

// Open your browser and navigate to http://localhost:3001/cars
// You should see the JSON data
func ApiServer() error {

   // Serves the car data stored in the data.json file at the /cars endpoint
    http.HandleFunc("/cars", func(w http.ResponseWriter, r *http.Request) {
        dataFilePath := filepath.Join("api", "data.json")

        file, err := os.Open(dataFilePath)
        if err != nil {
            http.Error(w, "Unable to open data file", http.StatusInternalServerError) // 500
            return
        }
        defer file.Close()

        var carsData CarsData

        // decodes file contents into a CarsData struct...
        decoder := json.NewDecoder(file)
        if err = decoder.Decode(&carsData); err != nil {
            http.Error(w, "Unable to parse data file", http.StatusInternalServerError) // 500
            return
        }
        w.Header().Set("Content-Type", "application/json")
        // .. and responds with this data in JSON format
        json.NewEncoder(w).Encode(carsData)
    })

    return http.ListenAndServe(":3001", nil) // Listens port 3001
}