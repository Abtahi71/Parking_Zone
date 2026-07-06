package dto

import "gotickets/internal/domain/user/types"

type RegisterResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    RegData  `json:"data"`
}

type RegData struct{
	Id        uint   `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      types.Role `json:"role"`
	CreatedAt string `json:"created_at"`
	UptedAt   string `json:"updated_at"`
}
type LoginResponse struct {
	Success bool      `json:"success"`
	Message string    `json:"message"`
	Data    LoginData `json:"data"`
}

type LoginData struct {
	Token string   `json:"token"`
	User  UserData `json:"user"`
}
type UserData struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  types.Role `json:"role"`
}
