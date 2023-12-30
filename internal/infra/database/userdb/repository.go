package userdb

import (
	"github.com/tomazcx/go-chat-app/internal/entity"
	"github.com/tomazcx/go-chat-app/internal/infra/database"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func (r *UserRepository) FindByLogin(login string) (*entity.User, error) {
	user := &entity.User{}
	err :=  r.db.First(user, "login = ?", login).Error

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) Create(user *entity.User) error {
	return r.db.Create(&user).Error
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		db: database.GetDBInstance(),
	}
}
