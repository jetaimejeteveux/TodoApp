package router

import (
	"TodoApp/controller"
	_ "TodoApp/docs"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

func StartApplication() *mux.Router {
	router := mux.NewRouter()
	//Read all todos
	router.HandleFunc("/todos", controller.GetAllTodos).Methods("GET")
	//create todo
	router.HandleFunc("/todos", controller.CreateTodo).Methods("POST")
	//read todo
	router.HandleFunc("/todos/{id}", controller.GetTodo).Methods("GET")
	//update todo
	router.HandleFunc("/todos/{id}", controller.UpdateTodo).Methods("PUT")
	//delete todo
	router.HandleFunc("/todos/{id}", controller.DeleteTodo).Methods("DELETE")
	//swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)

	return router
}
