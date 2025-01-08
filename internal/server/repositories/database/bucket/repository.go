// Package bucket представляет реализацию репозитория данных
package bucket

import (
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

func (r Repository) insertItem(item interface{}) error {
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
