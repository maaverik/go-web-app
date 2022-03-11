package models

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	gorm.Model // composition
	Name       string
	Email      string `gorm:"not null; unique index"`
}

// UserService is an abstraction around the DB connector that we use
type UserService struct {
	db *gorm.DB
}

// ErrNotFound signifies a resource not present in the DB
var ErrNotFound = errors.New("resource not found")

func NewUserService(connectionInfo string) (*UserService, error) {
	db, err := gorm.Open("postgres", connectionInfo)
	if err != nil {
		return nil, err
	}
	db.LogMode(true)
	return &UserService{
		db,
	}, nil
}

func (service *UserService) Close() error {
	return service.db.Close()
}

// ByID is used to look up a record with the provided id
// It returns a nil error if the record is found and ErrNotFound otherwise
// If there are other DB errors that come up, the corresponding error is returned
func (service *UserService) ByID(id uint) (*User, error) {
	var user User
	err := service.db.Where("id = ?", id).First(&user).Error
	switch err {
	case nil:
		return &user, nil
	case gorm.ErrRecordNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}

}

// ResetDB drops and recreats the user table to make testing easier
func (service *UserService) ResetDB() {
	service.db.DropTableIfExists(&User{})
	service.db.AutoMigrate(&User{})
}

// Create will create a user entry in the DB and update
// the provided user object with ID, CreatedAt, UpdatedAt fields
func (service *UserService) Create(user *User) error {
	return service.db.Create(user).Error
}
