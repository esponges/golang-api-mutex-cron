package main

import (
	"fmt"
	"net/http"
)

func (myInstance *MyStruct) Handler(w http.ResponseWriter, r *http.Request) {
	// Update the global variable within the handler
	myInstance.IncrementGlobalVariable()

	// Access the global variable
	value := myInstance.GetGlobalVariable()

	// Respond to the request
	fmt.Fprintf(w, "Global variable value: %d", value)
}

func (myInstance *MyStruct) ReadHandler(w http.ResponseWriter, r *http.Request) {
	// Read the global variable
	value := myInstance.GetGlobalVariable()

	// Respond to the request
	fmt.Fprintf(w, "Read global variable value: %d", value)
}
