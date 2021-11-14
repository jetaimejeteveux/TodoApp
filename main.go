package main

import (
	"TodoApp/router"
	"fmt"
	"net/http"
)

// @title Todo API
// @version 1.0
// @description This is a API for Todos
// @host localhost:8080
// @BasePath /

func main() {
	err := http.ListenAndServe(":8080", router.StartApplication())
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Listen Server at http://localhost:8080")
	}

}
