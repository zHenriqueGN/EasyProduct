package repository

import (
	"database/sql"
	"errors"

	"github.com/zHenriqueGN/EasyProduct/internal/entity"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) *UserRepository {
	return &UserRepository{DB}
}

func (u *UserRepository) Create(user *entity.User) error {
	stmt, err := u.DB.Prepare("INSERT INTO users (id, name, email, password) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(user.ID, user.Name, user.Email, user.Password)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRepository) FindByEmail(email string) (*entity.User, error) {
	stmt, err := u.DB.Prepare("SELECT id, name, email, password FROM users WHERE email = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row, err := stmt.Query(email)
	if err != nil {
		return nil, err
	}
	if row.Next() {
		var user entity.User
		err = row.Scan(&user.ID, &user.Name, &user.Email, &user.Password)
		if err != nil {
			return nil, err
		}
		return &user, nil
	}
	return nil, ErrUserNotFound
}
