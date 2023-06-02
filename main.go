package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/mux"
	"github.com/robfig/cron/v3"
)

// TheStruct for the routing handler
type IncrementJob struct {
	myInstance *MyStruct
	someArg    string
}

func (j IncrementJob) Run() {
	j.myInstance.IncrementGlobalVariable()
	fmt.Println("Incremented global variable", j.myInstance.GetGlobalVariable())
	fmt.Println("someArg", j.someArg)
	// could it request the increment endpoint here instead of calling the method?
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
	router.HandleFunc("/increment", myInstance.Handler).Methods("GET")
	router.HandleFunc("/read", myInstance.ReadHandler).Methods("GET")

	// Create a new cron instance
	c := cron.New()

	// Add a new cron job that runs the IncrementGlobalVariable method every 10 seconds
	incrementJob := IncrementJob{myInstance: myInstance, someArg: "someArgh"}
	fmt.Println("Adding cron job")
	c.AddJob("@every 10s", incrementJob)

	// Start the cron instance
	c.Start()

	// Start the server
	fmt.Println("Starting server on port 8080")
	http.ListenAndServe(":8080", nil)
}

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

func incrementGlobalVariable(m *MyStruct) {
	ticker := time.NewTicker(time.Hour)
	defer ticker.Stop()

	for range ticker.C {
		// Increment the global variable
		m.IncrementGlobalVariable()
	}
}

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
