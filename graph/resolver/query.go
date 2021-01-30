package resolver

import (
	"context"

	"github.com/Akshit8/go-meetup/graph/generated"
	"github.com/Akshit8/go-meetup/graph/model"
)

func (r *queryResolver) Meetups(ctx context.Context) ([]*model.Meetup, error) {
	return r.MeetupStore.GetMeetUps()
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
