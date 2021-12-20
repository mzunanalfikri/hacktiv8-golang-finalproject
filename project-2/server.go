package main

import (
	"log"
	"net/http"
	"project-2/config"
	"project-2/controller"

	"github.com/gorilla/mux"
)

func main() {
	config.StartDB()

	r := mux.NewRouter()

	r.HandleFunc("/users/register", controller.RegisterUser).Methods("POST")
	r.HandleFunc("/users/login", controller.LoginUser).Methods("POST")
	r.HandleFunc("/users/{id}", controller.UpdateUser).Methods("PUT")
	r.HandleFunc("/users/{id}", controller.DeleteUser).Methods("DELETE")

	r.NotFoundHandler = http.HandlerFunc(notfoundHandler)

	log.Println("Server start at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("page not found"))
}
