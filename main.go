// main.go
package main

import (
	"cars/gofiles"
    "log"
)


func main() {
	go func() {
		if err := gofiles.ApiServer(); err != nil {
			log.Fatalf("API server failed: %v", err)
		}
	}()

	if err := gofiles.GoServer(); err != nil {
		log.Fatalf("Go server failed: %v", err)
	}
}