// Package db impls db functionality
package db

import (
	"github.com/Akshit8/go-meetup/graph/model"
	"github.com/go-pg/pg/v10"
)

// MeetupRepo impls meetup storgae
type MeetupRepo struct {
	DB *pg.DB
}

// NewMeetupRepo inits a new MeetupRepo
func NewMeetupRepo(db *pg.DB) *MeetupRepo {
	return &MeetupRepo{
		DB: db,
	}
}

// GetMeetUps returns all meetups
func (mr *MeetupRepo) GetMeetUps() ([]*model.Meetup, error) {
	var meetups []*model.Meetup
	err := mr.DB.Model(&meetups).Select()
	if err != nil {
		return nil, err
	}
	return meetups, nil
}

// CreateMeetup creates a new meetup in db
func (mr *MeetupRepo) CreateMeetup(meetup *model.Meetup) (*model.Meetup, error) {
	_, err := mr.DB.Model(meetup).Insert()
	return meetup, err
}