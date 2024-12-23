package models

import (
	"time"

	"github.com/Eius/gochat-shared/vo"
	"github.com/google/uuid"
)

type User struct {
	Id             uuid.UUID         `json:"id"`
	Username       vo.Username       `json:"username"`
	Email          vo.Email          `json:"email"`
	HashedPassword vo.HashedPassword `json:"password"`
	CreatedAt      time.Time         `json:"createdAt"`
	UpdatedAt      time.Time         `json:"updatedAt"`
}

type UserRegister struct {
	Username        vo.Username `json:"username"`
	Email           vo.Email    `json:"email"`
	Password        vo.Password `json:"password"`
	ConfirmPassword vo.Password `json:"confirmPassword"`
}

type UserLogin struct {
	Email    vo.Email    `json:"email"`
	Password vo.Password `json:"password"`
}
