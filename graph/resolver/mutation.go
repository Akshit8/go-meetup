package resolver

import (
	"context"
	"errors"

	"github.com/Akshit8/go-meetup/graph/generated"
	"github.com/Akshit8/go-meetup/graph/model"
)

var (
	ErrInput = errors.New("input errors")
)

func (r *mutationResolver) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	isValid := validation(ctx, input)
	if !isValid {
		return nil, ErrInput
	}

	return r.Domain.Login(ctx, input)
}
func (r *mutationResolver) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	isValid := validation(ctx, input)
	if !isValid {
		return nil, ErrInput
	}

	return r.Domain.Register(ctx, input)
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	return r.Domain.CreateMeetup(ctx, input)
}

func (r *mutationResolver) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*model.Meetup, error) {
	return r.Domain.UpdateMeetup(ctx, id, input)
}

func (r *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	return r.Domain.DeleteMeetup(ctx, id)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
