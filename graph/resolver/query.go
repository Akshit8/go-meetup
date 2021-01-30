package resolver

import (
	"context"
	"fmt"

	"github.com/Akshit8/go-meetup/graph/generated"
	"github.com/Akshit8/go-meetup/graph/model"
)

func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {
	return r.MeetupStore.GetMeetUps()
}

func (r *queryResolver) User(ctx context.Context, id string) (*model.User, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
