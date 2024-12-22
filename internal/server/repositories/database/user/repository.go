package user

import (
	"fmt"

	"github.com/besean163/gophkeeper/internal/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type UserRepository struct {
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (r UserRepository) GetUser(id int) (*models.User, error) {
	connect, err := getDB()
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	connect.Find(user, id)
	return user, nil
}

func (r UserRepository) GetUserByLogin(login string) (*models.User, error) {
	connect, err := getDB()
	if err != nil {
		return nil, err
	}

	user := &models.User{}
	connect.Where("login  = ?", login).Find(user)
	return user, nil
}

func (r UserRepository) SaveUser(user *models.User) error {
	if user.ID == 0 {
		return r.insertUser(user)
	}

	return r.updateUser(user)
}

func (r UserRepository) updateUser(user *models.User) error {
	connect, err := getDB()
	if err != nil {
		return err
	}

	connect.Save(user)
	return nil
}

func (r UserRepository) insertUser(user *models.User) error {
	connect, err := getDB()
	if err != nil {
		return err
	}
	fmt.Println(user)
	connect.Create(user)
	return nil
}

func getDB() (*gorm.DB, error) {
	if db == nil {
		dsn := "postgres://gophkeeper:gophkeeper@localhost:5432/gophkeeper?sslmode=disable"
		conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		db = conn
	}
	return db, nil
}
