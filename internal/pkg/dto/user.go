package dto

type User struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
