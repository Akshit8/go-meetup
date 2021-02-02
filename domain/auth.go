package domain

import (
	"context"
	"errors"
	"log"

	"github.com/Akshit8/go-meetup/graph/model"
)

// Login login an existing user
func (d *Domain) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	user, err := d.UserStore.GetUserByUserEmail(input.Email)
	if err != nil {
		return nil, ErrBadCredentials
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		return nil, ErrBadCredentials
	}

	token, err := user.GenerateToken()
	if err != nil {
		return nil, errors.New("internal server error")
	}

	res := &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}

	return res, nil
}

// Register registers a new user
func (d *Domain) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	u, err := d.UserStore.GetUserByUserEmail(input.Email)
	log.Printf("user: %v", u)
	log.Printf("error: %v", err)
	if err == nil {
		return nil, errors.New("email already used")
	}

	_, err = d.UserStore.GetUserByUserName(input.Username)
	if err == nil {
		return nil, errors.New("username already used")
	}

	user := &model.User{
		Username:  input.Username,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	err = user.HashPassword(input.Password)
	if err != nil {
		log.Printf("error while hashing password: %v", err)
		return nil, errors.New("internal server error")
	}

	// TODO: Impl transaction
	_, err = d.UserStore.CreateUser(user)
	if err != nil {
		log.Printf("error creating user: %v", err)
		return nil, err
	}

	token, err := user.GenerateToken()
	if err != nil {
		log.Printf("error generating token: %v", err)
		return nil, errors.New("internal server error")
	}

	res := &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}

	return res, nil
}
