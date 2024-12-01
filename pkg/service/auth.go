package service

import (
	"crypto/sha1"

	"errors"
	"fmt"

	"orden/models"
	"orden/pkg/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	salt       = "asdassdfsdfdfdfsdfs"
	signingKey = "sdfsdfsdsdada"
	tokenTTL   = 12 * time.Hour
)

type tokenClaims struct {
	jwt.RegisteredClaims
	UserId int `json:"user_id"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.SignUp) (int, error) {
	checkUser, err := s.repo.GetUserByEmail(user.Email)
	if err != nil {
		checkUser.Id = 0
	}
	if checkUser.Id > 0 {
		return 0, errors.New("с такой почтой клиент существует")
	}

	user.Password = s.generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(email, s.generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{

		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().AddDate(10, 0, 0)),
			//ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			//IssuedAt:  time.Now().Unix(),
		},

		user.Id,
	})
	/*
		tokenString, err := token.SignedString([]byte(signingKey))
		if err != nil {
			return "", err
		}

		err = s.repo.SaveToken(models.Token{Token: tokenString, Revoked: false})
		if err != nil {
			return "", err
		}
	*/
	return token.SignedString([]byte(signingKey))

}

func (s *AuthService) ParseToken(accessToken string) (int, error) {

	token, err := jwt.ParseWithClaims(accessToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not  of  type *tokenCalims")

	}
	return claims.UserId, nil
}

func (s *AuthService) generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) IsAdmin(userId int) (int, error) {
	return s.repo.IsAdmin(userId)
}

func (s *AuthService) GetUserByEmail(email string) (models.User, error) {
	return s.repo.GetUserByEmail(email)
}

func (s *AuthService) ResetPasswd(userId int, passwd string) error {
	password := s.generatePasswordHash(passwd)
	return s.repo.ResetPasswd(userId, password)
}

func (s *AuthService) GetPasswdById(userId int) (string, error) {
	return s.repo.GetPasswdById(userId)
}

func (s *AuthService) CheckPasswd(userId int, passwd string) error {

	userPasswd, err := s.repo.GetPasswdById(userId)
	if err != nil {
		return err
	}

	if userPasswd != s.generatePasswordHash(passwd) {
		return errors.New("user password is not correct")
	}

	return nil
}

func (s *AuthService) UpdateUseranme(userId int, name string) error {
	return s.repo.UpdateUseranme(userId, name)
}

func (s *AuthService) GetUserInfo(userId int) (models.Profile, error) {
	return s.repo.GetUserInfo(userId)
}

func (s *AuthService) GetProfile(userId int) (models.Profile, error) {
	user, err := s.repo.GetUserInfo(userId)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *AuthService) SaveToken(token string) error {
	return s.repo.SaveToken(models.Token{Token: token, Revoked: false})
}

func (s *AuthService) InvalidateToken(token string) error {
	return s.repo.RevokeToken(token)
}

func (s *AuthService) IsTokenValid(token string) (bool, error) {
	revoked, err := s.repo.IsTokenRevoked(token)
	if err != nil {
		return false, err
	}
	return !revoked, nil
}

func (s *AuthService) DeleteUser(userId int) error {
	return s.repo.DeleteUser(userId)
}

func (s *AuthService) GetUsers() ([]models.GetUser, error) {
	return s.repo.GetUsers()
}
func (s *AuthService) GetAdmins() ([]models.GetUser, error) {
	return s.repo.GetAdmins()
}

func (s *AuthService) MakeAdmin(userId int) error {
	return s.repo.MakeAdmin(userId)
}

func (s *AuthService) UpdateDeviceToken(token string, userId int) error {
	return s.repo.UpdateDeviceToken(token, userId)
}

func (s *AuthService) GetRoleId(email string) (int, error) {
	return s.repo.GetRoleId(email)
}

func (s *AuthService) RemoveAdmin(userId int) error {
	return s.repo.RemoveAdmin(userId)
}
