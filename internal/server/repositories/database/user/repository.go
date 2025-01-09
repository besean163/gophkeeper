// Package bucket представляет реализацию репозитория хранения пользователей
package user

import (
	models "github.com/besean163/gophkeeper/internal/models/server"
	"gorm.io/gorm"
)

// Repository структура репозитория
type Repository struct {
	db *gorm.DB
}

// NewRepository создание структуры репозитория
func NewRepository(db *gorm.DB) Repository {
	return Repository{
		db: db,
	}
}

// GetUser получение пользователя
func (r Repository) GetUser(id int) (*models.User, error) {
	user := &models.User{}
	r.db.Find(user, id)
	if user.ID == 0 {
		return nil, nil
	}
	return user, nil
}

// GetUserByLogin получение пользователя по логину
func (r Repository) GetUserByLogin(login string) *models.User {
	user := new(models.User)
	r.db.Where("login = ?", login).First(user)

	if user.ID == 0 {
		return nil
	}

	return user
}

// SaveUser сохранение пользователя
func (r Repository) SaveUser(user *models.User) error {
	if user.ID == 0 {
		return r.insertUser(user)
	}

	return r.updateUser(user)
}

func (r Repository) updateUser(user *models.User) error {
	conn := r.db.Save(user)
	err := conn.Error
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) insertUser(user *models.User) error {
	conn := r.db.Create(user)
	err := conn.Error
	if err != nil {
		return err
	}
	return nil
}
