package service

import (
	"orden/models"
	"orden/pkg/repository"
)

type Authorization interface {
	CreateUser(user models.SignUp) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
	IsAdmin(userId int) (int, error)
	GetUserByEmail(email string) (models.User, error)

	ResetPasswd(userId int, passwd string) error

	GetPasswdById(userId int) (string, error)
	CheckPasswd(userId int, passwd string) error
	UpdateUseranme(userId int, name string) error

	GetUserInfo(userId int) (models.Profile, error)
	GetProfile(userId int) (models.Profile, error)

	SaveToken(token string) error
	InvalidateToken(token string) error
	IsTokenValid(token string) (bool, error)
	DeleteUser(userId int) error

	GetUsers() ([]models.GetUser, error)
	GetAdmins() ([]models.GetUser, error)
	MakeAdmin(userId int) error
	UpdateDeviceToken(token string, userId int) error
	GetRoleId(email string) (int, error)

	RemoveAdmin(userId int) error
}

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
