package resolver

import (
	"context"

	"github.com/Akshit8/go-meetup/graph/generated"
	"github.com/Akshit8/go-meetup/graph/model"
)

func (r *userResolver) Meetups(ctx context.Context, obj *model.User) ([]*model.Meetup, error) {
	panic("TODO")
}

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
