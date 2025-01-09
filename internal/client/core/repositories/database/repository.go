// Package database представляет реализацию репозитория данных клиента
package database

import (
	uuidcontroller "github.com/besean163/gophkeeper/internal/utils/uuid_controller"
	"gorm.io/gorm"
)

// Repository структура репозитория
type Repository struct {
	uuidcontroller.UUIDController
	DB *gorm.DB
}

// NewRepository создание структуры репозитория
func NewRepository(db *gorm.DB, uuidController uuidcontroller.UUIDController) Repository {
	return Repository{
		UUIDController: uuidController,
		DB:             db,
	}
}

func (r Repository) createItem(item interface{}) error {
	result := r.DB.Create(item)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) updateItem(item interface{}) error {
	result := r.DB.Save(item)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}
