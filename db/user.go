// Package db impls db functionality
package db

import (
	"fmt"

	"github.com/Akshit8/go-meetup/graph/model"
	"github.com/go-pg/pg/v10"
)

// UserRepo impls user storage
type UserRepo struct {
	DB *pg.DB
}

// NewUserRepo inits a new UserRepo
func NewUserRepo(db *pg.DB) *UserRepo {
	return &UserRepo{
		DB: db,
	}
}

// CreateUser creates a new user in db
func (ur *UserRepo) CreateUser(user *model.User) (*model.User, error) {
	_, err := ur.DB.Model(user).Insert()
	return user, err
}

// GetUserByField returns a user record against any filter
func (ur *UserRepo) GetUserByField(field, value string) (*model.User, error) {
	var user model.User
	err := ur.DB.Model(&user).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &user, err
}

// GetUserByID return a user by ID
func (ur *UserRepo) GetUserByID(id string) (*model.User, error) {
	return ur.GetUserByField("id", id)
}

// GetUserByUserEmail returns a user by email
func (ur *UserRepo) GetUserByUserEmail(email string) (*model.User, error) {
	return ur.GetUserByField("email", email)
}

// GetUserByUserName returns a user by username
func (ur *UserRepo) GetUserByUserName(username string) (*model.User, error) {
	return ur.GetUserByField("username", username)
}
