package resolver

import (
	"context"
	"errors"
	"fmt"
	auth "github.com/Akshit8/go-meetup/middleware"
	"log"

	"github.com/Akshit8/go-meetup/graph/generated"
	"github.com/Akshit8/go-meetup/graph/model"
)

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	u, err := r.UserStore.GetUserByUserEmail(input.Email)
	log.Printf("user: %v", u)
	log.Printf("error: %v", err)
	if err == nil {
		return nil, errors.New("email already used")
	}

	_, err = r.UserStore.GetUserByUserName(input.Username)
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
	_, err = r.UserStore.CreateUser(user)
	if err != nil {
		log.Printf("error creating user: %v", err)
		return nil, err
	}

	token, err := user.GenerateToken()
	if err != nil {
		log.Printf("error generating token: %v", err)
		return nil, errors.New("internal server error")
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	currentUser, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil,
	}
	if len(input.Name) < 3 {
		return nil, errors.New("description not long enough")
	}
	if len(input.Description) < 10 {
		return nil, errors.New("description not long enough")
	}
	meetup := &model.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      "1",
	}
	return r.MeetupStore.CreateMeetup(meetup)
}

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*model.Meetup, error) {
	meetup, err := r.MeetupStore.GetMeetupByID(id)
	if err != nil || meetup == nil {
		return nil, errors.New("meetup not exists")
	}

	didUpdate := false

	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, errors.New("name too short")
		}
		meetup.Name = *input.Name
		didUpdate = true
	}
	if input.Description != nil {
		if len(*input.Description) < 10 {
			return nil, errors.New("description too short")
		}
		meetup.Description = *input.Description
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("no update found")
	}

	meetup, err = r.MeetupStore.UpdateMeetup(meetup)
	if err != nil {
		return nil, fmt.Errorf("error while updating meetup: %v", err)
	}

	return meetup, nil
}

func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	meetup, err := r.MeetupStore.GetMeetupByID(id)
	if err != nil || meetup == nil {
		return false, errors.New("meetup not exists")
	}

	err = r.MeetupStore.DeleteMeetup(meetup)
	if err != nil {
		return false, errors.New("error while deleting meetup")
	}

	return true, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
