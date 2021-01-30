package resolver

import (
	"github.com/Akshit8/go-meetup/db"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver godoc
type Resolver struct {
	MeetupStore *db.MeetupRepo
	UserStore   *db.UserRepo
}
