// Package db impls db functionality
package db

import (
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

// GetUserByID return a new user by ID
func (ur *UserRepo) GetUserByID(id string) (*model.User, error) {
	var user model.User
	err := ur.DB.Model(&user).Where("id = ?", id).First()
	if err != nil {
		return nil, err
	}
	return &user, nil
}

// CreateUser creates a new user in db
func (ur *UserRepo) CreateUser(user *model.User) (*model.User, error) {
	_, err := ur.DB.Model(user).Insert()
	return user, err
}