package models

import "github.com/Eius/gochat-shared/vo"

type UserRegister struct {
	Username        vo.Username `json:"username"`
	Email           vo.Email    `json:"email"`
	Password        vo.Password `json:"password"`
	ConfirmPassword vo.Password `json:"confirmPassword"`
}

type UserLogin struct {
	Username vo.Username `json:"username"`
	Password vo.Password `json:"password"`
}
