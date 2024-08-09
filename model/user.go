package model

type User struct {
	UserID   uint64 `json:"user_id,omitempty"`
	Name     string `json:"name,omitempty"`
	Age      int    `json:"age,omitempty"`
	WorkType string `json:"working_status,omitempty"`
}
