package main

import (
	"fmt"
	"log"
	"net/http"

	muxtrace "gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

func main() {
	//initialize the tracer
	tracer.Start(
		tracer.WithEnv("test"),
		tracer.WithServiceVersion("1.0.0"),
		tracer.WithRuntimeMetrics(),
	)
	defer tracer.Stop()

	// Create a new router
	router := muxtrace.NewRouter(muxtrace.WithServiceName("source_code_integration_example"))

	// Define your routes
	router.HandleFunc("/", homeHandler).Methods("GET")
	router.HandleFunc("/users", usersHandler).Methods("GET")

	// Start the server
	fmt.Println("Server started on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the home page!")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "List of users:")
	fmt.Fprintln(w, "- User 1")
	fmt.Fprintln(w, "- User 2")
	fmt.Fprintln(w, "- User 3")
}
