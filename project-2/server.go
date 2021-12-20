package main

import (
	"log"
	"net/http"
	"project-2/config"
	"project-2/controller"
	"project-2/middleware"

	"github.com/gorilla/mux"
)

func main() {
	config.StartDB()

	r := mux.NewRouter()

	userGroup := r.PathPrefix("/users").Subrouter()
	userGroup.Use(middleware.JwtAuth)

	r.HandleFunc("/users/register", controller.RegisterUser).Methods("POST")
	r.HandleFunc("/users/login", controller.LoginUser).Methods("POST")
	userGroup.HandleFunc("", controller.UpdateUser).Methods("PUT")
	userGroup.HandleFunc("", controller.DeleteUser).Methods("DELETE")

	photoGroup := r.PathPrefix("/photos").Subrouter()
	photoGroup.HandleFunc("", controller.CreatePhoto).Methods("POST")
	photoGroup.HandleFunc("", controller.GetPhotos).Methods("GET")
	photoGroup.HandleFunc("/{id}", controller.UpdatePhoto).Methods("PUT")
	photoGroup.HandleFunc("/{id}", controller.DeletePhoto).Methods("DELETE")

	commentGroup := r.PathPrefix("/comments").Subrouter()
	commentGroup.HandleFunc("", controller.AddComments).Methods("POST")
	commentGroup.HandleFunc("", controller.GetComments).Methods("GET")
	commentGroup.HandleFunc("/{id}", controller.UpdateComment).Methods("PUT")
	commentGroup.HandleFunc("/{id}", controller.DeleteComment).Methods("DELETE")

	r.NotFoundHandler = http.HandlerFunc(notfoundHandler)

	log.Println("Server start at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func notfoundHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("page not found"))
}
