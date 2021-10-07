package model

type Todo struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Complete bool   `json:"complete"`
}
