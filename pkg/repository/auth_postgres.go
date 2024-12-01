package repository

import (
	"fmt"
	"orden/models"

	"github.com/jmoiron/sqlx"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user models.SignUp) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (email, name, password_hash, device_token) values($1, $2, $3, $4) RETURNING id", usersTable)
	row := r.db.QueryRow(query, user.Email, user.Name, user.Password, user.DeviceToken)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1 AND password_hash=$2", usersTable)
	err := r.db.Get(&user, query, email, password)
	return user, err
}

func (r *AuthPostgres) GetUserByEmail(email string) (models.User, error) {
	var user models.User
	query := fmt.Sprintf("SELECT id FROM %s WHERE email=$1", usersTable)
	err := r.db.Get(&user, query, email)
	return user, err
}

func (r *AuthPostgres) IsAdmin(userId int) (int, error) {
	var roleId int
	query := fmt.Sprintf("SELECT role_id FROM %s WHERE id=$1", usersTable)
	err := r.db.Get(&roleId, query, userId)
	return roleId, err
}

func (r *AuthPostgres) ResetPasswd(userId int, passwd string) error {

	query := fmt.Sprintf("UPDATE %s SET password_hash = $2 where id =$1", usersTable)
	_, err := r.db.Exec(query, userId, passwd)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthPostgres) GetPasswdById(userId int) (string, error) {
	var passwd string
	query := fmt.Sprintf("SELECT password_hash FROM %s WHERE id=$1 ", usersTable)
	err := r.db.Get(&passwd, query, userId)
	return passwd, err
}

func (r *AuthPostgres) UpdateUseranme(userId int, name string) error {

	query := fmt.Sprintf("UPDATE %s SET name = $1 where id =$2", usersTable)
	_, err := r.db.Exec(query, name, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthPostgres) GetUserInfo(userId int) (models.Profile, error) {
	var userInfo models.Profile
	query := fmt.Sprintf("select id, email,name, notifications from %s where id = $1", usersTable)
	err := r.db.Get(&userInfo, query, userId)
	return userInfo, err
}

func (r *AuthPostgres) SaveToken(token models.Token) error {
	_, err := r.db.NamedExec("INSERT INTO tokens (token, revoked) VALUES (:token, :revoked)", &token)
	return err
}

func (r *AuthPostgres) RevokeToken(token string) error {
	_, err := r.db.Exec("UPDATE tokens SET revoked = TRUE WHERE token = $1", token)
	return err
}

func (r *AuthPostgres) IsTokenRevoked(token string) (bool, error) {
	var revoked bool
	err := r.db.Get(&revoked, "SELECT revoked FROM tokens WHERE token = $1", token)
	if err != nil {
		return false, err
	}
	return revoked, nil
}

func (r *AuthPostgres) DeleteUser(userId int) error {

	query := fmt.Sprintf("UPDATE %s SET deleted_at =  NOW(), is_deleted = true WHERE id = $1", usersTable)
	_, err := r.db.Exec(query, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthPostgres) GetUsers() ([]models.GetUser, error) {
	var users []models.GetUser
	query := fmt.Sprintf("SELECT id, email, name, is_deleted FROM %s where role_id = 1", usersTable)
	err := r.db.Select(&users, query)
	if err != nil {
		return users, err
	}
	return users, err
}

func (r *AuthPostgres) GetAdmins() ([]models.GetUser, error) {
	var users []models.GetUser
	query := fmt.Sprintf("SELECT id, email, name, is_deleted FROM %s where role_id = 2", usersTable)
	err := r.db.Select(&users, query)
	if err != nil {
		return users, err
	}
	return users, err
}

func (r *AuthPostgres) MakeAdmin(userId int) error {

	query := fmt.Sprintf("UPDATE %s SET role_id = 2 WHERE id = $1", usersTable)
	_, err := r.db.Exec(query, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthPostgres) RemoveAdmin(userId int) error {

	query := fmt.Sprintf("UPDATE %s SET role_id = 1 WHERE id = $1", usersTable)
	_, err := r.db.Exec(query, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthPostgres) UpdateDeviceToken(token string, userId int) error {

	query := fmt.Sprintf("UPDATE %s SET device_token = $1 WHERE id = $2", usersTable)
	_, err := r.db.Exec(query, token, userId)
	if err != nil {
		return err
	}
	return nil
}

func (r *AuthPostgres) GetRoleId(email string) (int, error) {
	var roleId int
	query := fmt.Sprintf("select role_id from %s where email = $1", usersTable)
	err := r.db.Get(&roleId, query, email)
	return roleId, err
}
