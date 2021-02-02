package domain

import (
	"context"
	"errors"
	"fmt"

	"github.com/Akshit8/go-meetup/graph/model"
	auth "github.com/Akshit8/go-meetup/middleware"
)

func (d *Domain) CreateMeetup(ctx context.Context, input model.NewMeetup) (*model.Meetup, error) {
	currentUser, err := auth.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
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
		UserID:      currentUser.ID,
	}
	return d.MeetupStore.CreateMeetup(meetup)
}

func (d *Domain) UpdateMeetup(ctx context.Context, id string, input model.UpdateMeetup) (*model.Meetup, error) {
	meetup, err := d.MeetupStore.GetMeetupByID(id)
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

	meetup, err = d.MeetupStore.UpdateMeetup(meetup)
	if err != nil {
		return nil, fmt.Errorf("error while updating meetup: %v", err)
	}

	return meetup, nil
}

func (d *Domain) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	meetup, err := d.MeetupStore.GetMeetupByID(id)
	if err != nil || meetup == nil {
		return false, errors.New("meetup not exists")
	}

	err = d.MeetupStore.DeleteMeetup(meetup)
	if err != nil {
		return false, errors.New("error while deleting meetup")
	}

	return true, nil
}
