package repository

import "github.com/zHenriqueGN/EasyProduct/internal/entity"

type UserInterface interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
}
