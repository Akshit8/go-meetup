package resolver

import (
	"context"

	"github.com/Akshit8/go-meetup/graph/generated"
	"github.com/Akshit8/go-meetup/graph/model"
)

func (r *meetupResolver) User(ctx context.Context, obj *model.Meetup) (*model.User, error) {
	return r.UserStore.GetUserByID(obj.UserID)
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

type meetupResolver struct{ *Resolver }
