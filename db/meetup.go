// Package db impls db functionality
package db

import (
	"fmt"

	"github.com/Akshit8/go-meetup/graph/model"
	"github.com/go-pg/pg/v10"
)

// MeetupRepo impls meetup storage
type MeetupRepo struct {
	DB *pg.DB
}

// NewMeetupRepo inits a new MeetupRepo
func NewMeetupRepo(db *pg.DB) *MeetupRepo {
	return &MeetupRepo{
		DB: db,
	}
}

// GetMeetupByID returns a meetup for given id
func (mr *MeetupRepo) GetMeetupByID(id string) (*model.Meetup, error) {
	var meetup model.Meetup
	err := mr.DB.Model(&meetup).Where("id = ?", id).First()
	return &meetup, err
}

// GetMeetUps returns all meetups
func (mr *MeetupRepo) GetMeetUps(filter *model.MeetupFilter, limit, offset *int) ([]*model.Meetup, error) {
	var meetups []*model.Meetup
	query := mr.DB.Model(&meetups).Order("id")

	if filter != nil {
		if filter.Name != nil && *filter.Name != "" {
			query.Where("name LIKE ?", fmt.Sprintf("%%%s%%", *filter.Name))
		}
	}
	query.Limit(*limit)
	query.Offset(*offset)
	err := query.Select()

	if err != nil {
		return nil, err
	}
	return meetups, nil
}

// GetMeetupForUser returns all meetups for a given user
func (mr *MeetupRepo) GetMeetupForUser(user *model.User) ([]*model.Meetup, error) {
	var meetups []*model.Meetup
	err := mr.DB.Model(&meetups).Where("user_id = ?", user.ID).Select()
	return meetups, err
}

// CreateMeetup creates a new meetup in db
func (mr *MeetupRepo) CreateMeetup(meetup *model.Meetup) (*model.Meetup, error) {
	_, err := mr.DB.Model(meetup).Insert()
	return meetup, err
}

// UpdateMeetup updates an existing meetup
func (mr *MeetupRepo) UpdateMeetup(meetup *model.Meetup) (*model.Meetup, error) {
	_, err := mr.DB.Model(meetup).Where("id = ?", meetup.ID).Update()
	return meetup, err
}

// DeleteMeetup deletes a meetup
func (mr *MeetupRepo) DeleteMeetup(meetup *model.Meetup) error {
	_, err := mr.DB.Model(meetup).Where("id = ?", meetup.ID).Delete()
	return err
}
