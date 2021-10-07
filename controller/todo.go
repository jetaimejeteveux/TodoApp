package controller

import (
	"TodoApp/database"
	"TodoApp/model"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	//Connect DB
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	//Exec Query Real All Todo
	rows, err := db.Query("select * from todo")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer rows.Close()

	//Merubah Hasil rews Query ke struct
	var result []model.Todo
	for rows.Next() {
		var temp = model.Todo{}
		err := rows.Scan(&temp.Id, &temp.Name, &temp.Complete)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		result = append(result, temp)
	}

	//Merubah struct ke json untuk di tampilkan
	res, _ := json.Marshal(result)
	w.Write(res)

}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	//Connect DB
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	//Read Request Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//Merubah Request Body to struct
	var todo model.Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//Insert to database
	_, err = db.Exec("insert into todo(name) value(?)", todo.Name)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	w.Write([]byte("Create Data Succrsfull..."))

}

func GetTodo(w http.ResponseWriter, r *http.Request) {

}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {

}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {

}
