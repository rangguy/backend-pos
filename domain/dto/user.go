package dto

import "github.com/google/uuid"

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UserResponse struct {
	UUID        uuid.UUID `json:"uuid"`
	Name        string    `json:"name"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	Role        string    `json:"role,omitempty"`
	PhoneNumber string    `json:"phoneNumber"`
}

type LoginResponse struct {
	User  UserResponse `json:"user"`
	Token string       `json:"token"`
}

type RegisterRequest struct {
	Name            string `json:"name" validate:"required"`
	Username        string `json:"username" validate:"required"`
	Password        string `json:"password" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
	Email           string `json:"email" validate:"required,email"`
	PhoneNumber     string `json:"phoneNumber" validate:"required,number"`
	RoleID          uint
}

type RegisterResponse struct {
	User UserResponse `json:"user"`
}

type UpdateRequest struct {
	Name            string  `json:"name" validate:"required"`
	Username        string  `json:"username" validate:"required"`
	Password        *string `json:"password,omitempty"`
	ConfirmPassword *string `json:"confirmPassword,omitempty"`
	Email           string  `json:"email" validate:"required,email"`
	PhoneNumber     string  `json:"phoneNumber" validate:"required,number"`
	RoleID          uint
}
