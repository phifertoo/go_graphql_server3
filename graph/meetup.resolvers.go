package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/phifertoo/go_graphql_server/graph/generated"
	"github.com/phifertoo/go_graphql_server/models"
)

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	return getUserLoader(ctx).Load(obj.UserID)

	// return r.UsersRepo.GetUserByID(obj.UserID)
	// user := new(models.User)
	// for _, u := range users {
	// 	if u.ID == obj.UserID {
	// 		user = u
	// 		break
	// 	}
	// }

	// return user, nil
}

// Meetup returns generated.MeetupResolver implementation.
func (r *Resolver) Meetup() generated.MeetupResolver { return &meetupResolver{r} }

type meetupResolver struct{ *Resolver }
