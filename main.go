package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
)

type MyStruct struct {
	globalVariable int
	mutex          sync.Mutex
}

func NewInstance() *MyStruct {
	return &MyStruct{}
}

func (m *MyStruct) SetGlobalVariable(value int) {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.globalVariable = value
}

func (m *MyStruct) GetGlobalVariable() int {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	return m.globalVariable
}

func (m *MyStruct) IncrementGlobalVariable() {
	m.mutex.Lock()
	defer m.mutex.Unlock()
	m.globalVariable++
}

func (m *MyStruct) Handler(w http.ResponseWriter, r *http.Request) {
	// Update the global variable within the handler
	m.IncrementGlobalVariable()

	// Access the global variable
	value := m.GetGlobalVariable()

	// Respond to the request
	fmt.Fprintf(w, "Global variable value: %d", value)
}

func (m *MyStruct) ReadHandler(w http.ResponseWriter, r *http.Request) {
	// Read the global variable
	value := m.GetGlobalVariable()

	// Respond to the request
	fmt.Fprintf(w, "Read global variable value: %d", value)
}

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
