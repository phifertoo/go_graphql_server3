//go:generate go run github.com/99designs/gqlgen -v

package graph

import (
	"github.com/phifertoo/go_graphql_server/models"
	"github.com/phifertoo/go_graphql_server/postgres"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// In postgres/meetups.go, we have the MeetupsRepo struct
// In postgres/users.go, we have the UsersRepo struct
// We neeed to add them to this root resolver so that we have access to them
type Resolver struct {
	MeetupsRepo postgres.MeetupsRepo
	UsersRepo   postgres.UsersRepo
}

// setup a state
var meetups = []*models.Meetup{
	{
		ID:          "1",
		Name:        "A meetup",
		Description: "A description",
		UserID:      "1",
	},
	{
		ID:          "2",
		Name:        "A second meetup",
		Description: "A second description",
		UserID:      "2",
	},
}

var users = []*models.User{
	{
		ID:       "1",
		Username: "Bob",
		Email:    "bob@gmail.com",
	},
	{
		ID:       "2",
		Username: "Jon",
		Email:    "bob@gmail.com",
	},
}
