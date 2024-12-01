package models

import (
	"time"

	"github.com/go-playground/validator/v10"
)

var validate *validator.Validate

func init() {
	validate = validator.New()
}

type User struct {
	Id       int    `json:"id" db:"id"`
	Email    string `json:"email" validate:"required,email"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" db:"password_hash" binding:"required"`
	RoleId   int    `json:"role_id"`
}

func (i User) Validate() error {
	return validate.Struct(i)
}

type SignUp struct {
	Id          int    `json:"-" db:"id"`
	Email       string `json:"email" validate:"required,email"`
	Name        string `json:"name" binding:"required"`
	Password    string `json:"password" binding:"required"`
	DeviceToken string `json:"device_token" db:"device_token"`
}

func (i SignUp) Validate() error {
	return validate.Struct(i)
}

type Login struct {
	Email       string `json:"email" validate:"required,email"`
	Password    string `json:"password" binding:"required"`
	DeviceToken string `json:"device_token" db:"device_token"`
}

func (i Login) Validate() error {
	return validate.Struct(i)
}

type GetUser struct {
	Id        int    `json:"id" db:"id"`
	Email     string `json:"email" binding:"required"`
	Name      string `json:"name" binding:"required"`
	IsDeleted bool   `json:"is_deleted" db:"is_deleted"`
}

type ChangePasswd struct {
	Password     string `json:"password" binding:"required"`
	NewPassword  string `json:"new_password" binding:"required"`
	NewPassword2 string `json:"new_password2" binding:"required"`
}

type ChangeUserInfo struct {
	Name string `json:"name" binding:"required"`
}

type Profile struct {
	Id    int    `json:"id" db:"id"`
	Email string `json:"email" binding:"required"`
	Name  string `json:"name" binding:"required"`
}

type Token struct {
	Token     string    `db:"token"`
	Revoked   bool      `db:"revoked"`
	CreatedAt time.Time `db:"created_at"`
}

type ResetPasswd struct {
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" binding:"required"`
	Password2 string `json:"password2" binding:"required"`
}

func (i ResetPasswd) Validate() error {
	return validate.Struct(i)
}
