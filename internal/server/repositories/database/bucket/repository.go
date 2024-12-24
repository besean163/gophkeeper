package bucket

import (
	"github.com/besean163/gophkeeper/internal/logger"
	"github.com/besean163/gophkeeper/internal/server/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type BucketRepository struct {
}

func NewBucketRepository() BucketRepository {
	return BucketRepository{}
}

func (r BucketRepository) GetAccount(id int) (*models.Account, error) {
	connect, err := getDB()
	if err != nil {
		return nil, err
	}

	var account *models.Account
	connect.Find(&account)
	return account, nil
}

func (r BucketRepository) GetAccounts(user models.User) ([]*models.Account, error) {
	connect, err := getDB()
	if err != nil {
		return nil, err
	}

	accounts := []*models.Account{}
	connect.Where("user_id = ?", user.ID).Find(&accounts)
	return accounts, nil
}

func (r BucketRepository) CreateAccount(account models.Account) error {
	connect, err := getDB()
	if err != nil {
		return err
	}

	accounts := []*models.Account{}
	connect.Find(&accounts)
	return nil
}

func (r BucketRepository) SaveAccount(account models.Account) error {
	if account.ID == 0 {
		return r.insertAccount(account)
	}

	return r.updateAccount(account)
}

func (r BucketRepository) DeleteAccount(id int) error {
	connect, err := getDB()
	if err != nil {
		return err
	}

	logger.Debug("here")
	connect.Delete(&models.Account{}, id)
	return nil
}

func (r BucketRepository) UpdateAccount(account models.Account) error {
	connect, err := getDB()
	if err != nil {
		return err
	}

	connect.Save(account)
	return nil
}

func (r BucketRepository) updateAccount(account models.Account) error {
	connect, err := getDB()
	if err != nil {
		return err
	}

	connect.Save(account)
	return nil
}

func (r BucketRepository) insertAccount(account models.Account) error {
	connect, err := getDB()
	if err != nil {
		return err
	}
	connect.Create(&account)
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
