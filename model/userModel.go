package model

type User struct {
	ID       uint   `json:"user_id"`
	UserName string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
