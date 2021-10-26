package model

import "database/sql"

//Todo represents the model for todo list
type Todo struct {
	Id       int          `json:"id" schema:"id"`
	Name     string       `json:"name" schema:"name`
	Complete sql.NullBool `json:"complete" schema:"complete`
}
