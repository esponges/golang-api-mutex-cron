package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func main() {
	// Create an instance of the struct
	myInstance := NewInstance()

	// Set the initial value of the global variable
	myInstance.SetGlobalVariable(0)

	// Start the background goroutine to increment the variable
	go incrementGlobalVariable(myInstance)

	// Create a new Gorilla Mux router
	router := mux.NewRouter()

	// Register your API endpoints using Gorilla Mux router
	router.HandleFunc("/", myInstance.Handler).Methods("GET")
	router.HandleFunc("/read", myInstance.ReadHandler).Methods("GET")

	// Start the server
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}

func incrementGlobalVariable(m *MyStruct) {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Increment the global variable
			m.IncrementGlobalVariable()
		}
	}
}
