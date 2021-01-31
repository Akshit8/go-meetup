package resolver

import (
	"context"
	"errors"
	"fmt"

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

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*model.Meetup, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
