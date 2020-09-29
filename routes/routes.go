package routes

import (
	fooController "api/controllers"

	"github.com/gorilla/mux"
)

var Router = mux.NewRouter()

func init() {
	Router.HandleFunc("/api/foo", fooController.SearchHandler).Methods("GET")
	Router.HandleFunc("/api/foo/{id}", fooController.GetHandler).Methods("GET")
	Router.HandleFunc("/api/foo", fooController.PostHandler).Methods("POST")
	Router.HandleFunc("/api/foo/{id}", fooController.DeleteHandler).Methods("DELETE")
}
