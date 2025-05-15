package model

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Gender   string `json:"gender"`
	Password string `json:"password"`
}
