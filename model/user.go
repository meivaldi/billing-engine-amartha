package model

type User struct {
	name       string `json:"name"`
	age        int    `json:"age"`
	workStatus string `json:"working_status"`
}
