package repository

import (
	"orden/models"

	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user models.SignUp) (int, error)
	GetUser(username, password string) (models.User, error)
	IsAdmin(userId int) (int, error)
	GetUserByEmail(email string) (models.User, error)
	ResetPasswd(userId int, passwd string) error

	GetPasswdById(userId int) (string, error)
	UpdateUseranme(userId int, name string) error

	GetUserInfo(userId int) (models.Profile, error)

	DeleteUser(userId int) error

	GetUsers() ([]models.GetUser, error)
	GetAdmins() ([]models.GetUser, error)
	MakeAdmin(userId int) error

	GetRoleId(email string) (int, error)

	RemoveAdmin(userId int) error

	SaveToken(token models.Token) error
	RevokeToken(token string) error
	IsTokenRevoked(token string) (bool, error)
	UpdateDeviceToken(token string, userId int) error
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
