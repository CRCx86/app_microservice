package dto

import "strings"

type User struct {
	Id       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) Validate() (message string, result bool) {

	if !strings.Contains(u.Email, "@") { // TODO: заменить на регулярки
		return "Email address is required", false
	}

	if len(u.Password) < 6 {
		return "Password required and must be greater then 6 symbols", false
	}

	return "", true
}
