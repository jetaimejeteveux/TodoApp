package main

import (
	"TodoApp/router"
	"fmt"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", router.StartApplication())
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Listen Server at http://localhost:8080")
	}
}
