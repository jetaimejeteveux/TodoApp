package controller

import (
	"TodoApp/database"
	"TodoApp/model"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
)

//GetTodos godoc
//@Summary Get details of all todo
//@Tags todo
//@Accept json
//@Produce json
//@Success 200 {array} model.Todo
//@Router /todos [get]

func GetAllTodos(w http.ResponseWriter, r *http.Request) {
	//Connect DB
	w.Header().Set("Content-Type", "application/json")
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

//CreateTodo godoc
//@Summary Create a new todo
//@Description Create a new todo
//@Tags todo
//@Accept json
//Produce json
//@Param todo body model.Todo true "Create todos"
//@Success 200 {object} model.Todo
//@Router /todos [post]
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
	w.Write([]byte("Create Data Succesfull..."))

}

// GetTodo godoc
// @Summary get detail todo
// @Description get todo
// @Tags todo
// @Produce json
// @Accept json
// @Success 200 {object} model.Todo
// @Failure 400 {string} string "error"
// @Router /todo/{todoId} [get]
// @Param todoId path int true "id todo"

func GetTodo(w http.ResponseWriter, r *http.Request) {
	// Database Connection
	w.Header().Set("Content-Type", "application/json")
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	//Read HTTP Parameter
	params := mux.Vars(r)
	paramId := params["id"]

	//Exec Query Real All Todo
	rows, err := db.Query("select * from todo where id = ?", paramId)
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

//UpdateTodo godoc
//@Summary Update a todo
//@Description Update todo
//@Tags todo
//@Accept json
//Produce json
//@Param todo body model.Todo true "Update todos"
//@Success 200 {object} model.Todo
//@Failure 400 {string} string "error"
//@Router /todo [put]
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	// Databswagase Connection
	w.Header().Set("Content-Type", "application/json")
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	//Read HTTP Parameter
	params := mux.Vars(r)
	paramId := params["id"]

	//Read Req Body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	var todo model.Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	_, err = db.Exec("update todo set name = ?, complete = ? where id = ?", todo.Name, todo.Complete, paramId)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	w.Write([]byte("Update Data Succrsfull..."))

}

//DeleteTodo godoc
//@Summary Delete todo
//@Description Delete todo
//@Success 200 {string} string "ok"
//@Failure 400 {string} string "error"
//@Router /todo [delete]
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	// Database Connection
	w.Header().Set("Content-Type", "application/json")
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	//Read HTTP Parameter
	params := mux.Vars(r)
	paramId := params["id"]

	_, err = db.Exec("delete from todo where id =?", paramId)
	if err != nil {
		fmt.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte("Delete Data Succrsfull..."))
}
