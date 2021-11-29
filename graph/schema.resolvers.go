package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"

	"github.com/phifertoo/go_graphql_server/graph/generated"
	"github.com/phifertoo/go_graphql_server/graph/model"
	"github.com/phifertoo/go_graphql_server/models"
)

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	return r.UsersRepo.GetUserByID(obj.UserID)
	// user := new(models.User)
	// for _, u := range users {
	// 	if u.ID == obj.UserID {
	// 		user = u
	// 		break
	// 	}
	// }
	// return user, nil
}

func (r *mutationResolver) CreateMeetup(ctx context.Context, input model.NewMeetup) (*models.Meetup, error) {
	if len(input.Name) < 3 {
		return nil, errors.New("name not long enough")
	}
	if len(input.Description) < 3 {
		return nil, errors.New("description not long enough")
	}
	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      "1",
	}

	return r.MeetupsRepo.CreateMeetup(meetup)
	// panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	// return r.MeetupsRepo.GetMeetups()
	return meetups, nil
}

func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	var m []*models.Meetup

	for _, meetup := range meetups {
		if meetup.UserID == obj.ID {
			meetups = append(m, meetup)
		}
	}
	return meetups, nil
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// User returns generated.UserResolver implementation.
func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type meetupResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type userResolver struct{ *Resolver }
