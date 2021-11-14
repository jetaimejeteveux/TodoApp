package model

//swagger
type Todo struct {
	Id       int    `json:"id" schema:"id"`
	Name     string `json:"name" schema:"name`
	Complete bool   `json:"complete" schema:"complete`
}
