package controller

import (
	"TodoApp/database"
	"TodoApp/model"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
)

var decoder = schema.NewDecoder()

// GetTodos godoc
// @Summary Get details of all Todos
// @Description Get details of all Todos
// @Tags Todos
// @Accept json
// @Produce json
// @Success 200 {array} model.Todo
// @Router /todos [get]
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

// CreateTodos godoc
// @Summary Create a new Todo
// @Description Create a new Todos
// @Tags Todos
// @Accept json
// @Produce json
// @Success 200 {array} model.Todo
// @Router /todos [post]
// @Param name query string true "name for new Todo"
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	//Connect DB
	db, err := database.Connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer db.Close()

	//Read Parameter
	var todo model.Todo

	err = decoder.Decode(&todo, r.URL.Query())
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

// GetTodo godoc
// @Summary Get details of Todos
// @Description Get details of Todo by ID
// @Tags Todos
// @Accept json
// @Produce json
// @Success 200 {array} model.Todo
// @Router /todos/{id} [get]
// @Param id path int true "id for get todo"
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

// UpdateTodo godoc
// @Summary Update Todo
// @Description Update Todo by ID
// @Tags Todos
// @Accept json
// @Produce json
// @Success 200 {array} model.Todo
// @Router /todos/{id} [put]
// @Param id path int true "id for update todo"
// @Param name query model.Todo true "Name for update todo"
// // @Param complete query bool true "Status Complete for update todo"
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
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

	var todo model.Todo

	err = decoder.Decode(&todo, r.URL.Query())
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

// DeleteTodo godoc
// @Summary Delete Todo
// @Description Delete Todo by ID
// @Tags Todos
// @Accept json
// @Produce json
// @Success 200 {array} model.Todo
// @Router /todos/{id} [delete]
// @Param id path int true "id for delete todo"
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
