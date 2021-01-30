package resolver

import (
	"context"
	"errors"

	"github.com/Akshit8/go-meetup/graph/generated"
	"github.com/Akshit8/go-meetup/graph/model"
)

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

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
