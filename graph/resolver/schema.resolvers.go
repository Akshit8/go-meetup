package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Akshit8/go-meetup/graph/generated"
	"github.com/Akshit8/go-meetup/graph/model"
)

func (r *meetupResolver) User(ctx context.Context, obj *model.Meetup) (*model.User, error) {
	return dataloader.GetUserLoader(ctx).Load(obj.UserID)
}

func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
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

func (r *queryResolver) Meetups(ctx context.Context, filter *model.MeetupFilter, limit *int, offset *int) ([]*model.Meetup, error) {
	return r.MeetupStore.GetMeetUps(filter, limit, offset)
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	return r.UserStore.GetUserByID(id)
}

func (r *userResolver) FirstName(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) LastName(ctx context.Context, obj *model.User) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) Meetups(ctx context.Context, obj *model.User) ([]*model.Meetup, error) {
	return r.MeetupStore.GetMeetupForUser(obj)
}

func (r *userResolver) CreatedAt(ctx context.Context, obj *model.User) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *userResolver) UpdatedAt(ctx context.Context, obj *model.User) (*time.Time, error) {
	panic(fmt.Errorf("not implemented"))
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type meetupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
