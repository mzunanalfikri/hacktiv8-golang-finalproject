package main

import (
	"encoding/json"
	"log"
	"net/http"
	_ "project-1/docs"

	"project-1/controllers"
	"project-1/database"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title TOdo Application
// @version 1.0
// @description This todo application API
// @termsOfService http://swagger.io/terms
// @contact.name API Support
// @contact.email testing@gmail.com
// @license.name Apace 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /
func main() {
	database.StartDB()

	r := mux.NewRouter()

	r.HandleFunc("/", HelloWorld).Methods("GET")
	r.HandleFunc("/todos", controllers.GetAllTodos).Methods("GET")

	r.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	log.Println("Server start at http://localhost:8080")
	http.ListenAndServe(":8080", r)
}

func HelloWorld(rw http.ResponseWriter, r *http.Request) {
	json.NewEncoder(rw).Encode("Hello World")

}
