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

var (
	// ErrNotFound signifies a resource not present in the DB
	ErrNotFound = errors.New("resource not found")
	// ErrInvalidId signifies that an invalid ID was provided to a method like delete
	ErrInvalidId = errors.New("ID provided was invalid")
)

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

// first is a helper method for fetching the first item using the dest interface
// and storing the result into dest if found
func first(db *gorm.DB, dest interface{}) error {
	err := db.First(dest).Error
	if err == gorm.ErrRecordNotFound {
		return ErrNotFound
	}
	return err
}

// ByID is used to look up a user record with the provided id
// It returns a nil error if the record is found and ErrNotFound otherwise
// If there are other DB errors that come up, the corresponding error is returned
func (service *UserService) ByID(id uint) (*User, error) {
	var user User
	filtered := service.db.Where("id = ?", id)
	err := first(filtered, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// ByEmail is used to look up a user record with the given email id
func (service *UserService) ByEmail(email string) (*User, error) {
	var user User
	filtered := service.db.Where("email = ?", email)
	err := first(filtered, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// ResetDB drops and recreates the user table to make testing easier
func (service *UserService) ResetDB() error {
	err := service.db.DropTableIfExists(&User{}).Error
	if err != nil {
		return err
	}
	return service.AutoMigrate()
}

// Create will create a user entry in the DB and update
// the provided user object with ID, CreatedAt, UpdatedAt fields
func (service *UserService) Create(user *User) error {
	return service.db.Create(user).Error
}

func (service *UserService) Update(user *User) error {
	return service.db.Save(user).Error
}

func (service *UserService) Delete(id uint) error {
	if id == 0 { // if id is not provided or is 0, all records will be deleted
		return ErrInvalidId
	}
	// initialised like this but can be accessed by user.ID
	user := User{Model: gorm.Model{ID: id}}
	return service.db.Delete(&user).Error
}

// Automigrate will automatically migrate the user table
func (service *UserService) AutoMigrate() error {
	return service.db.AutoMigrate(&User{}).Error
}
