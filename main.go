// main.go
package main

import (
	"cars/gofiles"
	"log"
	//"time"
)

func main() {

	gofiles.GoServer()
/*
	go func() {
		if err := gofiles.ApiServer(); err != nil {
			log.Fatalf("API server failed: %v", err)
		}
	}()

	// Wait short time before GoServer
	time.Sleep(1 * time.Second)

	if err := gofiles.GoServer(); err != nil {
		log.Fatalf("Go server failed: %v", err)
	}*/
}
