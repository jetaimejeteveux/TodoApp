package router

import (
	"TodoApp/controller"

	"github.com/gorilla/mux"
)

func StartApplication() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/todos", controller.GetAllTodos).Methods("GET")
	router.HandleFunc("/todos", controller.CreateTodo).Methods("POST")
	router.HandleFunc("/todos/{id}", controller.GetTodo).Methods("GET")
	router.HandleFunc("/todos/{id}", controller.UpdateTodo).Methods("PUT")
	router.HandleFunc("/todos/{id}", controller.DeleteTodo).Methods("DELETE")

	return router
}
