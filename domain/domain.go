package domain

import (
	"errors"

	"github.com/Akshit8/go-meetup/db"
	"github.com/Akshit8/go-meetup/graph/model"
)

var (
	ErrBadCredentials  = errors.New("email and password don't match")
	ErrUnauthenticated = errors.New("unauthenticated")
	ErrForbidden       = errors.New("unauthorized")
)

type Domain struct {
	UserStore   db.UserRepo
	MeetupStore db.MeetupRepo
}

func NewDomain(userStore db.UserRepo, meetupStore db.MeetupRepo) *Domain {
	return &Domain{UserStore: userStore, MeetupStore: meetupStore}
}

type Owner interface {
	IsOwner(user *model.User) bool
}

func checkOwnerShip(o Owner, user *model.User) bool {
	return o.IsOwner(user)
}
