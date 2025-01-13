// Package bucket представляет реализацию репозитория данных
package bucket

import (
	uuidcontroller "github.com/besean163/gophkeeper/internal/utils/uuid_controller"
	"gorm.io/gorm"
)

// Repository структура репозитория
type Repository struct {
	db *gorm.DB
	uuidcontroller.UUIDController
}

// NewRepository создание структуры репозитория
func NewRepository(db *gorm.DB, uuidController uuidcontroller.UUIDController) Repository {
	return Repository{
		db:             db,
		UUIDController: uuidController,
	}
}

func (r Repository) createItem(item interface{}) error {
	result := r.db.Create(item)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) updateItem(item interface{}) error {
	result := r.db.Save(item)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}
