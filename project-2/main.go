package main

import (
	"log"
	"net/http"
	"project-2/config"

	"github.com/gorilla/mux"
)

func main() {
	config.StartDB()

	r := mux.NewRouter()

	// r.HandleFunc("/", HelloWorld).Methods("GET")

	r.NotFoundHandler = http.HandlerFunc(notfoundHandler)

	log.Println("Server start at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("page not found"))
}
