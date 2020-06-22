package user

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserID   string `json:"user_id"`
	Email    string `json:"email"`
	Address  string `json:"address"`
	Password string `json:"password"`
}

type Response struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

type DatabaseResponse struct {
	Success bool `json:"success"`
	Data    User `json:"data"`
}

type DeleteUserResponse struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	JWT     string `json:"token"`
}
