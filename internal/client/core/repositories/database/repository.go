// Package database представляет реализацию репозитория данных клиента
package database

import (
	"gorm.io/gorm"
)

// Repository структура репозитория
type Repository struct {
	DB *gorm.DB
}

// NewRepository создание структуры репозитория
func NewRepository(db *gorm.DB) Repository {
	return Repository{
		DB: db,
	}
}

func (r Repository) updateItem(item interface{}) error {
	result := r.DB.Save(item)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}

func (r Repository) insertItem(item interface{}) error {
	result := r.DB.Create(item)
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}
