package model

import "database/sql"

//Todo represents the model for todo list
type Todo struct {
	Id       int          `json:"id"`
	Name     string       `json:"name"`
	Complete sql.NullBool `json:"complete"`
}
