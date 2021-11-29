package postgres

import (
	"github.com/go-pg/pg/v9"
	"github.com/phifertoo/go_graphql_server/models"
)

type MeetupsRepo struct {
	DB *pg.DB
}

// GetMeetups method
func (m *MeetupsRepo) GetMeetups() ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	err := m.DB.Model(&meetups).Select()

	if err != nil {
		return nil, err
	}
	return meetups, nil
}

func (m *MeetupsRepo) CreateMeetup(meetup *models.Meetup) (*models.Meetup, error) {
	_, err := m.DB.Model(meetup).Returning("*").Insert()

	return meetup, err
}
